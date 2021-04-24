package main

import (
	"fmt"
	"github.com/luisfmelo/programming-challenges/tools"
)

var mapTelephoneKeys = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func backtrack(combs *[]string, digit rune) {
	var newCombs []string
	for _, comb := range *combs {
		for _, letter := range mapTelephoneKeys[digit] {
			newCombs = append(newCombs, comb + letter)
		}
	}
	*combs = newCombs
}

func letterCombinations(digits string) []string {
	var combinations []string
	if len(digits) == 0 {
		return combinations
	}
	combinations = mapTelephoneKeys[rune(digits[0])]

	for i := 1; i < len(digits); i++ {
		backtrack(&combinations, rune(digits[i]))
	}
	return combinations
}

func main() {
	var output, expectedOutput []string

	output = letterCombinations("23")
	expectedOutput = []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	fmt.Println(tools.CompareUnorderedStringArrays(expectedOutput, output))

	output = letterCombinations("")
	expectedOutput = []string{}
	fmt.Println(tools.CompareUnorderedStringArrays(expectedOutput, output))

	output = letterCombinations("2")
	expectedOutput = []string{"a", "b", "c"}
	fmt.Println(tools.CompareUnorderedStringArrays(expectedOutput, output))
}
