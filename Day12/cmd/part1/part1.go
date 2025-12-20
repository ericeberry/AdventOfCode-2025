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
	length     int
	width      int
	shapeCount []int
}

func main() {
	file, err := os.Open("./sample.txt")
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

			region := Region{length: length, width: width, shapeCount: shapeCount}
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

	fmt.Println(regions)
}
