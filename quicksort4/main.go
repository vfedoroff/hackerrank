package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
		if i < len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

type QuickSort struct {
	arr   sort.IntSlice
	swaps int
}

func (q *QuickSort) GetSwaps() int {
	return q.swaps
}

func (q *QuickSort) Sort(arr sort.IntSlice) {
	q.arr = arr
	q.quicksort(0, q.arr.Len()-1)
}

func (q *QuickSort) quicksort(lo, hi int) {
	var mutex = sync.WaitGroup{}
	if lo < hi {
		var p = q.partition(lo, hi)
		mutex.Add(2)
		go func() {
			q.quicksort(lo, p-1)
			mutex.Done()
		}()
		go func() {
			q.quicksort(p+1, hi)
			mutex.Done()
		}()
		mutex.Wait()
	}
}

func (q *QuickSort) partition(lo, hi int) int {
	var pivot = q.arr[hi]
	var i = lo - 1
	for j := lo; j < hi; j++ {
		if q.arr[j] <= pivot {
			i = i + 1
			q.arr.Swap(i, j)
			q.swaps++
		}
	}
	q.arr.Swap(i+1, hi)
	q.swaps++
	return i + 1
}

type InsertionSort struct {
	arr   sort.IntSlice
	swaps int
}

func (s *InsertionSort) Sort(arr sort.IntSlice) {
	s.arr = arr
	s.insertionSort()
}

func (i *InsertionSort) GetSwaps() int {
	return i.swaps
}

func (s *InsertionSort) insertionSort() {
	for i := 1; i < s.arr.Len(); i++ {
		var x = s.arr[i]
		var j = i - 1
		for j >= 0 && s.arr[j] > x {
			s.swaps++
			s.arr.Swap(j+1, j)
			j--
		}
		s.arr[j+1] = x
	}
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)
	s.Scan()
	var n, _ = strconv.Atoi(s.Text())
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		s.Scan()
		var txt = strings.TrimSpace(s.Text())
		if txt != "" {
			var d, _ = strconv.Atoi(s.Text())
			arr = append(arr, d)
		}
	}
	var insertSortArray = make([]int, len(arr))
	copy(insertSortArray, arr)
	var q = QuickSort{}
	var i = InsertionSort{}
	var mutex = sync.WaitGroup{}
	mutex.Add(2)
	go func() {
		q.Sort(sort.IntSlice(arr))
		mutex.Done()
	}()
	go func() {
		i.Sort(sort.IntSlice(insertSortArray))
		mutex.Done()
	}()
	mutex.Wait()
	fmt.Println(i.GetSwaps() - q.GetSwaps())
}
