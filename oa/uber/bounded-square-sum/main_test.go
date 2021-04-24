package main

import (
	"github.com/luisfmelo/programming-challenges/tools"
	"testing"
)

func TestBinaryCount(t *testing.T) {
	type testCase struct {
		arr            []int
		target         int
		expectedOutput int
	}
	testCases := []testCase{
		{[]int{3, 4, 5, 5, 5, 7, 7}, 5, 3},
		{[]int{3, 4, 5, 5, 5, 7, 7, 8, 9, 9, 9, 10}, 5, 3},
		{[]int{3, 4, 5, 5, 5, 7, 7, 8, 9, 9, 9, 10}, 9, 3},
		{[]int{3, 4, 5, 5, 5, 7, 7, 8, 9, 9, 9, 10}, 6, 0},
		{[]int{3, 4, 5, 5, 5, 7, 7, 8, 9, 9, 9, 10}, 11, 0},
		{[]int{3, 4, 5, 5, 5, 7, 7, 8, 9, 9, 9, 10}, 1, 0},
	}

	for _, testCase := range testCases {
		actualOutput := tools.BinaryCount(testCase.arr, testCase.target, 0, len(testCase.arr)-1)
		if actualOutput != testCase.expectedOutput {
			t.Errorf("BinaryCount(): Wanted %v, Got %v", testCase.expectedOutput, actualOutput)
		}
	}
}
