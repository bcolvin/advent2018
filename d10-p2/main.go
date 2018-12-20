package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var night sky
	var minX, minY, maxX, maxY int
	for s.Scan() {
		light := &light{}
		_, err := fmt.Sscanf(s.Text(), "position=<%d, %d> velocity=<%d, %d>", &light.x, &light.y, &light.vx, &light.vy)
		if err != nil {
			log.Fatalf("Error reading into value %v", err)
		}
		night.lights = append(night.lights, light)
		if light.x < minX {
			minX = light.x
		}
		if light.x > maxX {
			maxX = light.x
		}
		if light.y < minY {
			minY = light.y
		}
		if light.y > maxY {
			maxY = light.y
		}
	}
	night.x = maxX - minX
	night.y = maxY - minY

	night.xShift = int(math.Abs(float64(minX)))
	night.yShift = int(math.Abs(float64(minY)))

	for i := 1; i < 10; i++ {
		night.second = i
		night.grid()
	}
}

type light struct {
	x, y   int
	vx, vy int
}

func (l *light) currentPoint(second int) (int, int) {
	return l.x + (second * l.vx), l.y + (second * l.vy)
}

type sky struct {
	second         int
	xShift, yShift int
	x, y           int
	lights         []*light
}

func (s *sky) grid() {
	night := make(map[int][]rune, s.y)
	for _, v := range s.lights {
		x, y := v.currentPoint(s.second)
		str := night[y+s.yShift]
		if str == nil {
			str = []rune(strings.Repeat(" ", s.x))
			night[y+s.yShift] = str
		}
		str[x+s.xShift] = '#'
	}

	fmt.Println(s.second)
	for i := 0; i < s.y; i++ {
		r := string(night[i])
		if r == "" {
			fmt.Println()
		} else {
			fmt.Println(r)
		}
	}
}
