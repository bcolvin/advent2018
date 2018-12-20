package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var players, turns int
		_, err := fmt.Sscanf(s.Text(), "%d players; last marble is worth %d points", &players, &turns)
		if err != nil {
			log.Fatalf("Error reading into value")
		}
		//part 2 says last marble is 100x larger, weird wording, but the right result is interpretted as it should last 100x more than original turns
		g := initGame(players, turns*100)
		g.play()
		fmt.Println(g)
	}
}

func initGame(players, turns int) *game {
	g := &game{}
	for i := 1; i <= players; i++ {
		g.players = append(g.players, &player{id: i})
	}
	g.turns = turns
	g.currentMarble = &marble{val: 0}
	g.currentMarble.next = g.currentMarble
	g.currentMarble.prev = g.currentMarble
	return g
}

type game struct {
	players       []*player
	turns         int
	currentMarble *marble
}

func (g *game) String() string {
	if g.turns == 0 {
		winner := g.winner()
		return fmt.Sprintf("%d players; %d turns; Winner %s", len(g.players), g.turns, winner)
	}
	return fmt.Sprintf("%d players; %d turns; Not concluded", len(g.players), g.turns)
}

func (g *game) play() {
	curM := 1
	for {
		for _, p := range g.players {
			g.turns--
			m := &marble{val: curM}
			val := g.placeMarble(m)
			if val > 0 {
				p.points += val
			}
			curM++
			if g.turns == 0 {
				return
			}
		}
	}
}

func (g *game) placeMarble(m *marble) int {
	points := 0
	if m.val%23 == 0 {
		points += m.val
		for i := 0; i < 7; i++ {
			g.currentMarble = g.currentMarble.prev
		}
		points += g.currentMarble.val
		g.currentMarble = g.currentMarble.next

		g.currentMarble.prev.remove()
	} else {
		m.insert(g.currentMarble.next, g.currentMarble.next.next)
		g.currentMarble = m
	}
	return points
}

func (g *game) board() string {
	s := ""
	for cur := g.currentMarble.next; cur != g.currentMarble; cur = cur.next {
		s = fmt.Sprintf("%s %d", s, cur.val)
	}
	return s
}

func (g *game) winner() *player {
	var winner *player
	for _, p := range g.players {
		if winner == nil || winner.points < p.points {
			winner = p
		}
	}
	return winner
}

type marble struct {
	val  int
	next *marble
	prev *marble
}

func (m *marble) insert(m1, m2 *marble) {
	m.prev = m1
	m.next = m2
	m1.next = m
	m2.prev = m
}
func (m *marble) remove() {
	m1 := m.prev
	m2 := m.next
	m1.next = m2
	m2.prev = m1
}

func (m *marble) String() string {
	return fmt.Sprintf("Marble %d is between marble %d and %d", m.val, m.prev.val, m.next.val)
}

type player struct {
	id     int
	points int
}

func (p *player) String() string {
	return fmt.Sprintf("Player %d has %d points", p.id, p.points)
}
