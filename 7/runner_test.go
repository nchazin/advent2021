package main

import (
	"testing"
)

var locations = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func Test_calculate_min_fuel(t *testing.T) {
	fuel := calculate_min_fuel(locations)

	if fuel != 37 {
		t.Errorf("We expected 37 and got %d", fuel)
	}
}

func Test_median(t *testing.T) {
	result := calc_median(locations)

	if result != 2 {
		t.Errorf("We expected 2 and got %d", result)
	}

}

func Test_calculate_min_fuel_complicated(t *testing.T) {
	result := calculate_min_fuel_complicated(locations)

	if result != 168 {
		t.Errorf("We expected 168 and got %d", result)
	}
}

func Test_mean(t *testing.T) {
	result := calc_mean(locations)

	if result != 5 {
		t.Errorf("We expected 5  and got %d", result)
	}
}
