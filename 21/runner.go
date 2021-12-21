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

func (p *Player) move(move int) {
	for i := 0; i < move; i++ {
		p.position++
		if p.position > 10 {
			p.position = 1
		}
	}
	p.score += p.position
}

type GameState struct {
	players [2]Player
	count   int
}

// rolling a d3 3 times, this is  the frequency map out outcomes
// from this set:
// [3, 4, 5, 4, 5, 6, 5, 6, 7, 4, 5, 6, 5, 6, 7, 6, 7, 8, 5, 6, 7, 6, 7, 8, 7, 8, 9]
var DiceOutcomes = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

// this enumertes the rolls themselves, for ordering
var possible_rolls = [7]int{3, 4, 5, 6, 7, 8, 9}

func play_a_turn(g GameState, player int, universes uint64, wins *[2]uint64) {

	for _, r := range possible_rolls {
		/*if old_score != g.players[player].score {
			fmt.Errorf("Score not equal!!!")
			os.Exit(1)
		}*/
		newGameState := g
		newGameState.players[player].move(r)
		fanout := uint64(DiceOutcomes[r]) * universes

		if newGameState.players[player].score >= 21 {
			wins[player] += fanout
		} else {
			play_a_turn(newGameState, (player+1)%2, fanout, wins)
		}
	}
}

func quantum_gameplay() {
	g := GameState{}
	/*
		g.players[0].position = 4
		g.players[1].position = 8
	*/
	g.players[0].position = 9
	g.players[1].position = 3
	var wins [2]uint64
	play_a_turn(g, 0, 1, &wins)
	fmt.Println("W1: ", uint64(wins[0]), " W2: ", uint64(wins[1]))
	fmt.Println(wins)
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

	quantum_gameplay()
}
