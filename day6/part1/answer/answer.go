package answer

import (
	"os"
	"strconv"
	"strings"
)

func Parse(filePath string) ([][]int, []string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	operatorLine := lines[len(lines)-1]
	operators := strings.Fields(operatorLine)

	numberLines := lines[:len(lines)-1]
	questions := [][]int{}
	for _, line := range numberLines {
		numberStrings := strings.Fields(line)
		for len(questions) < len(numberStrings) {
			questions = append(questions, []int{})
		}

		for j, n := range numberStrings {
			num, _ := strconv.Atoi(n)
			questions[j] = append(questions[j], num)
		}
	}

	return questions, operators
}

func Compute(questions [][]int, operators []string) int {
	finalAnswer := 0
	for i, q := range questions {
		answer := q[0]
		operator := operators[i]
		switch operator {
		case "+":
			for _, n := range q[1:] {
				answer += n
			}
		case "*":
			for _, n := range q[1:] {
				answer *= n
			}
		}

		finalAnswer += answer
	}

	return finalAnswer
}
