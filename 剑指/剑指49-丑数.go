/*
* @Author: lishuang
* @Date:   2022/3/21 16:25
 */

package main

/*
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

*/

func nthUglyNumber(n int) int {
	// 丑数的提取性质：丑数只包含因子2，3，5，因此有“丑数=某较小丑数 * 某因子” （例如：10 = 5 * 2）
	// dp[a] dp[b] dp[c] 感觉想是三个数组组成一个丑数数组，分别对应小丑数2/3/5的倍数，每次都取最小倍数放到丑数数组
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n1, n2, n3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n1, n2), n3)
		if dp[i] == n1 {
			a++
		}
		if dp[i] == n2 {
			b++
		}
		if dp[i] == n3 {
			c++
		}
	}
	return dp[n-1]
}
