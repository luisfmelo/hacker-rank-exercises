package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)
	target := 1
	for _, n := range A {
		if n > target {
			break
		}
		if n == target {
			target++
		}
	}
	return target
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3}) == 4)
	fmt.Println(Solution([]int{-1, -3}) == 1)
	fmt.Println(Solution([]int{2, 3, 4, 5, 6}) == 1)
	fmt.Println(Solution([]int{6, 5, 4, 3, 2, 1}) == 7)
	fmt.Println(Solution([]int{6, 5, 4, 3, 2, 2}) == 1)
	fmt.Println(Solution([]int{1, 1, 1, 1}) == 2)
	fmt.Println(Solution([]int{2, 2, 2, 2, 2}) == 1)
}
