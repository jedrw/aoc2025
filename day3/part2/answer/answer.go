package answer

import (
	"bufio"
	"os"
	"strconv"
)

func Parse(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func findJoltage(bank string, current string, length int) string {
	highest := 0
	highestIdx := 0
	for i := 0; i < len(bank)-(length-1); i++ {
		bat, _ := strconv.Atoi(string(bank[i]))
		if bat == 9 {
			highestIdx = i
			break
		}
		if bat > highest {
			highest = bat
			highestIdx = i
		}
	}

	current += string(bank[highestIdx])
	if length > 1 {
		remainder := bank[highestIdx+1:]
		return findJoltage(remainder, current, length-1)
	} else {
		return current
	}
}

func Compute(input []string) int {
	total := 0
	for _, line := range input {
		joltageStr := findJoltage(line, "", 12)
		joltage, _ := strconv.Atoi(joltageStr)
		total += joltage
	}

	return total
}
