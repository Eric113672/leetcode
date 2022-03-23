/*
* @Author: lishuang
* @Date:   2022/3/23 11:19
 */

package main

/*
给定一个非负整数 n，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

示例 1:

输入: n = 2
输出: [0,1,1]
解释:
0 --> 0
1 --> 1
2 --> 10
动态转移方程： ret[i] = ret[i>>1]+ (i&1)
*/

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ret[i] = ret[i>>1] + (i & 1)
	}
	return ret
}
