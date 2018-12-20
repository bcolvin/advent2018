package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := []rune(string(b))
	for ok := false; !ok; {
		str, ok = react(str)
	}
	fmt.Printf("%d\n", len(str))
}

func react(oldStr []rune) ([]rune, bool) {
	var newStr []rune
	ok := true
	for i := 0; i < len(oldStr); i++ {
		if i+1 < len(oldStr) && oldStr[i] != oldStr[i+1] && unicode.ToLower(oldStr[i]) == unicode.ToLower(oldStr[i+1]) {
			i++
			ok = false
		} else {
			newStr = append(newStr, oldStr[i])
		}
	}
	return newStr, ok
}
