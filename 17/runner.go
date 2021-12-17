package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Probe struct {
	x  int
	y  int
	xv int
	yv int
}

type Target struct {
	minx int
	maxx int
	miny int
	maxy int
}

func (p *Probe) hit(target Target) bool {
	return p.x >= target.minx && p.x <= target.maxx && p.y >= target.miny && p.y <= target.maxy
}

func (p *Probe) ontarget(target Target) bool {
	return p.x <= target.maxx && p.y >= target.miny
}

func (p *Probe) step() {
	p.x += p.xv
	p.y += p.yv
	// move closer to 0
	if p.xv > 0 {
		p.xv--
	} else if p.xv < 0 {
		p.xv++
	}
	p.yv--
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)
	// TODO: parse this instead of just copypasta
	fmt.Println("target: ", input)

	minx := 25
	maxx := 67
	miny := -260
	maxy := -200
	/*
		minx := 20
		maxx := 30
		miny := -10
		maxy := -5
	*/

	target := Target{minx, maxx, miny, maxy}

	overall_max := 0
	hits := 0
	for yv := target.miny; yv <= (0 - target.miny); yv++ {
		for xv := 1; xv <= target.maxx; xv++ {
			probe := Probe{0, 0, xv, yv}
			maxy := 0
			for probe.ontarget(target) {
				probe.step()
				if probe.y > maxy {
					maxy = probe.y
				}
				if probe.hit(target) {
					if maxy > overall_max {
						overall_max = maxy
					}
					hits++
					break
				}

			}

		}

	}

	fmt.Println(overall_max)
	fmt.Println(hits)

	/*
			The probe's x position increases by its x velocity.
		The probe's y position increases by its y velocity.
		Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
		Due to gravity, the probe's y velocity decreases by 1.
	*/

}
