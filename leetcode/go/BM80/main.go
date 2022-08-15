package main

import "fmt"

func main() {

	prices := []int{8, 9, 2, 5, 4, 7, 1}
	dp := maxProfit(prices)
	fmt.Printf("dp %d", dp)
}

// 解题思路：
// 动态规划：只能买卖一次，记录历史上的最低价格min
// dp[i]=min(dp[i-1],min),由于只用到了i-1所以可以降维
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	min := prices[0]
	dp := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if dp < prices[i]-min {
			dp = prices[i] - min
		}
	}
	return dp
}
