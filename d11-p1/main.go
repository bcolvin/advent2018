package main

import (
	"fmt"
	"strings"
)

func main() {
	width := 300
	height := 300
	serialNumber := 5177
	grid := createGrid(height, width, serialNumber, power)
	fmt.Println(grid.findCell())
}

func power(x, y, serialNumber int) int {
	id := x + 10
	num := id * y
	num += serialNumber
	num *= id
	num = (num / 100) % 10
	return num - 5
}

type Grid struct {
	height, width int
	graph         [][]int
}

func createGrid(height, width, serialNumber int, pFunc func(x int, y int, serialNumber int) int) *Grid {
	grid := &Grid{height: height, width: width}
	for i := height; i > 0; i-- {
		var row []int
		for j := 0; j < width; j++ {
			row = append(row, pFunc((j+1), i, serialNumber))
		}
		grid.graph = append(grid.graph, row)
	}
	return grid
}

func (grid Grid) findCell() (int, int, int) {
	max := 0
	x, y := 0, 0
	for i := grid.height - 1; i > 2; i-- {
		for j := 0; j < grid.width-3; j++ {
			yE := i - 3
			xE := j + 3
			sum := grid.sum(j, xE, i, yE)
			//fmt.Println("xy: ",j,xE,yE,i," = ",sum)
			if sum > max {
				max = sum
				x = j + 1
				y = grid.width - yE - 2
			}
		}
	}
	return x, y, max
}

func (grid Grid) sum(x, xE, y, yE int) int {
	sum := 0
	for i := yE; i < y; i++ {
		for j := x; j < xE; j++ {
			sum += grid.graph[i][j]
		}
	}
	return sum
}

func (grid Grid) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "  0")
	for i := 1; i <= grid.width; i++ {
		if i < 10 {
			fmt.Fprintf(&sb, " ")
		}
		if i < 100 {
			fmt.Fprintf(&sb, " ")
		}
		fmt.Fprintf(&sb, " %d", i)
	}

	for i := grid.height - 1; i >= 0; i-- {
		fmt.Fprintln(&sb)

		rNum := grid.height - i
		if rNum < 10 {
			fmt.Fprintf(&sb, " ")
		}
		if rNum < 100 {
			fmt.Fprintf(&sb, " ")
		}
		fmt.Fprintf(&sb, "%d ", rNum)
		for j := 0; j < grid.width; j++ {
			val := grid.graph[i][j]
			if val >= 0 {
				fmt.Fprintf(&sb, "  %v ", val)
			} else {
				fmt.Fprintf(&sb, " %v ", val)
			}
		}
	}
	return sb.String()
}
