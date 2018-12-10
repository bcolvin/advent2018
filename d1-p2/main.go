package main

import (
 "fmt"
    "os"
    "log"
    "bufio"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    var nums []int
    s := bufio.NewScanner(f)
    for s.Scan() {
        var n int
        _, err := fmt.Sscanf(s.Text(), "%d", &n)
        if err != nil {
            log.Fatalf("Error reading into value")
        }
        nums = append(nums, n)
    }
    if err := s.Err(); err != nil {
        log.Fatal(err)
    }

    total := 0
    seen := map[int]bool{0:true}
    for {
        for _,n := range nums {
            total += n
            if seen[total] {
                fmt.Println(total)
                return
            }
            seen[total] = true
        }
    }
}
