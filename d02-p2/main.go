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

	var words []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(words); i++ {
		s1 := words[i]
		for j := i + 1; j < len(words); j++ {
			s2 := words[j]
			comm, b := compare(s1, s2)
			if b {
				fmt.Printf("%s\n", comm)
			} else {
				//fmt.Printf("NO: %s %s\n", s1,s2)
			}
		}
	}
}
func compare(a, b string) (string, bool) {
	diff := 0
	var str []rune
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			} else {
				str = append(str, rune(a[i]))
			}
		}
	}

	if diff == 1 {
		return string(str), true
	} else {
		return "", false
	}
}
