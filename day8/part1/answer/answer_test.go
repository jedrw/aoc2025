package answer_test

import (
	"testing"

	"day8/part1/answer"
)

func TestCompute(t *testing.T) {
	input := answer.Parse("sample.txt")
	expected := 40
	answer := answer.Compute(input, 10)

	if answer != expected {
		t.Errorf("expected %d, got %d", expected, answer)
	}
}
