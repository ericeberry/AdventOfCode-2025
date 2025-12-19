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

	memo := make(map[string]int)
	paths := findPaths(devices, "svr", memo, false, false)
	fmt.Println("Total is ", paths)
}

func findPaths(devices map[string][]string, name string, memo map[string]int, fft bool, dac bool) int {
	if name == "out" {
		if fft && dac {
			return 1
		}
		return 0
	}

	key := fmt.Sprintf("%s-%v-%v", name, fft, dac)
	if val, ok := memo[key]; ok {
		return val
	}

	result := 0
	if outputs, ok := devices[name]; ok {
		for _, output := range outputs {
			result += findPaths(devices, output, memo, fft || name == "fft", dac || name == "dac")
		}
	}

	memo[key] = result
	return result
}
