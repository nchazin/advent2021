package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var commands []string
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	fmt.Println(p1(commands))
	fmt.Println(p2(commands))

}

func p1(commands []string) int {
	curx, cury := 0, 0
	for _, command := range commands {
		dir, distance := get_movement(command)
		if dir == "forward" {
			curx += distance
		} else if dir == "up" {
			cury -= distance
		} else if dir == "down" {
			cury += distance
		}

	}
	fmt.Println("Curx ", curx, " Cury ", cury)
	return curx * cury
}

func p2(commands []string) int {
	curx, cury, aim := 0, 0, 0
	for _, command := range commands {
		dir, distance := get_movement(command)
		if dir == "forward" {
			curx += distance
			cury += aim * distance
		} else if dir == "up" {
			aim -= distance
		} else if dir == "down" {
			aim += distance
		}

	}
	fmt.Println("Curx ", curx, " Cury ", cury)
	return curx * cury
}

func get_movement(command string) (string, int) {
	var distance int
	tokens := strings.Fields(command)

	dir := tokens[0]

	distance, _ = strconv.Atoi(tokens[1])
	return dir, distance
}
