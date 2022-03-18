package main

/*
示例 1：

输入：n = 2
输出：2
示例 2：

输入：n = 7
输出：21
示例 3：

输入：n = 0
输出：1
*/

func numWays(n int) int {
	a, b := 1, 0
	for i := 0; i < n; i++ {
		a, b = (a+b)%1000000007, a
	}
	return a
}
