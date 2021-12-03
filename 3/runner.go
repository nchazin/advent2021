package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	//	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []string
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	fmt.Println(p1(numbers))
	fmt.Println(p2(numbers))

}

func p1(numbers []string) int64 {
	zeros, ones := counters(numbers)

	gamma := ""
	epsilon := ""
	bits := len([]rune(numbers[0]))
	for i := 0; i < bits; i++ {
		if zeros[i] > ones[i] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"

		}
	}
	fmt.Println(gamma, " --- ", epsilon)
	gammaval, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonval, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaval, " --- ", epsilonval)
	return gammaval * epsilonval
}

func counters(numbers []string) ([]int, []int) {
	bits := len([]rune(numbers[0]))

	var zeros []int
	var ones []int

	for i := 0; i < bits; i++ {
		zeros = append(zeros, 0)
		ones = append(ones, 0)
	}

	for _, number := range numbers {
		for i, c := range number {
			if c == '1' {
				ones[i] += 1
			} else {
				zeros[i] += 1
			}
		}

	}

	return zeros, ones
}

func p2(numbers []string) int64 {

	bits := len([]rune(numbers[0]))

	oldnumbers := numbers

	var zeros []int
	var ones []int
	//oxygen
	for i := 0; i < bits; i++ {
		zeros, ones = counters(oldnumbers)
		var newnumbers []string

		var mask byte
		if ones[i] >= zeros[i] {
			mask = '1'
		} else {
			mask = '0'
		}
		for _, number := range oldnumbers {
			if number[i] == mask {
				newnumbers = append(newnumbers, number)
			}
		}
		oldnumbers = newnumbers
		if len(oldnumbers) == 1 {
			break
		}
	}
	oxygen, _ := strconv.ParseInt(oldnumbers[0], 2, 64)

	//co2
	oldnumbers = numbers
	for i := 0; i < bits; i++ {
		zeros, ones = counters(oldnumbers)
		var newnumbers []string

		var mask byte
		if zeros[i] <= ones[i] {
			mask = '0'
		} else {
			mask = '1'
		}
		for _, number := range oldnumbers {
			if number[i] == mask {
				newnumbers = append(newnumbers, number)
			}
		}
		oldnumbers = newnumbers
		if len(oldnumbers) == 1 {
			break
		}
	}

	co2, _ := strconv.ParseInt(oldnumbers[0], 2, 64)

	return oxygen * co2

}

/*
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
*/
