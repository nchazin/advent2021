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
/*
	var numbers []int

	for scanner.Scan() {
		var val int
		val, _ = strconv.Atoi(scanner.Text())
		numbers = append(numbers, val)
	}
*/

}

/*
func pass() (str, int) {
	return "foo", 3
}
*/

/*
func get_movement(command str) {
	tokens := strings.Fields(command)

        dir := tokens[0]
        distance, _ := strconv(Atoi(tokens[1])
	fmt.Println(dir, distance)
}
*/
