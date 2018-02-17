package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		var x1, v1, x2, v2 int
		fmt.Sscanf(str, "%d %d %d %d", &x1, &v1, &x2, &v2)
		if (x1 < x2) && (v1 < v2) {
			fmt.Println("NO")
		} else if (v1 != v2) && ((x2-x1)%(v1-v2)) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
