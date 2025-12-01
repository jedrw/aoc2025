package answer

import (
	"bufio"
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

type dial struct {
	position int
	clicks   int
}

func (d *dial) Turn(turn int) {
	var negTurn bool
	if turn < 0 {
		negTurn = true
		turn = -turn
	}

	for range turn {
		if negTurn {
			d.position--
		} else {
			d.position++
		}

		d.position %= 100
		if d.position < 0 {
			d.position += 100
		}

		if d.position == 0 {
			d.clicks++
		}
	}
}

func Compute(input []int) int {
	dial := dial{position: 50, clicks: 0}
	for _, turn := range input {
		dial.Turn(turn)
	}

	return dial.clicks
}
