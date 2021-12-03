package main

import (
	"reflect"
	"testing"
)

var input = []string{"00100",
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

func Test_p1(t *testing.T) {
	commands := input
	result := p1(commands)

	if result != 198 {
		t.Errorf("Expected 198 got %d", result)
	}
}

func Test_p2(t *testing.T) {
	commands := input
	result := p2(commands)

	if result != 230 {
		t.Errorf("Expected 239 got %d", result)
	}
}

func Test_counters(t *testing.T) {
	zeros_expected := []int{5, 7, 4, 5, 7}
	ones_expected := []int{7, 5, 8, 7, 5}

	zeros, ones := counters(input)
	if !reflect.DeepEqual(zeros, zeros_expected) || !reflect.DeepEqual(ones, ones_expected) {
		t.Errorf("Expected %d %d got %d %d", zeros_expected, ones_expected, zeros, ones)
	}
}
