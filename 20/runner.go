package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type ImageMap struct {
	image  map[Coord]bool
	minrow int
	maxrow int
	mincol int
	maxcol int
}

func (im *ImageMap) print() {
	for i := im.minrow - 5; i < im.maxrow+5; i++ {
		for j := im.mincol - 5; j < im.maxcol+5; j++ {
			c := Coord{i, j}
			if im.image[c] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")

	}
}

func (im ImageMap) light() int {
	light := 0
	for _, v := range im.image {
		if v {
			light++
		}
	}
	return light
}

func (im ImageMap) enhance(key [512]byte, tick int) ImageMap {
	newimage := ImageMap{}
	newimage.image = make(map[Coord]bool)
	// infinite fill padding...
	var defaultfill byte

	// avoid the flapping of all 0s and all 1s in the evil input
	if tick%2 == 1 {
		if key[0] == byte('.') && key[511] == byte('#') {
			defaultfill = '1'
		} else {
			defaultfill = '0'
		}
	} else {
		if key[0] == byte('#') {
			defaultfill = '1'
		} else {
			defaultfill = '0'
		}
	}
	for i := im.minrow - 5; i < im.maxrow+5; i++ {
		for j := im.mincol - 5; j < im.maxcol+5; j++ {
			enhancer := make([]byte, 0)
			for a := i - 1; a <= i+1; a++ {
				for b := j - 1; b <= j+1; b++ {
					if val, ok := im.image[Coord{a, b}]; ok {
						if val {
							enhancer = append(enhancer, '1')
						} else {
							enhancer = append(enhancer, '0')
						}
					} else {
						enhancer = append(enhancer, defaultfill)
					}
				}
			}
			lookup, _ := strconv.ParseInt(string(enhancer[:]), 2, 32)
			if key[lookup] == '#' {
				newimage.image[Coord{i, j}] = true
			} else {
				newimage.image[Coord{i, j}] = false
			}
		}
	}
	newimage.mincol = im.mincol - 1
	newimage.maxcol = im.maxcol + 1
	newimage.minrow = im.minrow - 1
	newimage.maxrow = im.maxrow + 1
	return newimage
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var algorithm [512]byte
	input := strings.Split(string(data), "\n")

	for i, c := range input[0] {
		algorithm[i] = byte(c)
	}

	image := make(map[Coord]bool)

	rows := 0
	cols := 0
	for i := 1; i < len(input); i++ {
		if len(input[i]) == 0 {
			continue
		}
		cols = len(input[i])
		rows += 1
		for j, c := range input[i] {
			if c == '#' {
				image[Coord{i - 2, j}] = true
			} else if c == '.' {
				image[Coord{i - 2, j}] = false
			} else {
				fmt.Println("unexected char! ", c)
			}
		}
	}

	imagemap := ImageMap{image, 0, rows, 0, cols}
	fmt.Println(imagemap.light())
	for ticks := 1; ticks <= 50; ticks++ {
		imagemap = imagemap.enhance(algorithm, ticks)
		if ticks == 2 {
			//			imagemap.print()
			fmt.Println(imagemap.light())
		}
	}

	//	imagemap.print()
	fmt.Println(imagemap.light())

}
