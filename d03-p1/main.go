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

	s := bufio.NewScanner(f)
	var xys []*xy
	for s.Scan() {
	    var id string
		var x,y,w,h int
        _, err := fmt.Sscanf(s.Text(), "#%s @ %d,%d: %dx%d", &id,&x,&y,&w,&h)
        if err != nil {
            log.Fatalf("Error reading into value")
        }
        xys = append(xys,&xy{id,x,y,w,h})
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

    var fab [1000][1000]string
	for _,xy := range xys {
        for i := 0; i < xy.w; i++ {
            for j := 0; j < xy.h; j++ {
                if fab[j+xy.y][i+xy.x] == "" {
                    fab[j+xy.y][i+xy.x] = xy.id
                } else {
                    fab[j+xy.y][i+xy.x] = "X"
                }
            }
        }
    }
    cnt := 0
    for i:= 0; i < 1000; i++ {
        for j := 0; j < 1000; j++ {
            if fab[i][j] == "X" {
                cnt++
            }
        }
    }
    fmt.Println(cnt)
}

type xy struct {
    id string
    x,y int
    w,h int
}