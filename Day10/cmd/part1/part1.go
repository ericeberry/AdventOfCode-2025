package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ButtonPress struct {
	result int
	count  int
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

	var indicators []int
	var buttons [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		indicator := parseIndicator(parts[0])
		indicators = append(indicators, indicator)

		var row []int
		for i := 1; i < len(parts)-1; i++ {
			button := parseButtons(parts[i])
			row = append(row, button)
		}
		buttons = append(buttons, row)
	}

	var total int
	for i := 0; i < len(indicators); i++ {
		total += calculateButtonPresses(indicators[i], buttons[i])
	}
	fmt.Println("total is ", total)
}

func parseIndicator(s string) int {
	lights := stripRunes(s, "[]")
	result := 0
	for i, l := range lights {
		if l == '#' {
			result += 1 << i
		}
	}
	return result
}

func parseButtons(s string) int {
	buttons := strings.Split(stripRunes(s, "()"), ",")
	result := 0
	for _, b := range buttons {
		val, err := strconv.Atoi(b)
		if err != nil {
		}
		result += 1 << val
	}
	return result
}

func stripRunes(s string, toStrip string) string {
	stripSet := make(map[rune]struct{})
	for _, r := range toStrip {
		stripSet[r] = struct{}{}
	}
	return strings.Map(func(r rune) rune {
		if _, found := stripSet[r]; found {
			return -1
		}
		return r
	}, s)
}

func calculateButtonPresses(indicator int, buttons []int) int {
	queue := []ButtonPress{{result: 0, count: 0}}
	for len(queue) > 0 {
		buttonPress := queue[0]
		queue = queue[1:]

		for _, button := range buttons {
			newButtonPress := ButtonPress{result: buttonPress.result ^ button, count: buttonPress.count + 1}
			if newButtonPress.result == indicator {
				fmt.Println("count is ", newButtonPress.count)
				return newButtonPress.count
			}
			queue = append(queue, newButtonPress)
		}
	}

	return 0
}
