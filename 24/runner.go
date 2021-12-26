package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

/*
The ALU is a four-dimensional processing unit: it has integer variables w, x, y, and z.
These variables all start with the value 0. The ALU also supports six instructions:

inp a - Read an input value and write it to variable a.
add a b - Add the value of a to the value of b, then store the result in variable a.
mul a b - Multiply the value of a by the value of b, then store the result in variable a.
div a b - Divide the value of a by the value of b, truncate the result to an integer, then store
   the result in variable a. (Here, "truncate" means to round the value toward zero.)
mod a b - Divide the value of a by the value of b, then store the remainder in variable a.
   (This is also called the modulo operation.)
eql a b - If the value of a and b are equal, then store the value 1 in variable a. Otherwise,
   store the value 0 in variable a.
In all of these instructions, a and b are placeholders; a will always be the variable
  where the result of the operation is stored (one of w, x, y, or z),
  while b can be either a variable or a number. Numbers can be positive or negative,
   but will always be integers.

The ALU has no jump instructions; in an ALU program, every instruction is run exactly once in order
 from top to bottom. The program halts after the last instruction has finished executing.
*/

func clean_atoi(a string) int64 {
	val, _ := strconv.Atoi(a)
	return int64(val)
}

type Alu struct {
	w, x, y, z int64
}

func (a *Alu) imp(r byte, v int64) {
	switch r {
	case 'w':
		a.w = v
	case 'x':
		a.x = v
	case 'y':
		a.y = v
	case 'z':
		a.z = v
	default:
		fmt.Errorf("Unable to fetch %d in %c\n", v, a)
	}
}

func (a *Alu) store(r byte, v int64) {
	a.imp(r, v)
}

func (a *Alu) fetch(r byte) int64 {
	switch r {
	case 'w':
		return a.w
	case 'x':
		return a.x
	case 'y':
		return a.y
	case 'z':
		return a.z
	}
	return math.MaxInt64
}

func is_register(b byte) bool {
	if b == 'w' || b == 'x' || b == 'y' || b == 'z' {
		return true
	}
	return false
}

func (alu *Alu) operand(b []byte) int64 {
	var bval int64
	if is_register(b[0]) {
		bval = alu.fetch(b[0])
	} else {
		bval = clean_atoi(string(b[:]))
	}
	return bval
}

func (alu *Alu) add(a byte, b []byte) {
	bval := alu.operand(b)
	aval := alu.fetch(a)
	alu.store(a, aval+bval)
}

func (alu *Alu) mul(a byte, b []byte) {
	bval := alu.operand(b)
	aval := alu.fetch(a)
	alu.store(a, aval*bval)

}

func (alu *Alu) div(a byte, b []byte) {
	bval := alu.operand(b)
	aval := alu.fetch(a)
	alu.store(a, aval/bval)
}

func (alu *Alu) mod(a byte, b []byte) {
	bval := alu.operand(b)
	aval := alu.fetch(a)
	alu.store(a, aval%bval)
}

func (alu *Alu) eq(a byte, b []byte) {
	bval := alu.operand(b)
	aval := alu.fetch(a)
	if aval == bval {
		alu.store(a, int64(1))

	} else {
		alu.store(a, int64(0))
	}
}

func (a Alu) get_all() [4]int64 {
	ret := [4]int64{a.w, a.x, a.y, a.z}
	return ret
}

var input = []int64{3, 9, 9, 9, 9, 6, 9, 8, 7, 9, 9, 4, 2, 9}
var current = 0

func get_input() int64 {
	val := input[current]
	current++
	return val
}

func main() {
	fmt.Println("We did it")
	flag.Parse()
	fmt.Println("tail:", flag.Args())
	input_file := flag.Args()[0]
	fmt.Println("tail:", input_file)
	data, err := ioutil.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(data), "\n")

	alu := Alu{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		if parts[0] == "" {
			continue
		}

		switch parts[0] {
		case "inp":
			value := get_input()
			fmt.Println(value)
			alu.imp([]byte(parts[1])[0], value)
		case "add":
			alu.add([]byte(parts[1])[0], []byte(parts[2]))
		case "mul":
			alu.mul([]byte(parts[1])[0], []byte(parts[2]))
		case "div":
			alu.div([]byte(parts[1])[0], []byte(parts[2]))
		case "mod":
			alu.mod([]byte(parts[1])[0], []byte(parts[2]))
		case "eql":
			alu.eq([]byte(parts[1])[0], []byte(parts[2]))
		default:
			fmt.Println("Unpected operator: [", parts, "]")
		}
	}
	fmt.Println("Output: ", alu.get_all())
}
