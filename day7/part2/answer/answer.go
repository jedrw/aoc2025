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

var grid [][]string
var memo [][]int
var W, H int

func findPaths(x, y int) int {
	if x < 0 || x >= W {
		return 0
	}
	if y == H-1 {
		return 1
	}
	if memo[y][x] != -1 {
		return memo[y][x]
	}

	total := 0

	switch grid[y+1][x] {
	case ".":
		total += findPaths(x, y+1)
	case "^":
		total += findPaths(x-1, y+1)
		total += findPaths(x+1, y+1)
	}

	memo[y][x] = total
	return total
}

func Compute(input [][]string) int {
	grid = input
	H = len(grid)
	W = len(grid[0])

	memo = make([][]int, H)
	for i := range memo {
		memo[i] = make([]int, W)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	for y, row := range grid {
		for x, sym := range row {
			if sym == "S" {
				return findPaths(x, y)
			}
		}
	}

	return 0
}
