package main

import (
	"fmt"
	"reflect"
	"sort"
)

func nextPermutation(nums []int) {
	lastPositiveDerIndex := -1
	toSwapIndex := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			lastPositiveDerIndex = i - 1
			toSwapIndex = i
		} else if lastPositiveDerIndex != -1 && nums[i]-nums[lastPositiveDerIndex] > 0 && nums[toSwapIndex]-nums[lastPositiveDerIndex] > nums[i]-nums[lastPositiveDerIndex] {
			toSwapIndex = i
		}
	}

	if lastPositiveDerIndex != -1 {
		nums[lastPositiveDerIndex], nums[toSwapIndex] = nums[toSwapIndex], nums[lastPositiveDerIndex]
	}
	sort.Ints(nums[lastPositiveDerIndex+1:])
}

func main() {
	var nums, expectedNums []int

	nums = []int{2, 3, 1, 3, 3}
	expectedNums = []int{2, 3, 3, 1, 3}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{1, 2, 3}
	expectedNums = []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{3, 2, 1}
	expectedNums = []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{1, 1, 5}
	expectedNums = []int{1, 5, 1}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{1}
	expectedNums = []int{1}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{1, 3, 2}
	expectedNums = []int{2, 1, 3}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{1, 3, 2, 1}
	expectedNums = []int{2, 1, 1, 3}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{2, 3, 1}
	expectedNums = []int{3, 1, 2}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))

	nums = []int{1, 2, 3, 2, 4, 3, 2, 1}
	expectedNums = []int{1, 2, 3, 3, 1,2,2,4}
	nextPermutation(nums)
	fmt.Println(reflect.DeepEqual(nums, expectedNums))
}
