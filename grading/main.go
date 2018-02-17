package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func round(grade int) int {
	return grade + (5 - grade%5)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			grade, _ := strconv.Atoi(scanner.Text())
			if (5 - grade%5) < 3 {
				if round(grade) > 38 {
					grade = round(grade)
				}
			}
			fmt.Println(grade)
		} else {
			return
		}
	}
}
