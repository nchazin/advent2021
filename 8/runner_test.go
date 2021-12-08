package main

import (
	"testing"
)

func Test_line_1(t *testing.T) {
	input := "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"

	count := process_line1(input)

	if count != 3 {
		t.Errorf("Expected 3 and got %d", count)

	}
}

func Test_count_matches(t *testing.T) {
	a := "abdefg"
	b := "cf"

	result := count_matches(a, b)
	if result != 1 {
		t.Errorf("For 6  expected 1 and go: %d", result)
	}
}
