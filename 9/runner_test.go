package main

import (
	"fmt"
	"testing"
)

func Test_find_lowest(t *testing.T) {
	input := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	result, points := find_lowest_risks(input)

	fmt.Println(points)

	if result != 15 {
		t.Errorf("Expected 15 and got %d", result)
	}

}

func Test_calc_basin(t *testing.T) {
	tubes := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	checked := make(map[int]map[int]bool)
	for i, _ := range tubes {
		checked[i] = make(map[int]bool)
	}

	result := calc_basin(tubes, checked, 0, 1)

	if result != 3 {
		t.Errorf("Expected 3 and got %d", result)
	}
}
