package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func quickSort(arr []int, lo, hi int) {
	//algorithm quicksort(A, lo, hi) is
	//if lo < hi then
	//p := partition(A, lo, hi)
	//quicksort(A, lo, p – 1)
	//quicksort(A, p + 1, hi)
	if lo < hi {
		var p = partition(arr, lo, hi)
		quickSort(arr, lo, p-1)
		quickSort(arr, p+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	//algorithm partition(A, lo, hi) is
	//pivot := A[hi]
	//i := lo - 1
	//for j := lo to hi - 1 do
	//if A[j] ≤ pivot then
	//	i := i + 1
	//	swap A[i] with A[j]
	//swap A[i + 1] with A[hi]
	//return i + 1
	var pivot = arr[hi]
	var i = lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	printArr(arr)
	return i + 1
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
	quickSort(arr, 0, len(arr)-1)
}
