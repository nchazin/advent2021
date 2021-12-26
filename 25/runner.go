package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Field map[Coord]byte

type Trench struct {
	rows  int
	cols  int
	field Field
}

func (t Trench) print() {
	for row := 0; row < t.rows; row++ {
		for col := 0; col < t.cols; col++ {
			if val, ok := t.field[Coord{row, col}]; ok {
				fmt.Printf("%c", val)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (t *Trench) stepr() bool {
	new_field := Field{}
	moved := false
	for row := 0; row < t.rows; row++ {
		for col := 0; col < t.cols; col++ {
			coord := Coord{row, col}
			if val, ok := t.field[Coord{row, col}]; ok {
				if val == 'v' {
					new_field[coord] = 'v'
				} else {
					destcoord := Coord{row, (col + 1) % t.cols}
					if _, ok := t.field[destcoord]; ok {
						new_field[coord] = '>'
					} else {
						new_field[destcoord] = '>'
						moved = true
					}
				}
			}
		}
	}
	t.field = new_field
	return moved
}

func (t *Trench) stepd() bool {
	new_field := Field{}
	moved := false
	for row := 0; row < t.rows; row++ {
		for col := 0; col < t.cols; col++ {
			coord := Coord{row, col}
			if val, ok := t.field[Coord{row, col}]; ok {
				if val == '>' {
					new_field[coord] = '>'
				} else {
					destcoord := Coord{(row + 1) % t.rows, col}
					if _, ok := t.field[destcoord]; ok {
						new_field[coord] = 'v'
					} else {
						new_field[destcoord] = 'v'
						moved = true
					}
				}
			}
		}
	}
	t.field = new_field
	return moved
}

func (t *Trench) step() bool {
	moved := false
	if t.stepr() {
		moved = true
	}
	if t.stepd() {
		moved = true
	}
	return moved
}

type Coord struct {
	x, y int
}

func main() {
	flag.Parse()
	input_file := flag.Args()[0]
	data, err := ioutil.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(data), "\n")

	rows := len(input)
	cols := 0
	trench := Trench{}
	field := Field{}
	for row, line := range input {
		if len(line) > 0 {
			cols = len(line)
		} else {
			rows--
		}
		for col, val := range line {
			switch val {
			case '>':
				field[Coord{row, col}] = '>'
			case 'v':
				field[Coord{row, col}] = 'v'
			}
		}
	}

	trench.rows = rows
	trench.cols = cols
	trench.field = field
	steps := 1
	for {
		if !trench.step() {
			break
		}
		steps += 1
		// just in case
		if steps > 10000 {
			break
		}
	}
	trench.print()
	fmt.Printf("We took %d steps\n", steps)

}
