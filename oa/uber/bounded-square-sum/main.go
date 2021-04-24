package main

import (
	"fmt"
	"math"
	"sort"
)

// naive approach O(n * m)
func boundedSquareSum2(a, b []int, lower, upper int) int {
	n := 0
	for _, aElem := range a {
		for _, bElem := range b {
			v := aElem*aElem + bElem*bElem
			if v >= lower && v <= upper {
				n++
			}
		}
	}
	return n
}


func boundedSquareSum(a, b []int, lower, upper int) int {
	n := 0
	sort.Ints(b)
	for _, elem := range a {
		newLower := math.Ceil(math.Sqrt(float64(lower - elem*elem)))
		newUpper := math.Floor(math.Sqrt(float64(upper - elem*elem)))
		print(newLower)
		print(newUpper)
		//n += BinaryCount(b, int(newUpper), 0, len(b)-1) - BinaryCount(b, int(newLower), 0, len(b)-1)
	}
	return n
}

func main() {
	fmt.Println(boundedSquareSum([]int{3, -1, 9}, []int{100, 5, -2}, 7, 99) == 4)
	fmt.Println(boundedSquareSum([]int{1, 2, 3, -1, -2, -3}, []int{10}, 0, 100) == 0)
}
