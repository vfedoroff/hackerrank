package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
)

type Task struct {
	size []int32
	word string
}

type validator func(int64) bool


func (t *Task) checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func (t *Task) Read() *Task {
	scanner := bufio.NewScanner(os.Stdin)
	for i:=0; scanner.Scan(); i++ {
		s := scanner.Text()
		switch i {
		case 0:
			{
				t.strToSlice(s, func(x int64) bool { return i >= 1 || i <= 7 })
			}
		case 1:
			{
				t.word = s
			}
		}
		if i > 0 {
			break
		}
	}
	t.checkErr(scanner.Err())
	return t
}


func (t *Task) strToSlice(str string, validate validator) {
	for _, value := range strings.Split(str, " ") {
		h, err := strconv.ParseInt(value, 10, 32)
		t.checkErr(err)
		if ! validate(h) {
			os.Exit(1)
		}
		t.size = append(t.size, int32(h))
	}
}

func (t *Task) Run() {
	var max int32
	done := make(chan bool, 1)
	if len(t.size) != 26 {
		log.Fatalf("Array len is not equal to 26\n")
	}
	if len(t.word) > 10 {
		log.Fatalf("Word length should be less then 10\n")
	}
	go func() {
		for _, c := range t.word {
			if c < rune('a') || c > rune('z') {
				os.Exit(1)
			}
			s := t.size[c-rune('a')]
			if s > max {
				max = s
			}
		}
		done <- true
	}()
	<-done
	fmt.Println(int(max)*len(t.word))
}

func NewTask() *Task {
	return &Task{}
}


func main() {
	task := NewTask()
	task.Read().Run()
}
