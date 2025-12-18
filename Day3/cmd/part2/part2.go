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
		value += calculateJoltage12(line)
	}

	fmt.Println("result is ", value)
}

func calculateJoltage12(line string) int {
	arr := make([]int, 12)

	i := len(line) - 1
	j := 11
	for ; i >= len(line)-12; i-- {
		arr[j] = int(line[i] - '0')
		j--
	}

	for ; i >= 0; i-- {
		num := int(line[i] - '0')

		for j := 0; j < 12; j++ {
			if num >= arr[j] {
				tmp := arr[j]
				arr[j] = num
				num = tmp
			} else {
				break
			}
		}
	}

	num := 0
	for i := 0; i < 12; i++ {
		fmt.Print(arr[i])
		num = num*10 + arr[i]
	}

	fmt.Println()
	return num
}
