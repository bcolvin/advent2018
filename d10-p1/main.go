package main

import (
	"log"
    "fmt"
    "os"
    "bufio"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    s := bufio.NewScanner(f)
    for s.Scan() {
        var px,py int
        var vx,vy int
        _, err := fmt.Sscanf(s.Text(), "position=<%d, %d> velocity=<%d, %d>", &px,&py,&vx,&vy)
        if err != nil {
            log.Fatalf("Error reading into value %v",err)
        }
    }
}