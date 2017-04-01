package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strings"
	"sort"
)

type Task struct {
	m int
	n int
	magazine []string
	note []string
	dict map[string]int

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

func (t *Task) validateRange(x int)  {
	if  x < 1 || x > 30000 {
		log.Fatalln("Invalid value")
	}
}

func (t *Task) validateArrLength(arr *[]string, length *int) {
	if len(*arr) != *length {
		log.Fatalln("Invalid input length")
	}
}

func (t *Task) validateWord(str *string){
	if len(*str) < 1 ||  len(*str) > 5{
		log.Fatalln("Invalid str length")
	}
}

func (t *Task) Read() *Task {
	var err error
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	_, err = fmt.Scanf("%d %d", &t.m, &t.n)
	t.checkErr(err)
	t.validateRange(t.m)
	t.validateRange(t.n)
	line, err = t.readLine(scanner)
	t.checkErr(err)
	t.magazine = strings.Split(line, " ")
	t.validateArrLength(&t.magazine, &t.m)
	line, err = t.readLine(scanner)
	t.checkErr(err)
	t.note = strings.Split(line, " ")
	t.validateArrLength(&t.note, &t.n)
	return t
}


func (t *Task) Solve() {
	if len(t.note) > len(t.magazine) {
		fmt.Println("No")
		return
	}
	for _, v := range t.note {
		t.validateWord(&v)
		i, ok := t.dict[v]
		if !ok || i == 0 {
			fmt.Println("No")
			return
		} else {
			t.dict[v]--
		}
	}
	fmt.Println("Yes")
}


func (t *Task) Run() *Task {
	t.dict = make(map[string]int)
	sort.Strings(t.magazine)
	for _, v := range t.magazine {
		_, ok := t.dict[v]
		if ! ok {
			t.validateWord(&v)
			t.dict[v] = 1
		} else {
			t.dict[v]++
		}
	}
	return t
}

func NewTask() *Task {
	return &Task{}
}

func main() {
	task := NewTask()
	task.Read().Run().Solve()
}
