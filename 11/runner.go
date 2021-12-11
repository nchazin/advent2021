package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type OctoGrid [10][10]int
type FlashGrid [10][10]bool

func flash(octopus *OctoGrid, flashes *FlashGrid, i int, j int) {
	if i < 0 || i >= 10 || j < 0 || j >= 10 {
		return
	}
	if flashes[i][j] {
		return
	}

	flashes[i][j] = true
	octopus[i][j] = 0

	for x := i - 1; x <= +i+1; x++ {
		for y := j - 1; y <= +j+1; y++ {
			if x < 0 || x >= 10 || y < 0 || y >= 10 || (x == i && y == j) {
				continue
			}
			//fmt.Printf("energizing %d %d from %d %d\n", x, y, i, j)
			octopus[x][y]++
			if octopus[x][y] > 9 {
				flash(octopus, flashes, x, y)
			}
		}
	}
}

func printo(octopus *OctoGrid) {
	for i := 0; i < 10; i++ {
		fmt.Println(octopus[i])
	}
}

func energize(octopus *OctoGrid) int {
	var flashed FlashGrid
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			octopus[i][j] += 1
			if octopus[i][j] > 9 {
				flash(octopus, &flashed, i, j)
			}
		}
	}

	flashes := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if flashed[i][j] {
				octopus[i][j] = 0
				flashes++
			}
		}
	}
	return flashes
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dumbo OctoGrid
	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		for col, c := range line {
			dumbo[row][col] = int(c - '0')
		}
		row++
	}
	dumbo2 := dumbo

	flashes := 0
	for i := 1; i <= 100; i++ {
		flashed := energize(&dumbo)
		flashes += flashed
	}
	fmt.Println(flashes)
	printo(&dumbo)
	fmt.Println("")
	printo(&dumbo2)
	step := 1
	for {
		flashed := energize(&dumbo2)
		if flashed == 100 {
			break
		}
		step++
	}
	fmt.Println(step)

}
