package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbers [25]int
	marks   [25]bool
}

func (b *Board) mark(number int) {
	for i := 0; i < 25; i++ {
		if b.numbers[i] == number {
			b.marks[i] = true
			break
		}
	}
}

func (b *Board) bingo() bool {
	// rows
	for i := 0; i < 25; i += 5 {
		marks := 0
		for j := 0; j < 5; j++ {
			if b.marks[i+j] == false {
				break
			} else {
				marks += 1
			}
		}
		if marks == 5 {
			return true
		}
	}

	// cols
	for i := 0; i < 5; i++ {
		marks := 0
		for j := 0; j < 25; j += 5 {
			if b.marks[i+j] {
				marks += 1
			} else {
				break
			}
		}
		if marks == 5 {
			return true
		}
	}
	return false
}

func (b *Board) score(lastcall int) int {
	total := 0
	for i := 0; i < 25; i++ {
		if b.marks[i] == false {
			total += b.numbers[i]
		}
	}
	return total * lastcall
}

func getcalls(line string) []int {
	calls := make([]int, 0)
	vals := strings.Split(line, ",")
	for _, val := range vals {
		call, _ := strconv.Atoi(val)
		calls = append(calls, call)

	}
	return calls
}

func makeboard(lines []string) *Board {
	b := Board{}
	for i := 0; i < 5; i++ {
		line := lines[i]
		vals := make([]int, 0)
		for j := 0; j < 15; j += 3 {
			val, _ := strconv.Atoi(strings.Trim(line[j:j+2], " "))
			vals = append(vals, val)
		}
		for j, val := range vals {
			b.numbers[(5*i)+j] = val
		}
	}
	return &b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	calls := getcalls(lines[0])
	boards := make([]*Board, 0)
	for i := 2; i < len(lines); i += 6 {
		board := makeboard(lines[i : i+5])
		boards = append(boards, board)
	}
	score := 0
	for _, call := range calls {
		for _, board := range boards {
			board.mark(call)
			if board.bingo() {
				score = board.score(call)
				break
			}
		}
		if score != 0 {
			break
		}
	}
	fmt.Println("1: ", score)

	//part 2
	won := make([]bool, len(boards))
	wins := 0
	score = 0
	for _, call := range calls {
		for i, board := range boards {
			if won[i] {
				continue
			}
			board.mark(call)
			if board.bingo() {
				won[i] = true
				wins += 1
				if wins == len(boards) {
					score = board.score(call)
				}
			}
		}
		if score != 0 {
			break
		}
	}
	fmt.Println("2: ", score)

}
