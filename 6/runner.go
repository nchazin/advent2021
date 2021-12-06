package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func makefish(fishschool map[int]int, days int) map[int]int {
	for day := 1; day <= days; day++ {
		newschool := make(map[int]int)
		for timer, count := range fishschool {
			if timer == 0 {
				newschool[6] += count
				newschool[8] += count
			} else {
				newschool[timer-1] += count
			}

		}
		fishschool = newschool
	}

	return fishschool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fishschool := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fishes := strings.Split(scanner.Text(), ",")
		for _, fish := range fishes {
			fishtimer, _ := strconv.Atoi(fish)
			fishschool[fishtimer]++
		}
	}

	fmt.Println(countfish(makefish(fishschool, 80)))
	fmt.Println(countfish(makefish(fishschool, 256)))

}

func countfish(fishschool map[int]int) int {
	size := 0

	for _, count := range fishschool {
		size += count
	}
	return size
}
