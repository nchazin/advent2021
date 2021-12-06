package main

import (
	"testing"
)

var input = map[int]int{3: 2, 4: 1, 2: 1, 1: 1}

func Test_18(t *testing.T) {
	fishschool := makefish(input, 18)
	total_fish := countfish(fishschool)
	if total_fish != 26 {
		t.Errorf("Expected 26 and got: %d", total_fish)
	}
}

func Test_80(t *testing.T) {
	fishschool := makefish(input, 80)
	total_fish := countfish(fishschool)
	if total_fish != 5934 {
		t.Errorf("Expected 5934 and got: %d", total_fish)
	}
}
