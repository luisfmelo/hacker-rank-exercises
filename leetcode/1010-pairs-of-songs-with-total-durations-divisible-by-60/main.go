package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
	count := 0
	mapSongDurationMatch := map[int]bool{}

	for _, songTime := range time {
		timeRemain := songTime % 60
		var match bool
		if timeRemain == 0 {
			_, match = mapSongDurationMatch[60]
		} else {
			_, match = mapSongDurationMatch[60-timeRemain]
		}
		if match {
			count++
		}
		mapSongDurationMatch[60-timeRemain] = true
	}
	return count
}

func main() {
	//fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}) == 3)
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}) == 3)
	fmt.Println(numPairsDivisibleBy60([]int{0, 0, 60}) == 2)
	fmt.Println(numPairsDivisibleBy60([]int{0, 0, 0}) == 0)
}
