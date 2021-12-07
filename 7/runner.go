package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func clean_atoi(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

func calc_median(locs []int) int {
	l := len(locs)
	var median int
	sort.Ints(locs)
	if l%2 == 0 {
		median = (locs[l/2-1] + locs[l/2]) / 2
	} else {
		median = locs[l/2]
	}
	return median
}

func int_abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func calculate_min_fuel(locs []int) int {
	median := calc_median(locs)

	fuel := 0
	for _, loc := range locs {
		fuel += int_abs(loc - median)
	}
	return fuel
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	locs := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		locations := strings.Split(scanner.Text(), ",")
		for _, loc := range locations {
			locs = append(locs, clean_atoi(loc))
		}
	}
	fmt.Println(calculate_min_fuel(locs))
}
