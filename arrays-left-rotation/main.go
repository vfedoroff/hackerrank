package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"math"
)

type Task struct {

	n int
	d int
	arr []int
}

func (t *Task) checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func (t *Task) setN(x string)  {
	i, err := strconv.Atoi(x)
	t.checkErr(err)

	if ! t.validateN(i) {
		log.Fatalln("Invalid param")
	}
	t.n = i
}

func (t *Task) setD(x string)  {
	i, err := strconv.Atoi(x)
	t.checkErr(err)

	if ! t.validateD(i) {
		log.Fatalln("Invalid param")
	}
	t.d = i
}

func (t *Task) validateN(x int) bool {
	return  x >= 1 && x <= int(math.Pow10(5))
}

func (t *Task) validateD(x int) bool {
	return  x >= 1 && x <= t.n
}

func (t *Task) validateA(x int) bool {
	return  x >= 1 && x <= int(math.Pow10(6))
}

func (t *Task) arrAdd(x string) {
	i, err := strconv.Atoi(x)
	t.checkErr(err)
	if ! t.validateA(i) {
		log.Fatalln("Invalid param")
	}
	t.arr = append(t.arr, i)

}

func (t *Task) Read() *Task {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for i:=0; scanner.Scan(); i++ {
		s := scanner.Text()
		switch i {
		case 0:
			{
				dn := strings.Split(s, " ")
				if len(dn) != 2 {
					log.Fatalln("Invalid arguments")
				}
				t.setN(dn[0])
				t.setD(dn[1])
			}
		case 1:
			{
				t.strToSlice(s)
			}
		}
		if i > 0 {
			break
		}
	}
	t.checkErr(scanner.Err())
	return t
}

func (t *Task) strToSlice(str string) {
	arr := strings.Split(str, " ")
	if len(arr) != t.n {
		log.Fatalln("Invalid array length")
	}
	for _, value := range arr {
		t.arrAdd(value)
	}
}

func (t *Task) Run() *Task {
	out := make([]int, len(t.arr))
	for i, _ := range t.arr {
		newPos := (i + (len(t.arr) - t.d)) % len(t.arr)
		out[newPos] = t.arr[i]
	}
	t.arr = out
	return t
}

func (t *Task) Print() {
	str := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(t.arr)), " "), "[]")
	fmt.Println(str)
}

func NewTask() *Task {
	return &Task{}
}


func main() {
	task := NewTask()
	task.Read().Run().Print()
}
