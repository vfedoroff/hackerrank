package main

import (
	"fmt"
	"log"
	"math"
)

type Task struct {
	arr []int
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
	var a, n int
	_, err := fmt.Scan(&n)
	t.checkErr(err)
	if !t.inRange(n, 2, 600) {
		log.Fatalln("The value not in the expected range")
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if !t.inRange(a, 1, 2*int(math.Pow10(6))) {
			log.Fatalln("The value not in the expected range")
		}
		t.arr = append(t.arr, a)
	}
	return t
}

func (t *Task) swap(i, j int) {
	t.arr[i], t.arr[j] = t.arr[j], t.arr[i]
}

func (t *Task) bubbleSort() int {
	var swapNum int
	arrLen := len(t.arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < arrLen-1; i++ {
			if t.arr[i+1] < t.arr[i] {
				t.swap(i, i+1)
				swapNum++
				swapped = true
			}
		}
	}
	return swapNum
}

func (t *Task) Solve() {
	res := t.bubbleSort()
	fmt.Printf("Array is sorted in %d swaps.\n", res)
	fmt.Printf("First Element: %d\n", t.arr[0])
	fmt.Printf("Last Element: %d\n", t.arr[len(t.arr)-1])

}

func main() {
	NewTask().Read().Solve()
}
