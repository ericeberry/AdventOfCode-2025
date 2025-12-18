package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var input []string

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

	var buttons [][][]int
	var joltages [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
		parts := strings.Split(line, " ")

		var row [][]int
		for i := 1; i < len(parts)-1; i++ {
			button := parseButtons(parts[i])
			row = append(row, button)
		}
		buttons = append(buttons, row)

		joltages = append(joltages, parseJoltage(parts[len(parts)-1]))
	}

	var total int
	for i := 0; i < len(joltages); i++ {
		total += calculateJoltage(joltages[i], buttons[i])
		fmt.Println("  ", input[i])
	}
	fmt.Println("total is ", total)
}

func parseButtons(s string) []int {
	buttons := strings.Split(stripRunes(s, "()"), ",")
	var result []int

	for _, b := range buttons {
		val, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}
	return result
}

func parseJoltage(s string) []int {
	buttons := strings.Split(stripRunes(s, "{}"), ",")
	var result []int
	for _, b := range buttons {
		val, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
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

func calculateJoltage(joltage []int, buttons [][]int) int {
	cols := len(buttons)
	rows := len(joltage)

	b := make([]int, len(joltage))
	for i, eqResult := range joltage {
		b[i] = eqResult
	}

	// Build coefficient matrix
	A := make([][]int, rows)
	for i := range A {
		A[i] = make([]int, cols)
	}
	for col, b := range buttons {
		for _, row := range b {
			A[row][col] = 1
		}
	}

	// Solve Ax = b
	var freeVars []int
	result, vals, matrix := gaussJordanSolve(A, b)
	if result == 0 {
		// There are free variables. Add new rows for them and try all solutions.
		freeVars = vals
		freeVals := make([]int, len(freeVars))
		result, vals = trySolution(0, len(vals), 200, freeVals, freeVars, matrix)
	}

	valStr := fmt.Sprintf("%v", vals)
	freeStr := fmt.Sprintf("%v", freeVars)
	fmt.Printf("%3d  vals: %-35s  free vars: %-10s ", result, valStr, freeStr)
	return result
}

func trySolution(level int, maxDepth int, iterations int, current []int, freeVars []int, matrix [][]float64) (int, []int) {
	// Base Case: We've reached the innermost level
	if level == maxDepth {
		result, vals := backSubstitute(matrix, freeVars, current)
		return result, vals
	}

	// The actual loop for this specific level
	result := math.MaxInt
	var vals []int
	for i := 0; i < iterations; i++ {
		current[level] = i // Store current loop's index

		// Recurse to the next level
		currentResult, currentVals := trySolution(level+1, maxDepth, iterations, current, freeVars, matrix)
		if currentResult < result && currentResult > 0 {
			result = currentResult
			vals = currentVals
		}
	}

	return result, vals
}

func gaussJordanSolve(A [][]int, b []int) (int, []int, [][]float64) {
	rows := len(A)
	if rows == 0 {
		return 0, nil, nil
	}

	// Build augmented matrix [A|b]
	cols := len(A[0])
	aug := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		if len(A[i]) != cols {
			return 0, nil, nil
		}

		aug[i] = make([]float64, cols+1)
		for j := 0; j < len(A[i]); j++ {
			aug[i][j] = float64(A[i][j])
		}

		aug[i][cols] = float64(b[i])
	}

	// Initial pivot rows
	pivotRowForCol := make([]int, cols)
	for i := range pivotRowForCol {
		pivotRowForCol[i] = -1
	}

	currentRow := 0 // current row for placing pivot
	for c := 0; c < cols && currentRow < rows; c++ {
		// Find best pivot in column c at or below row currentRow
		pivotRow := -1
		for r := currentRow; r < rows; r++ {
			if aug[r][c] != 0 {
				pivotRow = r
				break
			}
		}
		if pivotRow == -1 {
			// No pivot in this column; variable is free
			continue
		}

		// Swap into row currentRow
		aug[currentRow], aug[pivotRow] = aug[pivotRow], aug[currentRow]
		pivotRowForCol[c] = currentRow

		// Normalize current row
		pivotVal := aug[currentRow][c]
		for j := c; j <= cols; j++ {
			aug[currentRow][j] /= pivotVal
		}

		// Eliminate this column from other rows
		for r := 0; r < rows; r++ {
			if aug[r][c] != 0 {
				var factor float64
				if r != currentRow {
					factor = aug[r][c]
				}

				for j := c; j <= cols; j++ {
					aug[r][j] -= factor * aug[currentRow][j]
					if math.Abs(aug[r][j]) < .00001 {
						aug[r][j] = 0
					}
				}
			}
		}

		currentRow++
	}

	var freeVars []int
	x := make([]int, cols)
	invalid := false
	total := 0
	for c := 0; c < cols; c++ {
		pivotRow := pivotRowForCol[c]
		// A column without a pivot row is a free var
		if pivotRow == -1 {
			freeVars = append(freeVars, c)
			continue
		}

		// Our solution will only have integer results
		mod := math.Mod(aug[pivotRow][cols], 1.0)
		if 0.00001 < mod && mod < 0.99999 {
			invalid = true
			continue
		}

		// Our solution will also not have negative numbers
		x[c] = int(math.Round(aug[pivotRow][cols]))
		if x[c] < 0 {
			invalid = true
			continue
		}

		total += x[c]
	}

	if len(freeVars) != 0 {
		return 0, freeVars, aug
	}

	if invalid {
		return math.MaxInt, nil, nil
	}

	return total, x, nil
}

func backSubstitute(matrix [][]float64, freeVars []int, freeValues []int) (int, []int) {
	n := len(matrix[0]) - 1 // number of variables
	solution := make([]float64, n)

	// Set free variables
	for i, idx := range freeVars {
		solution[idx] = float64(freeValues[i])
	}

	// Back-substitute for pivot variables
	freeSet := make(map[int]bool)
	for _, idx := range freeVars {
		freeSet[idx] = true
	}

	row := 0
	for col := 0; col < n && row < len(matrix); col++ {
		if freeSet[col] {
			continue
		}
		// Find the pivot row for this column
		if matrix[row][col] == 1 {
			val := matrix[row][n]
			for j := col + 1; j < n; j++ {
				val -= matrix[row][j] * solution[j]
			}
			solution[col] = val
			row++
		}
	}

	var total int
	sol := make([]int, n)
	invalid := false
	for i, val := range solution {
		mod := math.Mod(val, 1.0)
		if 0.00001 < mod && mod < 0.99999 {
			invalid = true
			continue
		}

		// Our solution will also not have negative numbers
		sol[i] = int(math.Round(val))
		if sol[i] < 0 {
			invalid = true
			continue
		}

		total += sol[i]
	}

	if invalid {
		return math.MaxInt, nil
	}

	return total, sol
}
