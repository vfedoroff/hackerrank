package main

import (
	"fmt"
	"log"
	"math"
)

type Task struct {
	dataSet [][]int
	swaps   int
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
	var a, n, d int
	_, err := fmt.Scan(&d)
	t.checkErr(err)
	if !t.inRange(d, 1, 15) {
		log.Fatalln("The value not in the expected range")
	}
	for i := 0; i < d; i++ {
		fmt.Scan(&n)
		if !t.inRange(n, 1, 2*int(math.Pow10(5))) {
			log.Fatalln("The value not in the expected range")
		}
		var arr []int
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			if !t.inRange(a, 1, int(math.Pow10(7))) {
				log.Fatalln("The value not in the expected range")
			}
			arr = append(arr, a)
		}
		t.dataSet = append(t.dataSet, arr)
	}
	return t
}

func (t *Task) mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := (len(arr)) / 2
	left := t.mergeSort(arr[:mid])
	right := t.mergeSort(arr[mid:])
	res := t.merge(left, right)
	return res

}

func (t *Task) merge(left, right []int) []int {
	lenLeft, lenRight := len(left), len(right)
	res := make([]int, 0, lenLeft+lenRight)
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			t.swaps += len(left)
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return right
}

func (t *Task) Solve() {

	for _, arr := range t.dataSet {
		t.mergeSort(arr)
		fmt.Printf("%d\n", t.swaps)
		t.swaps = 0
	}
}

func main() {
	NewTask().Read().Solve()
}
