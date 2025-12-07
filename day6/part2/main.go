package main

import (
	"fmt"

	"day6/part2/answer"
)

func main() {
	numbers, operators := answer.Parse("../input.txt")
	fmt.Println(answer.Compute(numbers, operators))
}
