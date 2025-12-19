package main

import (
	"bufio"
	"fmt"
	"os"
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

	devices := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		outputs := strings.Fields(parts[1])
		devices[parts[0]] = outputs
	}

	paths := findPaths(devices, "you")
	fmt.Println("Total is ", paths)
}

func findPaths(devices map[string][]string, next string) int {
	if next == "out" {
		return 1
	}

	result := 0
	if outputs, ok := devices[next]; ok {
		for _, output := range outputs {
			result += findPaths(devices, output)
		}
	}

	return result
}
