package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lens = map[int][]int{2: []int{1}, 3: []int{7}, 4: []int{4}, 7: []int{8}}

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

	const maxCapacity = 6889999 // your required line length
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, maxCapacity)
	count := 0
	count2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += process_line1(line)
		count2 += process_line2(line)
	}
	fmt.Println(count)
	fmt.Println(count2)
}

func process_line1(line string) int {
	count := 0
	parts := strings.Split(line, "|")
	outputs := strings.Fields(parts[1])
	for _, output := range outputs {
		//fmt.Println(len(output))
		if _, ok := lens[len(output)]; ok {
			//fmt.Println("Match on: ", output)
			count += 1
		}
	}
	return count
}

func process_line2(line string) int {
	parts := strings.Split(line, "|")
	patterns := strings.Fields(parts[0])
	outputs := strings.Fields(parts[1])
	chars := make(map[string]string)
	numbs := make(map[int]string)
	//pass 1 get the knowns
	for _, pattern := range patterns {
		if len(pattern) == 7 {
			chars[pattern] = "8"
			numbs[8] = pattern
		} else if len(pattern) == 4 {
			chars[pattern] = "4"
			numbs[4] = pattern
		} else if len(pattern) == 3 {
			chars[pattern] = "7"
			numbs[7] = pattern
		} else if len(pattern) == 2 {
			chars[pattern] = "1"
			numbs[1] = pattern
		}
	}
	//we can figure out the top piece...but
	//pass 2, we can resolve the 6 3:
	//6 has 6 parts and 1 match with 1
	//3 has 5 [arts and 2 matches with 1]
	for _, pattern := range patterns {
		if len(pattern) == 6 {
			matches := count_matches(pattern, numbs[1])
			if matches == 1 {
				chars[pattern] = "6"
				numbs[6] = pattern
			}
		} else if len(pattern) == 5 {
			matches := count_matches(pattern, numbs[1])
			if matches == 2 {
				chars[pattern] = "3"
				numbs[3] = pattern
			}
		}

	}
	// 0 has 4 matches with 3
	// 9 has 5
	// 6 has 5 - but we know 6 already!
	// do another pass
	// 5 has 5 matches with 6 and has 6 segments
	// 2 has 4 matched with 6 and has 6 segmetns
	for _, pattern := range patterns {
		if _, ok := chars[pattern]; ok {
			continue
		}
		if len(pattern) == 6 {
			matches := count_matches(pattern, numbs[3])
			if matches == 4 {
				chars[pattern] = "0"
				numbs[0] = pattern
			} else if matches == 5 {
				chars[pattern] = "9"
				numbs[9] = pattern
			}
		} else if len(pattern) == 5 {
			matches := count_matches(pattern, numbs[6])
			if matches == 5 {
				chars[pattern] = "5"
				numbs[5] = pattern
			} else if matches == 4 {
				chars[pattern] = "2"
				numbs[2] = pattern
			}
		}
	}
	return count_outputs(chars, outputs)
}

func count_outputs(patterns map[string]string, outputs []string) int {
	numeral := ""
	for _, output := range outputs {
		for k, v := range patterns {
			if len(k) != len(output) {
				continue
			}
			if count_matches(output, k) == len(output) {
				numeral += v
				break
			}
		}
	}
	return clean_atoi(numeral)
}

func count_matches(a string, b string) int {
	count := 0
	arunes := []rune(a)
	for _, rune := range arunes {
		if strings.Contains(b, string(rune)) {
			count += 1
		}
	}
	return count
}
