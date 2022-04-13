/*
* @Author: lishuang
* @Date:   2022/4/13 14:31
 */

package main

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。


示例1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

*/

/*
可以将每个整数看成图中的一个节点，如果两个整数之差为一个平方数，那么这两个整数所在的节点就有一条边。

要求解最小的平方数数量，就是求解从节点 n 到节点 0 的最短路径。

本题也可以用动态规划求解，在之后动态规划部分中会再次出现。
*/

func numSquares(n int) int {
	squas := generateSquares(n)
	return bfs(squas, n)

}

func dp(squas []int, save []int, n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}

	res := n
	for i := len(squas) - 1; i >= 0; i-- {
		if n-squas[i] < 0 {
			continue
		}
		sub := save[n-squas[i]]
		if sub == 0 {
			sub = dp(squas, save, n-squas[i])
		}
		if 1+sub < res {
			res = 1 + sub
		}
	}
	save[n] = res

	return res
}

func dp2(squas []int, save []int, n int) {
	for i := 1; i <= n; i++ {
		if save[i] != 0 {
			continue
		}
		local := i
		for _, v := range squas {
			if i-v > 0 && i-v < i {
				if local > save[i-v]+1 {
					local = save[i-v] + 1
				}
			}
		}
		save[i] = local
	}
}

//这题用bfs没理解
func bfs(squas []int, n int) int {
	queue := make([]int, 0, n)
	queue = append(queue, n)
	steps := 0

	for len(queue) > 0 {
		preLen := len(queue)
		steps++
		visit := make([]int, queue[0]+1)
		for k := 0; k < preLen; k++ {
			q := queue[0]
			queue = queue[1:]
			for _, v := range squas {
				if q == v {
					return steps
				} else if q < v {
					continue
				}
				if visit[q-v] == 0 {
					visit[q-v] = 1
					queue = append(queue, q-v)
				}
			}
		}
	}
	return steps
}

/**
 * 生成小于 n 的平方数序列
 * @return 1,4,9,...
 */
func generateSquares(n int) []int {
	var squares []int
	square := 1
	diff := 2
	for square < n {
		squares = append(squares, square)
		square += diff
		diff += 2
	}
	return squares
}
