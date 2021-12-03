package main

import (
	"testing"
)

func Test_p1(t *testing.T) {
	commands := []string{"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
	result := p1(commands)

	if result != 198 {
		t.Errorf("Expected 198 got %d", result)
	}
}
