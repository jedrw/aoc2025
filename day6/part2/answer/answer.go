package answer

import (
	"os"
	"strconv"
	"strings"
)

func Parse(path string) ([][]int, []string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	operatorLine := lines[len(lines)-1]
	operators := strings.Fields(operatorLine)

	grid := lines[:len(lines)-1]
	maxWidth := 0
	for _, row := range grid {
		if len(row) > maxWidth {
			maxWidth = len(row)
		}
	}

	for i := range grid {
		if len(grid[i]) < maxWidth {
			grid[i] += strings.Repeat(" ", maxWidth-len(grid[i]))
		}
	}

	problems := [][]int{}
	var current []int

	isBlankColumn := func(col int) bool {
		for _, row := range grid {
			if row[col] != ' ' {
				return false
			}
		}
		return true
	}

	for col := maxWidth - 1; col >= 0; col-- {
		if isBlankColumn(col) {
			if len(current) > 0 {
				for i, j := 0, len(current)-1; i < j; i, j = i+1, j-1 {
					current[i], current[j] = current[j], current[i]
				}
				problems = append(problems, current)
				current = nil
			}
			continue
		}

		numStr := ""
		for _, row := range grid {
			if row[col] != ' ' {
				numStr += string(row[col])
			}
		}

		n, _ := strconv.Atoi(numStr)
		current = append(current, n)
	}

	if len(current) > 0 {
		for i, j := 0, len(current)-1; i < j; i, j = i+1, j-1 {
			current[i], current[j] = current[j], current[i]
		}
		problems = append(problems, current)
	}

	return problems, operators
}

func Compute(questions [][]int, operators []string) int {
	total := 0

	for i := range questions {
		q := questions[i]
		op := operators[len(operators)-1-i]

		acc := q[0]
		switch op {
		case "+":
			for _, v := range q[1:] {
				acc += v
			}
		case "*":
			for _, v := range q[1:] {
				acc *= v
			}
		}
		total += acc
	}

	return total
}
