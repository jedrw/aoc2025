package answer

import (
	"bufio"
	"os"
	"strings"
)

func Parse(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	grid := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineString := scanner.Text()
		line := strings.Split(lineString, "")
		grid = append(grid, line)
	}

	return grid
}

type Position struct {
	x int
	y int
}

var allAdjacent = []Position{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func check4AdjacentFree(grid [][]string, pos Position) bool {
	free := 0
	for _, adj := range allAdjacent {
		x := pos.x + adj.x
		if x < 0 || x >= len(grid[0]) {
			free++
			continue
		}

		y := pos.y + adj.y
		if y < 0 || y >= len(grid) {
			free++
			continue
		}

		pos := grid[y][x]
		if pos == "." {
			free++
		}
	}

	return free >= 5
}

func Compute(input [][]string) int {
	free := 0
	for y, row := range input {
		for x, sym := range row {
			if sym == "@" {
				if check4AdjacentFree(input, Position{x, y}) {
					free++
				}
			}
		}

	}

	return free
}
