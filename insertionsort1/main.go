package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Print("\n")
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		arr[i], _ = strconv.Atoi(s.Text())
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			continue
		}
		v := arr[i+1]
		j := i
		for ; j > -1; j-- {
			if arr[j] > v {
				arr[j+1] = arr[j]
				printArr(arr)
			} else {
				break
			}
		}
		arr[j+1] = v
		printArr(arr)
		break
	}
}
