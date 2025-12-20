package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Rectangle struct {
	area        int
	leftCorner  Point
	rightCorner Point
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("error opening file input.txt:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var points []Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error parsing float:", err)
			return
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("error parsing float:", err)
			return
		}

		points = append(points, Point{x: x, y: y})
	}

	var rectangles []Rectangle
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			area := (abs(points[j].x-points[i].x) + 1) * (abs(points[j].y-points[i].y) + 1)
			rectangles = append(rectangles, Rectangle{area: area, leftCorner: points[i], rightCorner: points[j]})
		}
	}

	slices.SortFunc(rectangles, func(r1, r2 Rectangle) int {
		if r1.area > r2.area {
			return -1
		}
		if r1.area < r2.area {
			return 1
		}
		return 0
	})

	for _, rectangle := range rectangles {
		if checkRectangle(points, rectangle) {
			fmt.Println("result is ", rectangle.area)
			return
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func onSegment(p, q, r Point) bool {
	return q.x <= max(p.x, r.x) && q.x >= min(p.x, r.x) &&
		q.y <= max(p.y, r.y) && q.y >= min(p.y, r.y)
}

func orientation(p, q, r Point) int {
	val := (q.y-p.y)*(r.x-q.x) - (q.x-p.x)*(r.y-q.y)
	if val == 0 {
		return 0 // Collinear
	}
	if val > 0 {
		return 1 // Clockwise
	}
	return 2 // Counterclockwise
}

// isInsidePolygon checks if the point P is inside the polygon (including the boundary).
// The polygon is defined by the slice of Points.
func isInsidePolygon(polygon []Point, p Point) bool {
	n := len(polygon)
	if n < 3 {
		return false // Not a polygon
	}

	// Count the number of times a ray from P (to the right) intersects the polygon edges.
	intersections := 0

	// Ray direction: +X (to the right, theoretically at Y=p.y)

	for i := 0; i < n; i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%n] // Connects back to the first point

		// 1. Check if the point P is on the boundary of the current edge (p1, p2)
		if orientation(p1, p, p2) == 0 && onSegment(p1, p, p2) {
			return true // Point is on the boundary
		}

		// 2. Standard Ray Casting Logic (excluding horizontal edges that coincide with the ray)
		// Check if the edge (p1, p2) straddles the horizontal line y = p.y
		if (p1.y > p.y) != (p2.y > p.y) {
			// Calculate the X-coordinate of the intersection point of the edge and the ray
			// Intersection X = p1.x + (p.y - p1.y) * (p2.x - p1.x) / (p2.y - p1.y)

			// Handle vertical edges separately to avoid division by zero.
			if abs(p1.y-p2.y) == 0 { // Horizontal edge (already handled by orientation check)
				continue
			}

			xIntersect := p1.x + (p.y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y)

			// If the intersection is to the right of P (xIntersect > p.x), increment count
			if xIntersect > p.x {
				intersections++
			}
		}
	}

	// If the intersection count is odd, the point is inside; otherwise, it's outside.
	return intersections%2 == 1
}

// --- Main Check Function ---

// getRectangleCorners generates the four ordered corners of the rectangle.
func getRectangleCorners(r Rectangle) [4]Point {
	minX := min(r.leftCorner.x, r.rightCorner.x)
	maxX := max(r.leftCorner.x, r.rightCorner.x)
	minY := min(r.leftCorner.y, r.rightCorner.y)
	maxY := max(r.leftCorner.y, r.rightCorner.y)

	// Corners in clockwise order
	return [4]Point{
		{x: minX, y: minY}, // Bottom-Left
		{x: maxX, y: minY}, // Bottom-Right
		{x: maxX, y: maxY}, // Top-Right
		{x: minX, y: maxY}, // Top-Left
	}
}

// segmentIntersectStrictly checks if segment P1Q1 strictly intersects P2Q2.
// For the purpose of rectangle containment, we only care about intersections
// that occur at interior points of the rectangle edge P1Q1.
// If the intersection is at an endpoint of P1Q1, we've already verified it's
// inside or on the boundary of the polygon.
func segmentIntersectStrictly(p1, q1, p2, q2 Point) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// General Case: Segments intersect.
	if o1 != o2 && o3 != o4 {
		// If the intersection is at an endpoint of either segment,
		// it is not a "strict" crossing of the boundary for containment purposes.
		// Since both the polygon and rectangles are axis-aligned (rectilinear),
		// any crossing that would put the rectangle outside must occur at an
		// interior point of both segments.
		if o1 == 0 || o2 == 0 || o3 == 0 || o4 == 0 {
			return false
		}
		return true
	}

	return false
}

// CheckRectangle determines if the rectangle R resides entirely within the polygon.
// A residence on the boundary is allowed.
func checkRectangle(polygon []Point, r Rectangle) bool {
	rectCorners := getRectangleCorners(r)

	// --- 1. Check if all four corners of the rectangle are inside the polygon (or on the boundary) ---
	for _, corner := range rectCorners {
		if !isInsidePolygon(polygon, corner) {
			return false
		}
	}

	// --- 2. Check if any edge of the rectangle strictly intersects any edge of the polygon ---

	// If the rectangle corners are all inside, but an edge crosses, it means the
	// rectangle is too big for a concave section of the polygon.

	// Rectangle Edges (R1Q1)
	for i := 0; i < 4; i++ {
		rP1 := rectCorners[i]
		rP2 := rectCorners[(i+1)%4]

		// Polygon Edges (P1Q2)
		n := len(polygon)
		for j := 0; j < n; j++ {
			pP1 := polygon[j]
			pP2 := polygon[(j+1)%n]

			// If the segments strictly intersect, the rectangle is not contained.
			if segmentIntersectStrictly(rP1, rP2, pP1, pP2) {
				return false
			}
		}
	}

	// If all corners are inside/on the boundary, and no edges strictly cross, the rectangle is contained.
	return true
}
