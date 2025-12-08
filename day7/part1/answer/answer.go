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
		line := scanner.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	return grid
}

func Compute(grid [][]string) int {
	splits := 0
	for y, row := range grid {
		for x, sym := range row {
			if sym == "S" || sym == "|" {
				if y == len(grid)-1 {
					continue
				}

				if grid[y+1][x] == "." {
					grid[y+1][x] = "|"
				}

				if grid[y+1][x] == "^" {
					splits++
					grid[y+1][x+1] = "|"
					grid[y+1][x-1] = "|"
				}
			}
		}
	}

	return splits
}
