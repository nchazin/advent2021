package main

import (
	"reflect"
	"testing"
)

func Test_clean_atoi(t *testing.T) {
	result := clean_atoi("10")

	if result != 10 {
		t.Errorf("Expected 10 got %d", result)
	}
}

func Test_fill_field_x(t *testing.T) {
	expected := map[Point]int{Point{1, 1}: 1, Point{1, 2}: 1, Point{1, 3}: 1}
	field := make(map[Point]int)
	fill_field(field, Point{1, 1}, Point{1, 3})

	if !reflect.DeepEqual(field, expected) {
		t.Errorf("Expected %v got %v", expected, field)
	}

	field = make(map[Point]int)
	fill_field(field, Point{1, 3}, Point{1, 1})
	if !reflect.DeepEqual(field, expected) {
		t.Errorf("Expected %v got %v", expected, field)
	}
}

func Test_fill_field_y(t *testing.T) {
	expected := map[Point]int{Point{1, 1}: 1, Point{2, 1}: 1, Point{3, 1}: 1}
	field := make(map[Point]int)
	fill_field(field, Point{1, 1}, Point{3, 1})

	if !reflect.DeepEqual(field, expected) {
		t.Errorf("Expected %v got %v", expected, field)
	}

	field = make(map[Point]int)
	fill_field(field, Point{3, 1}, Point{1, 1})
	if !reflect.DeepEqual(field, expected) {
		t.Errorf("Expected %v got %v", expected, field)
	}
}
