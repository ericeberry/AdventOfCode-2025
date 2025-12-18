package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}

	num := 0
	result := 1
	for result > 0 {
		result, grid = countGrid(grid)
		fmt.Println("result is ", result)
		num += result
	}
	fmt.Println("num is ", num)
}

func countGrid(grid [][]rune) (int, [][]rune) {
	m := len(grid)
	n := len(grid[0])
	count := make([][]int, m)
	for i := range count {
		count[i] = make([]int, n)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '@' {
				count[i][j] = -1
				continue
			}

			if j < n-1 {
				if grid[i][j+1] == '@' {
					count[i][j+1]++
					count[i][j]++
				}

				if i < m-1 && grid[i+1][j+1] == '@' {
					count[i+1][j+1]++
					count[i][j]++
				}
			}

			if i < m-1 {
				if grid[i+1][j] == '@' {
					count[i+1][j]++
					count[i][j]++
				}

				if j-1 >= 0 && grid[i+1][j-1] == '@' {
					count[i+1][j-1]++
					count[i][j]++
				}
			}
		}
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if count[i][j] >= 0 && count[i][j] < 4 {
				grid[i][j] = '.'
				result++
			}
		}
	}

	return result, grid
}

