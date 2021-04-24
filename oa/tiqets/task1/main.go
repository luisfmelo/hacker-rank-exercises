package main

import (
	"fmt"
	"strconv"
)

func Solution(A int, B int) int {
	binRep := strconv.FormatInt(int64(A*B), 2)

	n1s := 0
	for _, d := range binRep {
		if d == '1' {
			n1s++
		}
	}
	return n1s
}

func main() {
	fmt.Println(Solution(0, 0) == 0)
	fmt.Println(Solution(1, 1) == 1)
	fmt.Println(Solution(8, 1) == 1)
	fmt.Println(Solution(7,3) == 3)
	fmt.Println(Solution(100000000, 100000000) == 20)
}
