package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	rec := initialize("37", 2)

	input := "580741"
	for i := 0; i < 1000000000; i++ {
		ok := rec.addResultWithCheck(input)
		if ok {
			fmt.Println(rec.cnt - 4)
			break
		}
		rec.moveElves()
	}
}

func initialize(input string, elves int) *recipe {
	var f, l *node
	for _, v := range []rune(input) {
		i := int(v - '0')
		if f == nil {
			f = &node{val: i}
			f.next = f
			f.prev = nil
			l = f
		} else {
			l = l.newNode(i)
		}
	}
	f.prev = l
	rec := &recipe{last: l}
	for i := 0; i < elves; i++ {
		rec.elves = append(rec.elves, f)
		f = f.next
	}
	return rec
}

type recipe struct {
	elves []*node
	last  *node
	cnt   int
}

func (r *recipe) lastN(n int) string {
	cur := r.last
	fn := make([]rune, n)
	for i := n - 1; i >= 0; i-- {
		fn[i] = rune('0' + cur.val)
		cur = cur.prev
	}
	return string(fn)
}

func (r *recipe) String() string {
	var sb strings.Builder
	//sb.WriteString(fmt.Sprintf("%d elves working\n",len(r.elves)))
	for cur := r.last.next; cur != r.last; cur = cur.next {
		sb.WriteString(strconv.Itoa(cur.val))
		sb.WriteRune(' ')
	}
	sb.WriteString(strconv.Itoa(r.last.val))
	for i := 0; i < len(r.elves); i++ {
		sb.WriteString(fmt.Sprintf("\nElf %d on %d", i+1, r.elves[i].val))
	}
	return sb.String()
}

func (r *recipe) addResultWithCheck(input string) bool {
	sum := 0
	for _, v := range r.elves {
		sum += v.val
	}
	rn := []rune(strconv.Itoa(sum))
	for i := 0; i < len(rn); i++ {
		r.last = r.last.newNode(int(rn[i] - '0'))
		r.cnt++
		out := r.lastN(len(input))
		if out == input {
			return true
		}
	}
	return false
}

func (r *recipe) addResult() {
	sum := 0
	for _, v := range r.elves {
		sum += v.val
	}
	rn := []rune(strconv.Itoa(sum))
	for i := 0; i < len(rn); i++ {
		r.last = r.last.newNode(int(rn[i] - '0'))
		r.cnt++
	}
}

func (r *recipe) moveElves() {
	for k, v := range r.elves {
		moves := v.val + 1
		for i := 0; i < moves; i++ {
			v = v.next
		}
		r.elves[k] = v
	}
}

func (n *node) newNode(val int) *node {
	newNode := &node{val, n, n.next}
	n.next = newNode
	return newNode
}

func (n *node) String() string {
	return fmt.Sprintf("node: %d prev %d next %d", n.val, n.prev.val, n.next.val)
}

type node struct {
	val  int
	prev *node
	next *node
}
