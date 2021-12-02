package main

import (
	"testing"
)

func Test_get_movement_forward(t *testing.T) {
	command := "forward 10"
	dir, distance := get_movement(command)
	if (dir != "forward") || (distance != 10) {
		t.Errorf("Expected forward, 10 and got %s, %d", dir, distance)
	}
}

func Test_p1(t *testing.T) {
	commands := []string{"forward 5", "down 10", "up 5", "down 1", "forward 1"}
	result := p1(commands)

	if result != 36 {
		t.Errorf("Expected 36 got %d", result)
	}
}

func Test_p2(t *testing.T) {
	commands := []string{"forward 5", "down 10", "up 5", "down 1", "forward 3"}
	result := p2(commands)

	if result != 144 {
		t.Errorf("Expected 144 got %d", result)
	}

}
