package main

import (
	"fmt"
	"math"
)

func scoreOfParentheses(S string) int {
	score := 0
	deep := 0

	for i := 1; i < len(S); i++ {
		if S[i] == '(' {
			deep++
		} else {
			if S[i-1] == '('{
				score += int(math.Pow(2, float64(deep)))
			}
			deep--
		}
	}
	return score
}

func main() {
	fmt.Println(scoreOfParentheses("()") == 1)
	fmt.Println(scoreOfParentheses("(())") == 2)
	fmt.Println(scoreOfParentheses("()()") == 2)
	fmt.Println(scoreOfParentheses("(()(()))") == 6)
}
