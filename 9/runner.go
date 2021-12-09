package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	scanner := bufio.NewScanner(file)

	tubes := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		tubeline := make([]int, 0)
		for _, char := range line {
			tubeline = append(tubeline, int(char)-48)
		}
		tubes = append(tubes, tubeline)
	}
	lowest, points := find_lowest_risks(tubes)
	fmt.Println(lowest)
	fmt.Println(find_basins(tubes, points))
}

type Point struct {
	x int
	y int
}

func calc_basin(tubes [][]int, checked map[int]map[int]bool, x, y int) int {
	if x < 0 || x > len(tubes)-1 || y < 0 || y > len(tubes[0])-1 {
		return 0
	}
	if checked[x][y] {
		return 0
	}
	if tubes[x][y] == 9 {
		return 0
	}
	checked[x][y] = true
	return 1 + (calc_basin(tubes, checked, x+1, y) +
		calc_basin(tubes, checked, x-1, y) +
		calc_basin(tubes, checked, x, y-1) +
		calc_basin(tubes, checked, x, y+1))
}

func find_basins(tubes [][]int, points []Point) int {
	basins := [3]int{0, 0, 0}
	checked := make(map[int]map[int]bool)
	for i, _ := range tubes {
		checked[i] = make(map[int]bool)
	}

	for _, p := range points {
		size := calc_basin(tubes, checked, p.x, p.y)
		for i, val := range basins {
			if size > val {
				basins[i] = size
				newval := val
				for j := i + 1; j < len(basins); j++ {
					if newval == 0 {
						break
					}
					oldval := basins[j]
					basins[j] = newval
					newval = oldval
				}
				break
			}
		}
	}

	val := 1
	for _, v := range basins {
		val *= v
	}
	return val
}

func find_lowest_risks(tubes [][]int) (int, []Point) {
	xlen := len(tubes)
	ylen := len(tubes[0])
	lovals := make([]int, 0)
	points := make([]Point, 0)
	for i, v := range tubes {
		for j, val := range v {
			if check_lowest(tubes, i, j, xlen, ylen) {
				lovals = append(lovals, val)
				points = append(points, Point{x: i, y: j})
			}
		}
	}
	sum := 0
	for _, v := range lovals {
		sum += v + 1
	}
	return sum, points
}

func check_lowest(tubes [][]int, x int, y int, xlen int, ylen int) bool {
	val := tubes[x][y]
	if x > 0 && tubes[x-1][y] <= val {
		return false
	}
	if x < xlen-1 && tubes[x+1][y] <= val {
		return false
	}
	if y > 0 && tubes[x][y-1] <= val {
		return false
	}
	if y < ylen-1 && tubes[x][y+1] <= val {
		return false
	}
	return true
}
