package main

import (
	"bufio"
	"log"
	"os"
    "fmt"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	twos,threes := 0,0
	s := bufio.NewScanner(f)
	for s.Scan() {
        re := s.Text()
        var cha [26]int
        for i := 0; i < len(re); i++ {
            cha[re[i] - 'a']++
        }

        isTwo,isThree := false,false
        for _,i := range cha {
            if i == 2 && !isTwo {
                twos++
                isTwo = true
            }
            if i == 3 && !isThree {
                threes++
                isThree = true
            }
        }
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d * %d = %d",twos, threes, twos*threes);
}