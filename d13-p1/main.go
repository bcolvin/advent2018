package main

import (
	"bufio"
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
	for i := 0; s.Scan(); i++ {

		if err != nil {
			log.Fatalf("Error reading into value %v", err)
		}
	}
}

type track struct {
	x, y int
}
