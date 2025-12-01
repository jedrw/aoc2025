package answer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Parse(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	turns := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		turnString := scanner.Text()
		turnDir := string(turnString[0])
		turn, err := strconv.Atoi(turnString[1:])
		if err != nil {
			panic(err)
		}

		if turnDir == "L" {
			turn = -turn
		}

		turns = append(turns, turn)
	}

	return turns
}

func Compute(input []int) int {
	dial := 50
	zeros := 0
	for _, turn := range input {
		dial += turn
		fmt.Println(dial)
		dial = dial % 100
		if dial == 0 {
			zeros++
		}
	}

	return zeros
}
