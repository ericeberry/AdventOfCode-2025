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

	var grid [][]rune
	var ops []rune
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		prevLine := scanner.Text()
		for scanner.Scan() {
			line := scanner.Text()
			runes := []rune(prevLine)
			if grid == nil {
				for _, ch := range runes {
					row := []rune{ch}
					grid = append(grid, row)
				}
			} else {
				for j, ch := range runes {
					if j >= len(grid) {
						row := []rune{ch}
						grid = append(grid, row)
					} else {
						grid[j] = append(grid[j], ch)
					}
				}
			}
			prevLine = line
		}

		ops = []rune(prevLine)
		for i := len(ops); i < len(grid); i++ {
			ops = append(ops, ' ')
		}
	}

	var nums []int
	for _, row := range grid {
		str := string(row)
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			nums = append(nums, -1)
		} else {
			nums = append(nums, num)
		}
	}

	var currentOp rune = ' '
	result := 0
	val := -1
	for i, num := range nums {
		if num == -1 {
			result += val
			val = -1
			currentOp = ' '
		} else if val == -1 {
			val = num
			currentOp = ops[i]
		} else if currentOp == '+' {
			val += num
		} else if currentOp == '*' {
			val *= num
		}
	}

	result += val
	fmt.Println("result is ", result)

}
