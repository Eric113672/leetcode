/*
* @Author: lishuang
* @Date:   2022/4/8 10:51
 */

package main

/*
给定一个数组 prices ，其中prices[i] 表示股票第 i 天的价格。

在每一天，你可能会决定购买和/或出售股票。你在任何时候最多只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
返回 你能获得的 最大 利润。

示例 1:

输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3


思路解析:
对于 [a, b, c, d]，如果有 a <= b <= c <= d ，
那么最大收益为 d - a。而 d - a = (d - c) + (c - b) + (b - a) ，
因此当访问到一个 prices[i] 且 prices[i] - prices[i-1] > 0，那么就把 prices[i] - prices[i-1] 添加到收益中。

*/
func maxProfit2(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit = profit + prices[i] - prices[i-1]
		}
	}
	return profit
}
