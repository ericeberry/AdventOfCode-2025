package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var counter = 0

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("error opening file input:", err)
		return
	}

	ranges := bytes.Split(data, []byte(","))
	for _, r := range ranges {
		bounds := bytes.Split(r, []byte("-"))
		analyzeRange(bounds)
	}

	fmt.Println("final counter:", counter)
}

func analyzeRange(bounds [][]byte) {
	lower, err := strconv.Atoi(string(bounds[0]))
	if err != nil {
		fmt.Println("error converting lower bound", err)
		return
	}

	upper, err := strconv.Atoi(string(bounds[1]))
	if err != nil {
		fmt.Println("error converting lower bound", err)
		return
	}

	for i := lower; i <= upper; i++ {
		s := strconv.Itoa(i)
		mid := len(s) / 2

		part1 := s[:mid]
		part2 := s[mid:]
		if part1 == part2 {
			fmt.Println("found matching number:", i)
			counter += i
		}
	}
}
