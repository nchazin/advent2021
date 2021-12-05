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

func fill_field_diagonal(field map[Point]int, a Point, b Point) {
	y := a.y
	if a.x < b.x {
		if a.y < b.y {
			for x := a.x; x <= b.x; x++ {
				field[Point{x, y}]++
				y++
			}
		} else {
			for x := a.x; x <= b.x; x++ {
				field[Point{x, y}]++
				y--
			}
		}
	} else {
		if a.y < b.y {
			for x := a.x; x >= b.x; x-- {
				field[Point{x, y}]++
				y++
			}
		} else {
			for x := a.x; x >= b.x; x-- {
				field[Point{x, y}]++
				y--
			}
		}
	}
}

func fill_field(field map[Point]int, a Point, b Point, diagonal bool) {
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
	} else if diagonal {
		fill_field_diagonal(field, a, b)
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

	fieldp1 := make(map[Point]int)
	fieldp2 := make(map[Point]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		startstr := strings.Split(parts[0], ",")
		endstr := strings.Split(parts[1], ",")
		start := Point{clean_atoi(startstr[0]), clean_atoi(startstr[1])}
		end := Point{clean_atoi(endstr[0]), clean_atoi(endstr[1])}
		fill_field(fieldp1, start, end, false)
		fill_field(fieldp2, start, end, true)
	}

	count := 0
	for _, p := range fieldp1 {
		if p > 1 {
			count++
		}
	}
	fmt.Println(count)

	count = 0
	for _, p := range fieldp2 {
		if p > 1 {
			count++
		}
	}
	fmt.Println(count)
}
