/*
* @Author: lishuang
* @Date:   2022/4/22 10:25
 */

package main

/*
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用'.'表示。

*/
// 我是fw..
func solveSudoku(board [][]byte) {
	isValid := func(row, col int, val byte) bool {
		for i := 0; i < 9; i++ {
			if board[i][col] == val {
				return false
			}
		}
		for i := 0; i < 9; i++ {
			if board[row][i] == val {
				return false
			}
		}
		beginRow, beginCol := (row/3)*3, (col/3)*3
		for i := beginRow; i < beginRow+3; i++ {
			for j := beginCol; j < beginCol+3; j++ {
				if board[i][j] == val {
					return false
				}
			}
		}
		return true
	}
	var helper func() bool
	helper = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k)) {
						board[i][j] = byte(k)
						if helper() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	helper()
}
