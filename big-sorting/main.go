package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
			ri := unicode.ToLower(iRunes[idx])
			rj := unicode.ToLower(jRunes[idx])
			if rj == ri {
				continue
			}
			return ri < rj
		}
	}
	return len(iRunes) < len(jRunes)
}

func main() {
	s := bufio.NewScanner(bufio.NewReader(os.Stdin))
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		s.Scan()
		var txt = strings.TrimSpace(s.Text())
		if len(txt) > 0 {
			arr[i] = s.Text()
		}
	}
	sort.Sort(BigSort(arr))
	printArr(arr)
}
