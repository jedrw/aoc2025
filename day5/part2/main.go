package main

import (
	"fmt"

	"day5/part2/answer"
)

func main() {
	ranges, ingredients := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(ranges, ingredients))
}
