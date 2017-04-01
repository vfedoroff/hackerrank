package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
)

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() interface{}   { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An MaxHeap is a max-heap of ints.
type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool {
	return h.MinHeap[i] > h.MinHeap[j]
}

type Queue struct {
	lower  *MaxHeap
	higher *MinHeap
}

func NewQueue() *Queue {
	lower := &MaxHeap{}
	higher := &MinHeap{}
	heap.Init(lower)
	heap.Init(higher)
	q := Queue{lower: lower, higher: higher}
	return &q
}

func (q *Queue) Add(x int) {
	if q.lower.Len() == 0 {
		heap.Push(q.lower, x)
	} else {
		if q.lower.Len() > q.higher.Len() {
			// lower is bigger
			if q.lower.Top().(int) > x {
				// Balance - Take highest from lower, put it in higher
				heap.Push(q.higher, q.lower.Top())
				heap.Pop(q.lower)
				heap.Push(q.lower, x)
			} else {
				heap.Push(q.higher, x)
			}
		} else {
			// higher is bigger
			if q.higher.Top().(int) >= x {
				heap.Push(q.lower, x)
			} else {
				heap.Push(q.lower, q.higher.Top())
				heap.Pop(q.higher)
				heap.Push(q.higher, x)
			}
		}
	}
}

func (q *Queue) Median() float64 {
	n := q.lower.Len() + q.higher.Len()
	if n%2 == 0 {
		return float64(q.lower.Top().(int)+q.higher.Top().(int)) / 2
	} else {
		return float64(q.lower.Top().(int))
	}

}

type Task struct {
	a []int
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
	var n int
	_, err = fmt.Scanf("%d", &n)
	t.checkErr(err)
	if !t.inRange(n, 1, int(math.Pow10(5))) {
		log.Fatalln("The value not in the expected range")
	}
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)
		if !t.inRange(x, 0, int(math.Pow10(5))) {
			log.Fatalln("The 'a' is not in the valid range")
		}
		t.a = append(t.a, x)
	}
	return t
}

func (t *Task) Solve() {
	q := NewQueue()
	for _, a := range t.a {
		q.Add(a)
		fmt.Printf("%.1f\n", q.Median())
	}
}

func main() {
	task := NewTask()
	task.Read().Solve()
}
