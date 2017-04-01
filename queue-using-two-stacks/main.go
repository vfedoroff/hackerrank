package main

import (
	"fmt"
	"log"
	"math"
)

type stack []interface{}

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Peek() interface{} {
	return s[len(s)-1]
}

func (s *stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *stack) Pop() interface{} {
	d := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return d
}

type Query struct {
	qtype int
	x     int
}

type Queue struct {
	backOfQueue  stack
	frontOfQueue stack
}

func (q *Queue) moveBackToFront() {
	for !q.backOfQueue.Empty() {
		v := q.backOfQueue.Pop()
		q.frontOfQueue.Push(v)
	}
}

func (q *Queue) Empty() bool {
	return q.frontOfQueue.Empty() && q.backOfQueue.Empty()
}

func (q *Queue) Push(v interface{}) {
	q.backOfQueue.Push(v)
}

func (q *Queue) Peek() interface{} {
	if q.Empty() {
		log.Fatalln("Queue underflow")
	}
	if q.frontOfQueue.Empty() {
		q.moveBackToFront()
	}
	return q.frontOfQueue.Peek()
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		log.Fatalln("Queue underflow")
	}
	if q.frontOfQueue.Empty() {
		q.moveBackToFront()
	}
	return q.frontOfQueue.Pop()
}

type Task struct {
	queries []Query
}

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

func (t *Task) Read() *Task {
	var err error
	var q int
	_, err = fmt.Scanf("%d", &q)
	t.checkErr(err)
	if !t.inRange(q, 1, int(math.Pow10(5))) {
		log.Fatalln("The value not in the expected range")
	}
	for i := 0; i < q; i++ {
		var qtype, x int
		fmt.Scanf("%d %d", &qtype, &x)
		//t.checkErr(err)
		if !t.inRange(qtype, 1, 3) {
			log.Fatalln("The query is not in the valid range")
		}

		if qtype == 1 && !t.inRange(x, 1, int(math.Pow10(9))) {
			log.Fatalln("The x is not in the valid range")
		}
		t.queries = append(t.queries, Query{qtype: qtype, x: x})
	}
	return t
}

func (t *Task) Solve() {
	var queue Queue
	for _, query := range t.queries {
		switch query.qtype {
		case 1:
			queue.Push(query.x)
		case 2:
			queue.Pop()
		case 3:
			v := queue.Peek()
			fmt.Printf("%d\n", v)
		}
	}
}

func main() {
	task := NewTask()
	task.Read().Solve()
}
