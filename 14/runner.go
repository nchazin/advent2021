package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Digraph map[string]byte

type Insert struct {
	index uint
	value byte
}

func insert(a []byte, index uint, value byte) []byte {
	if uint(len(a)) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func step(pattern []byte, digraphs Digraph) []byte {
	inserts := make([]Insert, 0)

	for i := 0; i < len(pattern)-1; i++ {
		d := string(pattern[i : i+2])
		inserts = append(inserts, Insert{uint(2*i + 1), digraphs[d]})
	}

	for _, isrt := range inserts {
		pattern = insert(pattern, isrt.index, isrt.value)
	}
	return pattern
}

func count_chars(pattern []byte) map[byte]int {
	counts := make(map[byte]int)
	for _, c := range pattern {
		counts[c]++
	}
	return counts
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	digraphs := make(Digraph)
	var pattern []byte
	for scanner.Scan() {
		if index == 0 {
			pattern = scanner.Bytes()
			index++
			continue
		}
		if index == 1 {
			index++
			continue
		}
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		digraphs[parts[0]] = parts[1][0]
		index++
	}

	for i := 1; i <= 10; i++ {
		pattern = step(pattern, digraphs)
		//fmt.Println("pattern at step ", i, ":", string(pattern[0:len(pattern)]))
		fmt.Println(i, " - ", len(pattern))
	}
	counts := count_chars(pattern)
	maxv := 0
	minv := len(pattern)
	for _, v := range counts {
		if v > maxv {
			maxv = v
		}
		if v < minv {
			minv = v
		}
	}
	fmt.Println(maxv - minv)
}
