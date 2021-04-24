package main

import (
	"fmt"
)

const targetLength = 7

var targetLetterOccurrences = map[rune]int{
	'B': 1,
	'A': 1,
	'L': 2,
	'O': 2,
	'N': 1,
}

func Solution(S string) int {
	mapLetterToNOccurrences := map[rune]int{}
	for _, letter := range S {
		// verify if it's part of target letter
		if _, isTargetLetter := targetLetterOccurrences[letter]; !isTargetLetter {
			continue
		}
		// count occurrences
		if _, exists := mapLetterToNOccurrences[letter]; !exists {
			mapLetterToNOccurrences[letter] = 1
		} else {
			mapLetterToNOccurrences[letter]++
		}
	}

	n := len(S) / targetLength
	for letter, occurrences := range mapLetterToNOccurrences {
		nCandidate := occurrences/targetLetterOccurrences[letter]
		if nCandidate < n {
			n = nCandidate
		}
	}
	return n
}

func main() {
	fmt.Println(Solution("BAONXXOLL") == 1)
	fmt.Println(Solution("BAOOLLNNOLOLGBAX") == 2)
	fmt.Println(Solution("QAWABAWONL") == 0)
	fmt.Println(Solution("ONLABLABLOON") == 1)
	fmt.Println(Solution("BALLOONBALLOONBALLOONBALLOONBALLOONBALLOONBALLOONBALLOONBALLOONBALLOON") == 10)
	fmt.Println(Solution("BALLOONBALLOONBALLOONBALLOONBALLOONBALOONBALLOONBALLOONBALLOONBALLOON") == 9)
	fmt.Println(Solution("BALOON") == 0)
}
