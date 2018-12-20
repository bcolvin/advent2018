package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	nodes := make(map[int]*node)
	for s.Scan() {
		var pr, st int
		_, err := fmt.Sscanf(s.Text(), "Step %c must be finished before step %c can begin.", &pr, &st)
		if err != nil {
			log.Fatalf("Error reading into value")
		}

		step := nodes[st]
		if step == nil {
			step = &node{id: st}
			nodes[step.id] = step
		}

		pre := nodes[pr]
		if pre == nil {
			pre = &node{id: pr}
			nodes[pre.id] = pre
		}
		step.addPrereq(pre)
		pre.addDepend(step)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	process := addSteps(nodes)
	for len(process) > 0 {
		process[0].process()
		process = addSteps(nodes)
	}
	fmt.Println()
}

func addSteps(steps map[int]*node) []*node {
	var process []*node
	for _, v := range steps {
		if v.complete != true && len(v.prereq) == 0 {
			process = append(process, v)
		}
		//fmt.Printf(v)
	}
	if process != nil {
		sort.Slice(process, func(i, j int) bool {
			return process[i].id < process[j].id
		})
	}
	return process
}

type node struct {
	id       int
	prereq   map[int]*node
	depend   map[int]*node
	complete bool
}

func (n *node) String() string {
	if n.complete {
		return fmt.Sprintf("Node %c has \n no prereqs\n %d children %v\n and is complete\n", n.id, len(n.depend), n.depend)
	}
	return fmt.Sprintf("Node %c has \n %d prereqs: %v\n %d children %v\n and is not complete\n", n.id, len(n.prereq), n.prereq, len(n.depend), n.depend)
}

func (n *node) addPrereq(pre *node) {
	if n.prereq == nil {
		n.prereq = make(map[int]*node, 10)
	}
	n.prereq[pre.id] = pre
}

func (n *node) addDepend(depend *node) {
	if n.depend == nil {
		n.depend = make(map[int]*node, 10)
	}
	n.depend[depend.id] = depend
}

func (n *node) process() error {
	if len(n.prereq) > 0 {
		return fmt.Errorf("Prerequisites not completed yet, %d left\n", len(n.prereq))
	}
	n.complete = true
	fmt.Printf("%c", n.id)
	if n.depend != nil {
		for _, v := range n.depend {
			delete(v.prereq, n.id)
		}
	}
	return nil
}
