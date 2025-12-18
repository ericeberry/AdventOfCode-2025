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

	slices.SortFunc(points, func(p1, p2 Point) int {
		if p1.x < p2.x {
			return -1
		}
		if p1.x > p2.x {
			return 1
		}
		if p1.y < p2.y {
			return -1
		}
		if p1.y > p2.y {
			return 1
		}
		return 0
	})

	var rectangles []Rectangle
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			var area int
			if points[i].y < points[j].y {
				area = (points[j].x - points[i].x + 1) * (points[j].y - points[i].y + 1)
			} else {
				area = (points[i].x - points[j].x + 1) * (points[i].y - points[j].y + 1)
			}
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

	fmt.Println("result is ", rectangles[0].area)
}
