/*
* @Author: lishuang
* @Date:   2022/4/22 10:40
 */

package main

/*
n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

*/

import "strings"

var res4 [][]string

func isValid(board [][]string, row, col int) (res4 bool) {
	n := len(board)
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if board[row][i] == "Q" {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

func backtrack8(board [][]string, row int) {
	size := len(board)
	if row == size {
		temp := make([]string, size)
		for i := 0; i < size; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res4 = append(res4, temp)
		return
	}
	for col := 0; col < size; col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		backtrack8(board, row+1)
		board[row][col] = "."
	}
}

func solveNQueens(n int) [][]string {
	res4 = [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backtrack8(board, 0)

	return res4
}
