/*
* @Author: lishuang
* @Date:   2022/3/15 14:43
 */

package main

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

示例 1：
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1

解题思路:
使用动态规划的思想，一条长绳子可以拆成两段的子问题组合
公式：
dp[i] = max(dp[i-1]*dp[1], dp[i-2]*dp[2]......)
dp[i] = max((i-1) * 1, (i-2)*2......)

nannan

*/

// 动态规划
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1 //当绳子长为0， 1， 2的情况
	var temp1, temp2 int
	for i := 3; i < n+1; i++ {
		//将绳子分成i-1，1，两段，可以看出dp[i]必然大于等于dp[i-1]
		dp[i] = dp[i-1]
		for j := 1; j <= i/2; j++ {
			//将绳子分为两部分，两部分自由组合分段然后与不分段比较，选择较大的两部分组合相乘
			temp1 = max(dp[j], j)
			temp2 = max(dp[i-j], i-j)
			dp[i] = max(dp[i], temp1*temp2)
		}
	}
	return dp[n]
}
