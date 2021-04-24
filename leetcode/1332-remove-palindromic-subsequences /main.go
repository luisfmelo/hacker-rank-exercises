package main

import (
	"fmt"
	"log"
)

func isPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if isPalindrome(s) {
		return 1
	}
	return 2
}

func main() {
	type testCase struct {
		input    string
		expected int
	}

	testCases := []testCase{
		{"ababa", 1},
		{"abb", 2},
		{"baabb", 2},
		{"", 0},
		{"aabababba", 2},
	}

	for i, tc := range testCases {
		msg := fmt.Sprintf("Test #%d", i+1)
		actual := removePalindromeSub(tc.input)
		if actual != tc.expected {
			log.Fatalf("%s - Error: Expected %v, Got %v\n", msg, tc.expected, actual)
		}
		log.Printf("%s, Success\n", msg)
	}
}
