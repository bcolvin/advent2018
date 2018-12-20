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

	total := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("Error reading into value")
		}
		total += n
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
