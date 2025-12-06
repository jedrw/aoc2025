package answer_test

import (
	"testing"

	"day5/part1/answer"
)

func TestCompute(t *testing.T) {
	ranges, ingredients := answer.Parse("sample.txt")
	expected := 3
	answer := answer.Compute(ranges, ingredients)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
