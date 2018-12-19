package main

import (
	"log"
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func main() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    vals := strings.Split(string(b), " ")
    id := int('A')
    offset := 0
    root := readNode(id,vals,&offset)
    fmt.Println(root.sumMeta())
}

func readNode(id int, vals []string, offset *int) *node {
    nVals := vals[*offset:]
    children, err := strconv.Atoi(nVals[0])
    if err != nil {
        panic(fmt.Errorf("Error reading child count for %c", id, err))
    }
    metas, err := strconv.Atoi(nVals[1])
    if err != nil {
        panic(fmt.Errorf("Error reading meta count for %c", id, err))
    }
    n := &node{id: id}
    *offset += 2
    for i:= 1; i <= children; i++ {
        nn := readNode(n.id+i,vals,offset)
        n.children = append(n.children,nn)
    }
    n.meta = vals[*offset:(*offset+metas)]
    *offset += metas
    return n
}

type node struct {
    id int
    children []*node
    meta []string
}

func (n *node) String() string {
    r := fmt.Sprintf("Node %c: %d children and %d meta", n.id, len(n.children), len(n.meta))
    r += fmt.Sprintf("\n Meta: %v",n.meta)
    for _,v := range n.children {
        r += fmt.Sprintf("\n Child %c:\n  %s",v.id,v)
    }
    return r
}

func (n *node) sumMeta() int {
    meta := 0
    var cMeta []int
    if len(n.children) > 0 {
        cMeta = make([]int, len(n.children))
        for i, v := range n.children {
            cMeta[i] = v.sumMeta()
        }
    }
    if len(n.meta) > 0 {
        for _,v := range n.meta {
            i, err := strconv.Atoi(v)
            if err != nil {
                fmt.Errorf("Error converting meta to int %s, %v\n", v,err)
            }

            if len(cMeta) > 0 {
                i--
                if len(cMeta) > i {
                    meta += cMeta[i]
                }
            } else {
                meta += i
            }
        }
    }
    return meta
}
