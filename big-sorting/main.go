package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func printArr(arr []string) {
	for _, v := range arr {
		fmt.Printf("%s\n", v)
	}
}

type BigSort []string

func (b BigSort) Len() int      { return len(b) }
func (b BigSort) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b BigSort) Less(i, j int) bool {
	iRunes := []rune(b[i])
	jRunes := []rune(b[j])
	if len(iRunes) == len(jRunes) {
		for idx := 0; idx < len(iRunes); idx++ {
			ri := iRunes[idx]
			rj := jRunes[idx]
			if rj == ri {
				continue
			}
			return ri < rj
		}
	} else {
		return len(iRunes) < len(jRunes)
	}
	return false
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)
	s.Scan()
	arr := make([]string, 0)
	for s.Scan() {
		var txt = strings.TrimSpace(s.Text())
		if txt != "" {
			arr = append(arr, s.Text())
		}
	}
	sort.Sort(BigSort(arr))
	printArr(arr)
}
