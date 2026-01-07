package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape [][]rune

type Region struct {
	Length     int
	Width      int
	ShapeCount []int
}

type Point struct {
	X, Y int
}
type PlacedShape struct {
	Shape Shape
	Pos   Point
}

type Bin struct {
	Width  int
	Length int
	Grid   [][]rune
	Placed []PlacedShape
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

	shape := make([][]rune, 0)
	var shapes []Shape
	var regions []Region
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsRune(line, 'x') {
			parts := strings.Split(line, ":")
			dims := strings.Split(parts[0], "x")
			length, err := strconv.Atoi(dims[0])
			if err != nil {
				fmt.Println("error converting string to int:", err)
				return
			}

			width, err := strconv.Atoi(dims[1])
			if err != nil {
				fmt.Println("error converting string to int:", err)
				return
			}

			shapeCountStr := strings.Fields(parts[1])
			shapeCount := make([]int, len(shapeCountStr))
			for i, count := range shapeCountStr {
				shapeCount[i], err = strconv.Atoi(count)
				if err != nil {
					fmt.Println("error converting string to int:", err)
					return
				}
			}

			region := Region{Length: length, Width: width, ShapeCount: shapeCount}
			regions = append(regions, region)
			continue
		} else if strings.ContainsRune(line, ':') {
			shape = make([][]rune, 0)
			continue
		} else if strings.ContainsRune(line, '#') {
			shape = append(shape, []rune(line))
		} else {
			shapes = append(shapes, shape)
		}
	}

	var numRegions int
	for _, region := range regions {
		bin := Bin{Width: region.Width, Length: region.Length, Grid: make([][]rune, region.Length)}
		for i := range bin.Grid {
			bin.Grid[i] = make([]rune, region.Width)
			for j := range bin.Grid[i] {
				bin.Grid[i][j] = '.'
			}
		}
		if CanFitShapes(&bin, shapes, region.ShapeCount) {
			numRegions++
		}
	}
	fmt.Println("Number of regions:", numRegions)
}

// Rotate shape 90 degrees clockwise
func Rotate(shape Shape) Shape {
	h, w := len(shape), len(shape[0])
	rotated := make(Shape, w)
	for i := range rotated {
		rotated[i] = make([]rune, h)
		for j := range rotated[i] {
			rotated[i][j] = shape[h-j-1][i]
		}
	}
	return rotated
}

// Generate all rotations (0°, 90°, 180°, 270°)
func AllRotations(shape Shape) []Shape {
	rotations := []Shape{shape}
	for i := 0; i < 3; i++ {
		shape = Rotate(shape)
		rotations = append(rotations, shape)
	}
	return rotations
}

func (b *Bin) CanPlace(shape Shape, x, y int) bool {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] == '#' {
				bx, by := x+j, y+i
				if bx >= b.Width || by >= b.Length || b.Grid[by][bx] == '#' {
					return false
				}
			}
		}
	}
	return true
}

func (b *Bin) Place(shape Shape, x, y int) {
	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] == '#' {
				b.Grid[y+i][x+j] = '#'
			}
		}
	}
	b.Placed = append(b.Placed, PlacedShape{Shape: shape, Pos: Point{X: x, Y: y}})
}

func PackShapes(bin *Bin, shapes []Shape, shapeCounts []int) bool {
	for i, shape := range shapes {
		count := shapeCounts[i]
		for c := 0; c < count; c++ {
			placed := false
			for _, rot := range AllRotations(shape) {
				for y := 0; y <= bin.Length-len(rot); y++ {
					for x := 0; x <= bin.Width-len(rot[0]); x++ {
						if bin.CanPlace(rot, x, y) {
							bin.Place(rot, x, y)
							placed = true
							break
						}
					}
					if placed {
						break
					}
				}
				if placed {
					break
				}
			}
			if !placed {
				return false
			}
		}
	}
	return true
}

func CanFitShapes(bin *Bin, shapes []Shape, shapeCounts []int) bool {
	binArea := bin.Width * bin.Length
	totalShapeCells := 0
	maxPossibleCells := 0
	for i, shape := range shapes {
		shapeArea := len(shape) * len(shape[0])
		count := shapeCounts[i]
		if count == 0 {
			continue
		}
		shapeCells := 0
		for _, row := range shape {
			for _, cell := range row {
				if cell == '#' {
					shapeCells++
				}
			}
		}
		totalShapeCells += shapeCells * count
		maxPossibleCells += shapeArea * count
	}

	if totalShapeCells > binArea {
		return false // Too many shape cells
	}
	if maxPossibleCells <= binArea {
		return true // All max-sized shapes fit
	}
	return PackShapes(bin, shapes, shapeCounts) // Try to pack
}
