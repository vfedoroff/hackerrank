package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

type iceCream struct {
	ID   int
	Cost int
}

type CostSorted []iceCream

func (a CostSorted) Len() int           { return len(a) }
func (a CostSorted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CostSorted) Less(i, j int) bool { return a[i].Cost < a[j].Cost }

type trip struct {
	Money   int
	Flavors []iceCream
}

type Task struct {
	trips []trip
}

func NewTask() *Task {
	return &Task{}
}

func (task *Task) checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func (task *Task) inRange(v, left, right int) bool {
	if left > right {
		left, right = right, left
	}
	return v >= left && v <= right
}

func (task *Task) Read() *Task {
	var err error
	var t int
	_, err = fmt.Scanf("%d", &t)
	task.checkErr(err)
	if !task.inRange(t, 1, 50) {
		log.Fatalln("The value not in the expected range")
	}
	task.trips = make([]trip, t)
	for i := 0; i < t; i++ {
		var m int
		_, err = fmt.Scanf("%d", &m)
		task.checkErr(err)
		if !task.inRange(m, 2, int(math.Pow10(4))) {
			log.Fatalln("The value not in the expected range")
		}
		tr := trip{
			Money: m,
		}
		var n int
		_, err = fmt.Scanf("%d", &n)
		task.checkErr(err)
		if !task.inRange(n, 2, int(math.Pow10(4))) {
			log.Fatalln("The value not in the expected range")
		}
		for j := 1; j <= n; j++ {
			var cost int
			_, err = fmt.Scanf("%d", &cost)
			task.checkErr(err)
			if !task.inRange(cost, 1, int(math.Pow10(4))) {
				log.Fatalln("The value not in the expected range")
			}
			ic := iceCream{
				ID:   j,
				Cost: cost,
			}
			tr.Flavors = append(tr.Flavors, ic)
		}
		sort.Sort(CostSorted(tr.Flavors))
		task.trips[i] = tr
	}
	return task
}

func (task *Task) Solve() {
	for _, t := range task.trips {
		task.search(t.Flavors, t.Money)
	}
}

func (task *Task) search(a []iceCream, m int) {
	mapRemaining := make(map[int]iceCream)
	for _, f := range a {
		remaining := m - f.Cost
		r, ok := mapRemaining[remaining]
		if ok {
			first := f.ID
			second := r.ID
			if first != second { // We want unique elements
				fmt.Printf("%d %d\n", int(math.Min(float64(first), float64(second))), int(math.Max(float64(first), float64(second))))
				break
			}
		}
		mapRemaining[f.Cost] = f
	}
}

func main() {
	NewTask().Read().Solve()
}
