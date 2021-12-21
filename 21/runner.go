package main

import (
	"fmt"
)

type Die struct {
	current, count, max int
}

func (d *Die) roll() int {
	ret := d.current
	d.current++
	if d.current > d.max {
		d.current = 1
	}
	d.count++
	return ret
}

type Player struct {
	position, score int
}

func (p *Player) play(d *Die) bool {
	move := d.roll() + d.roll() + d.roll()
	for i := 0; i < move; i++ {
		p.position++
		if p.position > 10 {
			p.position = 1
		}
	}

	p.score += p.position

	return p.score >= 1000
}

func main() {

	d := Die{1, 0, 100}
	/*
		p1 := Player{4, 0}
		p2 := Player{8, 0}
	*/
	p1 := Player{9, 0}
	p2 := Player{3, 0}
	var loser Player

	for i := 0; i < 10000; i++ {
		if p1.play(&d) {
			loser = p2
			break
		}
		if p2.play(&d) {
			loser = p1
			break
		}
	}
	fmt.Println(loser.score, d.count, loser.score*d.count)
}
