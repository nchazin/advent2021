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

type Paper map[Point]bool

func clean_atoi(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

func fold_y(paper Paper, y int) Paper {
	for point, _ := range paper {
		if point.y > y {
			newpoint := Point{point.x, y - (point.y - y)}
			paper[newpoint] = true
			delete(paper, point)
		}
	}
	return paper
}

func fold_x(paper Paper, x int) Paper {
	for point, _ := range paper {
		if point.x > x {
			newpoint := Point{x - (point.x - x), point.y}
			paper[newpoint] = true
			delete(paper, point)
		}
	}
	return paper
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	paper := make(Paper)
	folds := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if len(line) > 4 && line[0:4] == "fold" {
			folds = append(folds, line[11:])
			continue
		}
		coords := strings.Split(line, ",")
		//fmt.Println(coords, line)
		point := Point{clean_atoi(coords[0]), clean_atoi(coords[1])}
		paper[point] = true
	}

	for i, fold := range folds {
		parts := strings.Split(fold, "=")
		//fmt.Println("-->", folds, parts)
		if parts[0] == "x" {
			paper = fold_x(paper, clean_atoi(parts[1]))
		} else {
			paper = fold_y(paper, clean_atoi(parts[1]))
		}
		if i == 0 {
			fmt.Println(len(paper))
		}
	}

	var paper2 [10][50]int

	for x, _ := range paper {
		paper2[x.y][x.x] = 1
	}
	for _, p := range paper2 {
		for _, d := range p {
			if d == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")

	}
}
