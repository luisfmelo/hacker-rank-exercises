package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	LAND  = 0
	WATER = 1
)

func longestNiceSubstring(s string) string {
	mapLettersToIndexes := map[string][]int{}

	for index, char := range s {
		letter := string(char)
		if indexes, exists := mapLettersToIndexes[letter]; !exists {
			mapLettersToIndexes[letter] = []int{index}
		} else {
			mapLettersToIndexes[letter] = append(indexes, index)
		}
	}

	var ptrEnd = 0
	var longestNiceSubstring = ""
	for i, _ := range s {
		ptrEnd = len(s) - 1

		for j := i; j < len(s); j++ {
			char := s[j]
			letter := string(char)
			var brotherLetter string
			if char >= 'a' && char <= 'z' {
				brotherLetter = strings.ToUpper(letter)
			} else if char >= 'A' && char <= 'Z' {
				brotherLetter = strings.ToLower(letter)
			}

			brotherIndexes, brotherExists := mapLettersToIndexes[brotherLetter]
			if !brotherExists && j < ptrEnd {
				if ptrEnd == j && len(longestNiceSubstring) < j-i {
					longestNiceSubstring = s[i:j]
				}
				break
			}

			sort.Ints(brotherIndexes)
			foundBrother := false
			for _, botherIdx := range brotherIndexes {
				if botherIdx <= ptrEnd {
					foundBrother = true
					break
				}
			}
			if !foundBrother {
				break
			}
			if j == ptrEnd {
				if len(longestNiceSubstring) < ptrEnd-i+1 {
					longestNiceSubstring = s[i : ptrEnd+1]
				}
				break
			}
		}

	}

	return longestNiceSubstring
}

func canChoose(groups [][]int, nums []int) bool {
return true
}

func highestPeak(isWater [][]int) [][]int {
	var heights = make([][]int, len(isWater))

	for i := range isWater {
		for j := range isWater[i] {
			if isWater[i][j] == WATER {
				heights[i][j] = 0

			}
		}
	}

	return heights
}

func main() {
	//fmt.Printf("1: %s\n", longestNiceSubstring("YazaAay"))
	//fmt.Printf("2: %s\n", longestNiceSubstring("bB"))
	//fmt.Printf("3: %s\n", longestNiceSubstring("c"))
	//fmt.Printf("4: %s\n", longestNiceSubstring("dDzeE"))

	fmt.Printf("1: %v\n", highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}))
}
