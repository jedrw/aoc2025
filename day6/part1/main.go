package main

import (
	"fmt"

	"day6/part1/answer"
)

func main() {
	numbers, operators := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(numbers, operators))
}
