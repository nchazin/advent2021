package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Digraph map[string]byte

type Pattern map[string]int

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

func step_patternmap(patternmap Pattern, digraphs Digraph) Pattern {
	inserted := make(Pattern)
	// for instance we have NN -> B
	// So we have NBN, and we now add an NB and an BN
	for k, v := range patternmap {
		var newv [2]byte
		newc := digraphs[string(k)]
		newv[0] = k[0]
		newv[1] = newc
		inserted[string(newv[:])] += v
		newv[0] = newc
		newv[1] = k[1]
		inserted[string(newv[:])] += v
	}

	return inserted
}

func count_chars(pattern []byte) map[byte]int {
	counts := make(map[byte]int)
	for _, c := range pattern {
		counts[c]++
	}
	return counts
}

func count_chars2(patternmap Pattern) map[byte]uint64 {
	counts := make(map[byte]uint64)
	for dg, c := range patternmap {
		counts[dg[0]] += uint64(c)
		counts[dg[1]] += uint64(c)
	}
	//We double because say we have NBNC
	//We will count the NB 1 1, then the B N 1 1
	//dp 2 N and 2 B - we will need to add the first
	//char once more at teh end...
	for b, count := range counts {
		counts[b] = count / 2
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
	patternmap := make(Pattern)
	var pattern []byte
	for scanner.Scan() {
		if index == 0 {
			pattern = scanner.Bytes()
			for i := 0; i <= len(pattern)-2; i++ {
				patternmap[string(pattern[i:i+2])]++
			}

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
	fmt.Println(patternmap)

	for i := 1; i <= 10; i++ {
		pattern = step(pattern, digraphs)
	}
	counts := count_chars(pattern)
	for k, v := range counts {
		fmt.Println(string(k), " -> ", v)
	}
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

	for i := 1; i <= 40; i++ {
		patternmap = step_patternmap(patternmap, digraphs)
	}

	counts2 := count_chars2(patternmap)
	// dividing by 2 we undercount the first and last characters, which are fixed in place!
	counts2[pattern[0]] += 1
	counts2[pattern[len(pattern)-1]] += 1
	for k, v := range counts2 {
		fmt.Println(string(k), " -> ", v)
	}

	maxv2 := uint64(0)
	minv2 := uint64(math.MaxUint64)
	for _, v := range counts2 {
		if v > maxv2 {
			maxv2 = v
		}
		if v < minv2 {
			minv2 = v
		}
	}
	fmt.Println(maxv2 - minv2)
}
