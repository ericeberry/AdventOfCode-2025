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

	grid := make([][]Node, 0)
	x := 0
	y := 0
	var start Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := []rune(line)
		row := make([]Node, 0)
		for _, part := range parts {
			node := Node{X: x, Y: y, Val: part, Count: 0}
			row = append(row, node)
			x++

			if node.Val == 'S' {
				start = node
			}
		}

		grid = append(grid, row)
		x = 0
		y++
	}

	result := scanTree(grid, start)
	fmt.Println("result is ", result)

}

func scanTree(grid [][]Node, start Node) int {
	var queue []*Node
	queue = append(queue, &start)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == '.' {
			if len(grid) > node.Y+1 {
				grid[node.Y+1][node.X].Count += node.Count
				if !Contains(queue, node.Y+1, node.X) {
					queue = append(queue, &grid[node.Y+1][node.X])
				}
			}
		} else if node.Val == '^' {
			if node.X-1 >= 0 && len(grid) > node.Y+1 {
				left := &grid[node.Y+1][node.X-1]
				left.Count += node.Count
				if !Contains(queue, node.Y+1, node.X-1) {
					queue = append(queue, &grid[node.Y+1][node.X-1])
				}
			} else {
				fmt.Println("Error node out of bounds")
				break
			}

			if len(grid[node.Y]) > node.X+1 && len(grid) > node.Y+1 {
				right := &grid[node.Y+1][node.X+1]
				right.Count += node.Count
				if !Contains(queue, node.Y+1, node.X+1) {
					queue = append(queue, &grid[node.Y+1][node.X+1])
				}
			} else {
				fmt.Println("Error node out of bounds")
				break
			}
		} else {
			grid[node.Y+1][node.X].Count = 1
			queue = append(queue, &grid[node.Y+1][node.X])
		}
	}

	result := 0
	for _, row := range grid[len(grid)-1] {
		result += row.Count
	}

	return result
}

func Contains(queue []*Node, y int, x int) bool {
	for _, node := range queue {
		if node.X == x && node.Y == y {
			return true
		}
	}
	return false
}
