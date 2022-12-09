package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func clean_atoi(a string) int {
	val, _ := strconv.Atoi(a)
	return val
}

type cuboid struct {
	xspace [2]int
	yspace [2]int
	zspace [2]int
}

func (c *cuboid) split(i cuboid) int {
	return 0
	//return make([]cuboid{}, 0)

}

func main() {
	fmt.Println("x")
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), "\n")

	for _, line := range input {
		command := strings.Split(line, " ")
		onoff := command[0]
		fmt.Println(onoff)
		metrics := strings.Split(command[1], ",")
		fmt.Println("-------------", metrics)
		x := strings.Split(metrics[0], "=")[1]
		fmt.Println("--,0;sd", x)
		/*
			y := strings.Split(metrics[1], "..")
			z := strings.Split(metrics[2], "..")
			c := cuboid{[2]int{clean_atoi(x[0]), clean_atoi(x[1])}, [2]int{clean_atoi(y[0]), clean_atoi(y[1])}, [2]int{clean_atoi(z[0]), clean_atoi(z[1])}}
			fmt.Println(onoff, c)
		*/
	}
}
