package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var counter2 = 0

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("error opening file input:", err)
		return
	}

	ranges := bytes.Split(data, []byte(","))
	for _, r := range ranges {
		bounds := bytes.Split(r, []byte("-"))
		analyzeRange2(bounds)
	}

	fmt.Println("final counter:", counter2)
}

func analyzeRange2(bounds [][]byte) {
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
		length := len(s)
		for j := 2; j <= length; j++ {
			partSize := length / j
			remainder := length % j
			if remainder != 0 {
				continue
			}

			match := true
			firstPart := s[:partSize]
			for k := 1; k < length/partSize; k++ {
				if (k+1)*partSize > length {
					fmt.Println("checking part", k, "of size", partSize)
				}
				newPart := s[k*partSize : (k+1)*partSize]
				if newPart != firstPart {
					match = false
					break
				}
			}
			if match {
				fmt.Println("found matching number:", i)
				counter2 += i
				break
			}
		}
	}
}
