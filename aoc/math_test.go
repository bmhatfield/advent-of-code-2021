package aoc

import "testing"

func TestSummation(t *testing.T) {
	s := Summation(5)

	expected := 1 + 2 + 3 + 4 + 5
	if s != expected {
		t.Error("expected", expected, "got", s)
	}

	s2 := Summation(12)

	expected2 := 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 + 11 + 12
	if s2 != expected2 {
		t.Error("expected", expected, "got", s2)
	}
}
