package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func fill_field(field map[Point]int, a Point, b Point) {
	if a.x == b.x {
		if a.y < b.y {

			for y := a.y; y <= b.y; y++ {
				field[Point{a.x, y}]++
			}
		} else {
			for y := a.y; y >= b.y; y-- {
				field[Point{a.x, y}]++
			}
		}
	} else if a.y == b.y {
		if a.x < b.x {
			for x := a.x; x <= b.x; x++ {
				field[Point{x, a.y}]++
			}
		} else {
			for x := a.x; x >= b.x; x-- {
				field[Point{x, a.y}]++
			}

		}
	}
}

func clean_atoi(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	field := make(map[Point]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		startstr := strings.Split(parts[0], ",")
		endstr := strings.Split(parts[1], ",")
		start := Point{clean_atoi(startstr[0]), clean_atoi(startstr[1])}
		end := Point{clean_atoi(endstr[0]), clean_atoi(endstr[1])}
		fill_field(field, start, end)
	}

	count := 0
	for _, p := range field {
		if p > 1 {
			count++
		}
	}
	fmt.Println(count)
}
