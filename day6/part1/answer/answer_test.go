package answer_test

import (
	"testing"

	"day6/part1/answer"
)

func TestCompute(t *testing.T) {
	numbers, operators := answer.Parse("sample.txt")
	expected := 4277556
	answer := answer.Compute(numbers, operators)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
