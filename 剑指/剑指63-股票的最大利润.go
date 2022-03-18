package main

import "math"

/*
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

*/

// 贪心算法解析 大致就是记录目前遇到的最小值 用当前的股票价减去low 就是当前的最大利润
func maxProfit(prices []int) int {
	low := math.MaxInt32
	result := 0

	for i := 0; i < len(prices); i++ {
		low = min(low, prices[i])
		result = max(result, prices[i]-low)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//动态规划算法解析 其实和贪心算法大致是差不多的 主体就是讲暂时的low存起来dp数组里 用当前的股票价减去当前的low与上一个[1]对比
/*
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}
*/
