package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

const workers = 5
const delay = 60

const (
	none = iota
	ready
	started
	complete
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

	queue := make(chan *node)
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	for i := 0; i < workers; i++ {
		go workElf(queue, &wg)
	}

	start := time.Now()
	for done := false; !done; done = addNewSteps(nodes, queue) {
	}
	wg.Wait()
	close(queue)
	fmt.Printf("\n%f seconds\n", time.Now().Sub(start).Seconds())
}

func addNewSteps(steps map[int]*node, queue chan<- *node) bool {
	var process []*node
	done := true
	for _, v := range steps {
		if v.status == none && len(v.prereq) == 0 {
			v.status = ready
			process = append(process, v)
		}
		if v.status != complete {
			done = false
		}
	}
	if process != nil {
		sort.Slice(process, func(i, j int) bool {
			return process[i].id < process[j].id
		})
		for _, v := range process {
			queue <- v
		}
	}
	return done
}

type node struct {
	id     int
	prereq map[int]*node
	depend map[int]*node
	status byte
}

func (n *node) String() string {
	switch n.status {
	case ready:
		return fmt.Sprintf("Node %c has \n no prereqs\n %d children %v\n and is ready to start\n", n.id, len(n.depend), n.depend)
	case started:
		return fmt.Sprintf("Node %c has \n no prereqs\n %d children %v\n and is started\n", n.id, len(n.depend), n.depend)
	case complete:
		return fmt.Sprintf("Node %c has \n no prereqs\n %d children %v\n and is complete\n", n.id, len(n.depend), n.depend)
	}
	return ""
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

func (n *node) finish() error {
	if n.status != started {
		return fmt.Errorf("Work not started, can't finish\n", n.id)
	}
	fmt.Printf("%c", n.id)
	n.status = complete
	if n.depend != nil {
		for _, v := range n.depend {
			delete(v.prereq, n.id)
		}
	}
	return nil
}

func workElf(step <-chan *node, done *sync.WaitGroup) {
	for s := range step {
		if s.status != ready {
			fmt.Printf("Ahh step %c not ready how did we get here?!\n", s.id)
		}
		s.status = started
		workTime := time.Duration(delay + s.id - 'A' + 1)
		time.Sleep(time.Second * workTime)
		s.finish()
		done.Done()
	}
}
