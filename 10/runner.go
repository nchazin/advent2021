package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

/*
): 3 points.
]: 57 points.
}: 1197 points.
>: 25137 points.
*/
var closers = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var openers = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var fixers = map[byte]uint{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func score_incomplete_lines(line []byte) uint {
	stack := make([]byte, 0)

	for _, c := range line {
		// assume input is safe - no extra chars, and always an opener first!
		if val, ok := openers[c]; ok {
			stack = append(stack, val)
		} else if _, ok := closers[c]; ok {
			if stack[len(stack)-1] != c {
				return 0
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			fmt.Errorf("Unexpected input: %c", c)
		}
	}
	var score uint

	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += fixers[stack[i]]
	}
	return score
}

func score_line(line []byte) int {
	stack := make([]byte, 0)

	for _, c := range line {
		// assume input is safe - no extra chars, and always an opener first!
		if val, ok := openers[c]; ok {
			stack = append(stack, val)
		} else if points, ok := closers[c]; ok {
			if stack[len(stack)-1] != c {
				return points
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			fmt.Errorf("Unexpected input: %c", c)
		}
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := 0
	score_incomplete := make([]uint, 0)

	for scanner.Scan() {
		line := scanner.Bytes()
		score += score_line(line)
		output := score_incomplete_lines(line)
		if output != 0 {
			score_incomplete = append(score_incomplete, output)
		}
	}
	fmt.Println(score)
	sort.Slice(score_incomplete, func(i, j int) bool { return score_incomplete[i] < score_incomplete[j] })
	fmt.Println(score_incomplete[len(score_incomplete)/2])
}
