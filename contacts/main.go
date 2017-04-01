package main

import (
	"fmt"
	"log"
	"math"
)

type Tries struct {
	count int
	tree  [26]*Tries
}

func (t *Tries) Add(str string) {
	if str == "" || t == nil {
		return
	}
	pos := int(str[0] - byte(rune('a')))
	if t.tree[pos] == nil {
		t.tree[pos] = &Tries{count: 1}
	} else {
		t.tree[pos].count++
	}
	t.tree[pos].Add(str[1:])
}

func (t *Tries) Search(str string) int {
	if str == "" || t == nil {
		return 0
	}
	pos := int(str[0] - byte(rune('a')))
	if len(str) == 1 {
		if st := t.tree[pos]; st != nil {
			return st.count
		}
		return 0
	}
	return t.tree[pos].Search(str[1:])
}

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func (t *Task) inRange(v, left, right int) bool {
	if left > right {
		left, right = right, left
	}
	return v >= left && v <= right
}

func (t *Task) Solve() {
	var n int
	var action, line string
	_, err := fmt.Scan(&n)
	t.checkErr(err)
	if !t.inRange(n, 1, int(math.Pow10(5))) {
		log.Fatalln("The value not in the expected range")
	}
	tries := new(Tries)

	for i := 0; i < n; i++ {
		fmt.Scan(&action, &line)
		if !t.inRange(len(line), 1, 21) {
			log.Fatalln("The value not in the expected range")
		}
		if action == "add" {
			tries.Add(line)
		} else {
			fmt.Println(tries.Search(line))
		}
	}
}

func main() {
	NewTask().Solve()
}
