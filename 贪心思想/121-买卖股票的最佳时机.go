/*
* @Author: lishuang
* @Date:   2022/4/8 10:10
 */

package main

/*
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

贪心: 只要记录前面的最小价格，将这个最小价格作为买入价格，然后将当前的价格作为售出价格，查看当前收益是不是最大收益。

*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	minBuy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		} else {
			profit = max(profit, prices[i]-minBuy)
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
