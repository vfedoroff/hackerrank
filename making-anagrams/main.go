package main

import (
	"log"
	"bufio"
	"os"
	"math"
	"fmt"
)

type Task struct {
	a string
	b string
	res int
}

func (t *Task) checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func (t *Task) readLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return scanner.Text(), nil
}

func (t *Task) Read() *Task {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	t.checkErr(scanner.Err())
	t.a, err = t.readLine(scanner)
	t.checkErr(err)
	t.b, err = t.readLine(scanner)
	t.checkErr(err)
	return t
}

func (t *Task) Print() {
	fmt.Printf("%d\n", t.res)
}


func (t *Task) Run() *Task {
	cnt := make([]int, 26)
	for _, s := range t.a {
		cnt[s - rune('a')]++
	}
	for _, s := range t.b {
		cnt[s - rune('a')]--
	}
	for _, i := range cnt {
		t.res += int(math.Abs(float64(i)))
	}
	return t
}

func NewTask() *Task {
	return &Task{}
}

func main() {
	task := NewTask()
	task.Read().Run().Print()
}
