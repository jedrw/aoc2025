package answer_test

import (
	"testing"

	"day5/part2/answer"
)

func TestCompute(t *testing.T) {
	ranges, ingredients := answer.Parse("sample.txt")
	expected := 14
	answer := answer.Compute(ranges, ingredients)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
