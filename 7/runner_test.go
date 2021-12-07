package main

import (
	"testing"
)

func Test_calculate_min_fuel(t *testing.T) {
	locations := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fuel := calculate_min_fuel(locations)

	if fuel != 37 {
		t.Errorf("We expected 37 and got %d", fuel)
	}
}

func Test_median(t *testing.T) {
	locations := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	result := calc_median(locations)

	if result != 2 {
		t.Errorf("We expected 2 and got %d", result)
	}

}
