package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

var counter = 0

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

	value := 50
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("error converting string to int:", err)
			return
		}

		if direction == 'L' {
			value = left(value, count)
		} else if direction == 'R' {
			value = right(value, count)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file:", err)
	}

	fmt.Println("Final count:", counter)
}

func left(initialValue int, count int) int {
	counter = counter + ((count - initialValue + 99) / 100)
	temp := ((initialValue-count)%100 + 100) % 100
	if temp == 0 && count > 0 {
		counter++
	}
	if initialValue == 0 {
		counter--
	}
	return temp
}

func right(initialValue int, count int) int {
	counter = counter + ((initialValue + count) / 100)
	temp := (initialValue + count) % 100
	return temp
}
