package main

import (
	"fmt"
	"math"
)

func maxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k <= 0 {
		return 0
	}

	profit := make([][]int, k+1)
	profit[0] = make([]int, len(prices))

	for t := 1; t <= k; t++ {
		profit[t] = make([]int, len(prices))
		maxThusFar := math.MinInt64
		for d := range prices {
			// update max until now in previous transaction
			maxThusFar = maxInt(maxThusFar, -prices[d]+profit[t-1][d])

			if d == 0 {
				profit[t][d] = 0
			} else {
				profit[t][d] = maxInt(profit[t][d-1], prices[d]+maxThusFar)
			}
		}
	}

	return profit[len(profit)-1][len(profit[0])-1]
}

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}) == 7)
	fmt.Println(maxProfit(2, []int{}) == 0)
	fmt.Println(maxProfit(0, []int{5, 4, 3, 2, 1}) == 0)
	fmt.Println(maxProfit(2, []int{2, 4, 1}) == 2)
	fmt.Println(maxProfit(2, []int{5, 4, 3, 2, 1}) == 0)
}
