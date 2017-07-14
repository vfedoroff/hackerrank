package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strToSlice(str string) (*[]int, error) {
	var result = make([]int, 0)
	var err error
	var item int
	str = strings.TrimSpace(str)
	for _, value := range strings.Split(str, " ") {
		item, err = strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return &result, nil
}

func countElements(arr *[]int, v int) int {
	var cnt int
	for _, i := range *arr {
		if i == v {
			cnt++
		}
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var text, _ = reader.ReadString('\n')
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	arr, _ := strToSlice(text)
	sort.Ints(*arr)
	var max = (*arr)[len(*arr)-1]
	fmt.Println(countElements(arr, max))
}
