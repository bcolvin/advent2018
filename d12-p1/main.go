package main

import (
	"bufio"
	"fmt"
	"log"
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
	var combos []*genetics
	var initial PlantRow
	for i := 0; s.Scan(); i++ {
		var err error
		var input string
		if i == 0 {
			_, err = fmt.Sscanf(s.Text(), "initial state: %s", &input)
			initial = PlantRow{[]rune(input), 0}
		} else if i != 1 {
			var output rune
			_, err = fmt.Sscanf(s.Text(), "%s => %c", &input, &output)
			combos = append(combos, &genetics{[]rune(input), output})
		}

		if err != nil {
			log.Fatalf("Error reading into value %v", err)
		}
	}

	g := 20
	initial.zero = initial.row[0]
	initial.row[0] = 'z'
	for i := 0; i < g; i++ {
		fmt.Println(initial)
		initial.reproduce(combos)
		initial.trim()
	}
	fmt.Println("Sum:", initial.sum())
}

type PlantRow struct {
	row  []rune
	zero rune
}

func (pr PlantRow) String() string {
	var sbB strings.Builder
	for i := 0; i < len(pr.row); i++ {
		fmt.Fprintf(&sbB, "%c", pr.row[i])
	}
	return sbB.String()
}

func (pr *PlantRow) reproduce(combos []*genetics) {
	curLen := len(pr.row)
	comp := []rune(".....")
	var newZero rune
	for i := 0; i < curLen+4; i++ {
		r := '.'
		if i < curLen {
			r = pr.row[i]
		}
		comp = append(comp[1:], r)
		r = '.'
		zero := false
		for _, v := range combos {
			a, b := pr.getOffspring(comp, v.combo)
			if b {
				zero = b
			}
			if a {
				r = v.output
				break
			}
		}
		if zero {
			newZero = r
			r = 'z'
		}
		if i < curLen {
			pr.row[i] = r
		} else {
			pr.row = append(pr.row, r)
		}
	}
	pr.zero = newZero
}

func (pr *PlantRow) getOffspring(iter []rune, formula []rune) (bool, bool) {
	var zero bool
	for i := 0; i < len(iter); i++ {
		r := iter[i]
		if r == 'z' {
			r = pr.zero
			if i == len(iter)/2 {
				zero = true
			}
		}
		if r != formula[i] {
			return false, zero
		}
	}
	return true, zero
}

func (pr *PlantRow) trim() {
	var firstPlant, lastPlant int
	for i := 0; i < len(pr.row); i++ {
		if pr.row[i] == '#' || pr.row[i] == 'z' {
			firstPlant = i
			break
		}
	}
	for i := len(pr.row) - 1; i >= 0; i-- {
		if pr.row[i] == '#' || pr.row[i] == 'z' {
			lastPlant = i
			break
		}
	}
	pr.row = pr.row[firstPlant : lastPlant+1]
}

func (pr PlantRow) sum() int {
	ret := 0
	zero := 0
	for i := 0; i < len(pr.row); i++ {
		if pr.row[i] == 'z' {
			zero = i
		}
	}
	for i := 0; i < len(pr.row); i++ {
		if pr.row[i] == '#' {
			ret += i - zero
		}
	}
	return ret
}

type genetics struct {
	combo  []rune
	output rune
}
