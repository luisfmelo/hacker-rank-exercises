package main

import "fmt"

func findLongestWord2(s string, d []string) string {
	mapDictWordToPointer := map[string]int{}
	for _, word := range d {
		mapDictWordToPointer[word] = 0
	}

	longestWord := ""
	for _, l := range s {
		for word, ptr := range mapDictWordToPointer {
			if rune(word[ptr]) == l {
				mapDictWordToPointer[word]++
			}
			if mapDictWordToPointer[word] == len(word) {
				if len(word) > len(longestWord) || len(word) == len(longestWord) && word < longestWord {
					longestWord = word
				}
				delete(mapDictWordToPointer, word)
			}
		}
	}
	return longestWord
}

// All elements of str1 appear in str2 in order
// example: str1: "ab", str2: "abcd", output: true
func HasElementsInOrder(str1, str2 string) bool {
	ptr := 0
	for _, c := range str2 {
		if c == rune(str1[ptr]) {
			ptr++
		}
		if ptr == len(str1) {
			return true
		}
	}
	return ptr == len(str1)
}

func findLongestWord(s string, d []string) string {
	longestWord := ""
	for _, word := range d {
		if HasElementsInOrder(word, s) {
			if len(word) > len(longestWord) || len(word) == len(longestWord) && word < longestWord {
				longestWord = word
			}
		}
	}
	return longestWord
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}) == "apple")
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}) == "a")
	fmt.Println(findLongestWord("aaa", []string{"aaa", "aa", "a"}) == "aaa")
	fmt.Println(findLongestWord("bab", []string{"ba", "ab", "a", "b"}) == "ab")
}
