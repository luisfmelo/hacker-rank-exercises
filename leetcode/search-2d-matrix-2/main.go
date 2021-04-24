package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for col >= 0 && row <= len(matrix)-1 {
		if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		} else if matrix[row][col] == target {
			return true
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if found := binarySearch(target, row); found {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("\nApproach 1)")
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5) == true)
	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20) == false)

	fmt.Println("\nApproach 2) Slower")
	fmt.Println(searchMatrix2([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5) == true)
	fmt.Println(searchMatrix2([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20) == false)
}
