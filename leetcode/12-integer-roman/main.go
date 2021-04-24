package main

import (
	"fmt"
	"log"
	"strings"
)

func intToRoman(num int) string {
	var mapRomanLetters = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	var roman []string
	powerOf10 := 1
	for num > 0 {
		n := num % (10 * powerOf10)

		// Convert to Roman
		possibilities := map[int]string{
			0:             "",
			powerOf10:     mapRomanLetters[powerOf10],
			2 * powerOf10: mapRomanLetters[powerOf10] + mapRomanLetters[powerOf10],
			3 * powerOf10: mapRomanLetters[powerOf10] + mapRomanLetters[powerOf10] + mapRomanLetters[powerOf10],
			4 * powerOf10: mapRomanLetters[powerOf10] + mapRomanLetters[5*powerOf10],
			5 * powerOf10: mapRomanLetters[5*powerOf10],
			6 * powerOf10: mapRomanLetters[5*powerOf10] + mapRomanLetters[powerOf10],
			7 * powerOf10: mapRomanLetters[5*powerOf10] + mapRomanLetters[powerOf10] + mapRomanLetters[powerOf10],
			8 * powerOf10: mapRomanLetters[5*powerOf10] + mapRomanLetters[powerOf10] + mapRomanLetters[powerOf10] + mapRomanLetters[powerOf10],
			9 * powerOf10: mapRomanLetters[powerOf10] + mapRomanLetters[10*powerOf10],
		}

		roman = append([]string{possibilities[n]}, roman...)

		num -= n
		powerOf10 *= 10
	}

	return strings.Join(roman, "")
}

func intToRoman2(num int) string {
	var romanLetters = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var integers = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var roman []string
	for i := 0; num > 0; i++ {
		if integers[i] <= num {
			roman = append(roman, romanLetters[i])
			num -= integers[i]
		}
		if integers[i] <= num {
			i--
		}
	}

	return strings.Join(roman, "")
}

func main() {
	testCases := []struct {
		input          int
		expectedOutput string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i, tc := range testCases {
		msg := fmt.Sprintf("Test #%d", i+1)
		actual := intToRoman(tc.input)
		if actual != tc.expectedOutput {
			log.Printf("%s - Error: Expected %v, Got %v\n", msg, tc.expectedOutput, actual)
		} else {
			log.Printf("%s - Success\n", msg)
		}
	}

	log.Println()
	log.Println()

	for i, tc := range testCases {
		msg := fmt.Sprintf("Test #%d", i+1)
		actual := intToRoman2(tc.input)
		if actual != tc.expectedOutput {
			log.Printf("%s - Error: Expected %v, Got %v\n", msg, tc.expectedOutput, actual)
		} else {
			log.Printf("%s - Success\n", msg)
		}
	}
}
