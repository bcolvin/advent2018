package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	r := 48
	xys := make(map[xy]int)
	maxX, maxY := 0, 0
	for s.Scan() {
		p := xy{0, 0}
		fmt.Sscanf(s.Text(), "%d, %d", &p.x, &p.y)
		xys[p] = r
		r++
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	grid := make(map[xy]int)
	//-1 +1 for infinite check
	for i := -1; i <= maxY+1; i++ {
		for j := -1; j <= maxX+1; j++ {
			curPos := xy{j, i}
			val := xys[curPos]
			if val != 0 {
				grid[curPos] = val
				//fmt.Printf("%c", val)
				continue
			}
			closest := 0
			for k, v := range xys {
				dist := distance(curPos, k)
				if closest == 0 || dist < closest {
					closest = dist
					grid[curPos] = v
				} else if dist == closest {
					grid[curPos] = '.'
				}
			}
			//fmt.Printf("%c", grid[curPos])
		}
		//fmt.Println()
	}

	finA := make(map[int]int)
	finI := make(map[int]int)
	for k, v := range grid {
		if k.x <= maxX && k.x >= 0 && k.y <= maxY && k.y >= 0 {
			finA[v]++
		}
		finI[v]++
	}
	max := 0
	val := 0
	for k, v := range finA {
		if v != finI[k] {
			//fmt.Printf("%c is infinite continue\n",k)
			continue
		}
		if v > max {
			max = v
			val = k
		}
	}
	fmt.Printf("MAX %c : %d\n", val, max)
}

type xy struct {
	x, y int
}

func distance(p, q xy) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}
