package main

import (
	"testing"
)

func TestHasElementsInOrder(t *testing.T) {
	type testCase struct {
		str1, str2 string
		expectedOutput bool
	}
	testCases := []testCase{
		{"ab", "abcd", true},
		{"bc", "abcd", true},
		{"ac", "abcd", true},
		{"ca", "abcd", false},
	}

	for _, testCase := range testCases {
		actualOutput := HasElementsInOrder(testCase.str1, testCase.str2)
		if actualOutput != testCase.expectedOutput {
			t.Errorf("HasElementsInOrder(): Wanted %v, Got %v", testCase.expectedOutput, actualOutput)
		}
	}
}