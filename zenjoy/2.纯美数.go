/*
* @Author: lishuang
* @Date:   2022/4/22 20:51
 */

package main

/*
思路解析  美数只包含因子2，3，5, 则可以理解为 美数=某较小美数 * 某因子（例如：10 = 5 * 2）
则有 dp[i] = min(min(n1, n2), n3) n1 n2 n3 分别是当前倍数的美数 最小值 则得到得dp数组是有序的
*/

func nthBeautifulNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1 // 初始化dp数组
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
