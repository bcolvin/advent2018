package main

import (
"bufio"
"log"
"os"
"time"
"strings"
"sort"
"fmt"
"strconv"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    s := bufio.NewScanner(f)
    var shifts []*entry
    for s.Scan() {
        sp := strings.SplitAfter(s.Text(),"]")
        tm ,err := time.Parse("[2006-01-02 15:04] ",sp[0])
        if err != nil {
            log.Fatalf("Error parsing time: %v",err)
        }
        sp = strings.Split(sp[1]," ")
        var e *entry
        switch sp[1] {
        case "Guard":
            g, err := strconv.Atoi(sp[2][1:])
            if err != nil {
                log.Fatalf("Error parsing guard: %v",err)
            }
            e = &entry{tm, g, "start"}
        case "falls":
            e = &entry{tm,0,"sleep"}
        case "wakes":
            e = &entry{tm,0,"wake"}
        }
        if e != nil {
            shifts = append(shifts,e)
        }
    }
    if err := s.Err(); err != nil {
        log.Fatal(err)
    }
    sort.Slice(shifts, func(i,j int) bool {
        return shifts[i].time.Before(shifts[j].time)
    })
    var guard int
    for _,s := range shifts {
        if s.guard != 0 && guard != s.guard {
            guard = s.guard
        }
        s.guard = guard
    }

    sleep := make(map[int][]int)
    var start time.Time
    var min,v int
    for _,s := range shifts {
        if s.action == "sleep" {
            start = s.time
        } else if s.action == "wake" && start.Before(s.time) {
            _,ok := sleep[s.guard]
            if !ok {
                sleep[s.guard] = make([]int,60)
            }
            for i := start.Minute(); i < s.time.Minute(); i++ {
                sleep[s.guard][i]++
                cur := sleep[s.guard][i]
                if v < cur {
                    v = cur
                    min = i
                    guard = s.guard
                }
            }
        }
    }
    fmt.Printf("Guard %d most frequently slept on min %d\n",guard,min)
    fmt.Println(guard*min)
}

type entry struct {
    time time.Time
    guard int
    action string
}