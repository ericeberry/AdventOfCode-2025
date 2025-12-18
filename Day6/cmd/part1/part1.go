package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	var grid [][]int
	var ops []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if grid == nil {
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("error converting part to int:", err)
					return
				}
				row := []int{num}
				grid = append(grid, row)
			}
		} else {
			for j, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					ops = append(ops, part)
				} else {
					grid[j] = append(grid[j], num)
				}
			}
		}
	}

	result := 0
	for i, row := range grid {
		var num int
		for j, val := range row {
			if j == 0 {
				num = val
			} else if ops[i] == "+" {
				num += val
			} else if ops[i] == "*" {
				num *= val
			}
		}

		result += num
	}
	fmt.Println("result is ", result)

}
