package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(arr *[]int) int {
	var sum int
	for _, v := range *arr {
		sum += v
	}
	return sum
}

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	arr, _ := strToSlice(text)
	sort.Ints(*arr)
	var min = (*arr)[0]
	var max = (*arr)[len(*arr)-1]
	var asum = sum(arr)
	var minimum = asum - max
	var maximum = asum - min
	fmt.Println(minimum, maximum)
}
