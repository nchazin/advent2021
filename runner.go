package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		var val int
		val, _ = strconv.Atoi(scanner.Text())
		numbers = append(numbers, val)
	}

	fmt.Println(calc_window(numbers, 1))
	fmt.Println(calc_window(numbers, 3))

}

func calc_window(numbers []int, size int) int {
	var count = 0
	for i := 0; i < len(numbers)-size; i++ {
		if sum(numbers[i+1:i+size+1]) > sum(numbers[i:i+size]) {
			count++
		}
	}
	return count
}

func sum(numbers []int) int {
	var sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
