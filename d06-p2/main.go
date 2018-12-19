package main

import (
    "log"
    "os"
    "bufio"
    "fmt"
    "math"
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
    maxX, maxY := 0,0
    for s.Scan() {
        p := xy{0,0}
        fmt.Sscanf(s.Text(),"%d, %d", &p.x,&p.y)
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
    maxDist := 10000
    total := 0
    for i := 0; i <= maxY; i++ {
        for j := 0; j <= maxX; j++ {
            curPos := xy{j, i}
            dist := 0
            for k := range xys {
                dist += distance(curPos, k)
                if dist > maxDist {
                    break
                }
            }
            v := xys[curPos]
            if v != 0 {
                grid[curPos] = v
            } else if dist < maxDist {
                grid[curPos] = '#'
            } else {
                grid[curPos] = '.'
            }
            //fmt.Printf("%c",grid[curPos])
            if dist < maxDist {
                total++
            }
        }
        //fmt.Println()
    }
    fmt.Printf("Total area: %d\n",total)
}

type xy struct {
    x,y int
}

func distance(p,q xy) int{
    return int(math.Abs(float64(p.x - q.x))+math.Abs(float64(p.y - q.y)))
}