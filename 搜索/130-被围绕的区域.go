/*
* @Author: lishuang
* @Date:   2022/4/14 17:26
 */

package main

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的'O' 用 'X' 填充。

解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的'O'都不会被填充为'X'。
任何不在边界上，或不与边界上的'O'相连的'O'最终都会被填充为'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

先把与边缘连接的o全部替换‘#’ 然后再将剩余的0 换成x 再把#换回0
*/

func solve(board [][]byte) {
	if len(board) <= 2 { // 不存在需要替换的O
		return
	}

	// 1. 遍历，从边缘找到O,然后dfs所有连接的O为#
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(i, 0, board)
		dfs(i, n-1, board)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, board)
		dfs(m-1, j, board)
	}

	// 2. 将剩余的O变为X，将剩余的#变成O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(i, j int, board [][]byte) {
	// terminal
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}

	// process
	board[i][j] = '#'

	// drill down
	dfs(i-1, j, board) // 上
	dfs(i+1, j, board) // 下
	dfs(i, j-1, board) // 左
	dfs(i, j+1, board) // 右
}
