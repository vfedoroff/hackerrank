package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

type Task struct {
	lines []string
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

func (t *Task) getScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	return scanner
}

func (t *Task) readLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return scanner.Text(), nil
}

func (t *Task) Read() *Task {
	var err error
	var line string
	var n int
	scanner := t.getScanner()
	_, err = fmt.Scanf("%d", &n)
	t.checkErr(err)
	if !t.inRange(n, 1, int(math.Pow10(3))) {
		log.Fatalln("The value not in the expected range")
	}
	for i := 0; i < n; i++ {
		line, err = t.readLine(scanner)
		t.checkErr(err)
		if !t.inRange(len(line), 1, int(math.Pow10(3))) {
			log.Fatalln("The string length is not in the expected range")
		}
		t.lines = append(t.lines, line)
	}
	return t
}

func (t *Task) isBalanced(str string) bool {
	s := &stack{}
	for _, c := range str {
		switch {
		case c == '{':
			s.Push('}')
		case c == '[':
			s.Push(']')
		case c == '(':
			s.Push(')')
		default:
			if s.Empty() || c != s.Peek() {
				return false
			}
			s.Pop()
		}
	}
	return s.Empty()
}

func (t *Task) Solve() {
	for _, str := range t.lines {
		if t.isBalanced(str) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func NewTask() *Task {
	return &Task{}
}

func main() {
	task := NewTask()
	task.Read().Solve()
}
