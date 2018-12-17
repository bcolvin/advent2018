package main

import (
    "log"
    "unicode"
    "fmt"
    "io/ioutil"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    orig := []rune(string(b))
    bestRem := 'A'
    bestLen := len(orig)
    for _, v := range "abcdefghijklmnopqrstuvwxyz" {
        str := removeRune(v,orig)
        for ok := false; !ok; {
            str, ok = react(str)
        }
        resLen := len(str)
        if resLen < bestLen {
            bestLen = resLen
            bestRem = v
        }
    }
    fmt.Printf("Best remove was %c with new length %d\n",bestRem,bestLen)
}

func removeRune(rem rune, oldStr []rune) []rune {
    var newStr []rune
    for i:= 0; i < len(oldStr); i++  {
        if unicode.ToLower(oldStr[i]) != rem {
            newStr = append(newStr,oldStr[i])
        }
    }
    return newStr
}

func react(oldStr []rune) ([]rune, bool) {
    var newStr []rune
    ok := true
    for i := 0; i < len(oldStr); i++ {
        if i+1 < len(oldStr) && oldStr[i] != oldStr[i+1] && unicode.ToLower(oldStr[i]) == unicode.ToLower(oldStr[i+1]) {
            i++
            ok = false
        } else {
            newStr = append(newStr,oldStr[i])
        }
    }
    return newStr, ok
}