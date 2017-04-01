package main

import (
	"bufio"
	"os"
	"log"
	"strconv"
	"fmt"
)

type Task struct {
	t int
	n []int
}

type empty struct {}

func (t *Task) checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func (t *Task) validateT(x int) bool {
	return x >= 1 && x <= 10
}


func (t *Task) validateN(x int) bool {
	return x >= 0 && x <= 60
}


func (t *Task) Read() *Task {
	scanner := bufio.NewScanner(os.Stdin)
	for i:=0; scanner.Scan(); i++ {
		s := scanner.Text()
		num, err := strconv.Atoi(s)
		t.checkErr(err)
		if i==0 {
            if ! t.validateT(num) {
				log.Fatalln("T value not in the valid range")
			}
			t.t = num
		} else {
			if ! t.validateN(num) {
				log.Fatalln("T value not in the valid range")
			}
			t.n = append(t.n, num)
		}
		if i == t.t {
			break
		}

	}
	t.checkErr(scanner.Err())
	return t
}

// 0 = 1
// 1 = (1 * 2) = 2
// 2 = (1 * 2) + 1 = 2 + 1
// 3 = ( (1 * 2) + 1 ) * 2 = 6
// 4 = ( (1 * 2) + 1 ) * 2) + 1 = 6 + 1

func (t *Task) getHeight(c int) int {
	height := 1
	i := 0
	for i < c {
		height *= 2
		i++
		if i == c {
			break
		}
		height++
		i++
		if i == c {
			break
		}
	}
	return height
}

func (t *Task) Run() {
	for _,n := range t.n {
		h := t.getHeight(n)
		fmt.Printf("%d\n", h)
	}
}

func NewTask() *Task {
	return &Task{}
}

func main() {
	task := NewTask()
	task.Read().Run()
}
