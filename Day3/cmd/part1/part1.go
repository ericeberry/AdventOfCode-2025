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

	value := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value += calculateJoltage(line)
	}

	fmt.Println("result is ", value)
}

func calculateJoltage(line string) int {
	pos := 0
	max := 0
	for i := 0; i < len(line)-1; i++ {
		val := int(line[i] - '0')
		if max < val {
			max = val
			pos = i
		}

		if max == 9 {
			break
		}
	}

	num := max
	max = 0
	for i := pos + 1; i < len(line); i++ {
		val := int(line[i] - '0')
		if max < val {
			max = val
		}

		if max == 9 {
			break
		}
	}

	num = num*10 + max
	fmt.Println(num)
	return num
}
