/*
* @Author: lishuang
* @Date:   2022/4/18 15:15
 */

package main

import "fmt"

/*
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

*/
/*
func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] { // 寻找到匹配的第一个字符
				if search(i, j, board, word) {
					return true
				}
			}
		}
	}

	return false
}

func search(i, j int, board [][]byte, words string) bool {

	m, n := len(board), len(board[0])

	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != words[0] { // 不符合的条件
		return false
	}

	if len(words) == 1 { // 匹配到最后一个值， 返回
		return true
	}

	tmp := board[i][j]
	board[i][j] = '1' // 由于words只能是字母，所以'1'不会被匹配

	if search(i+1, j, board, words[1:]) || search(i, j+1, board, words[1:]) || search(i-1, j, board, words[1:]) || search(i, j-1, board, words[1:]) {

		return true

	} else {

		//注意由于board是slice引用类型，所以函数的修改会真正的修改原slice的值，所以需要重新改正回来
		board[i][j] = tmp
		return false
	}
}
*/

var ways = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	var dfs func(board [][]byte, x, y, i int, res *bool)

	dfs = func(board [][]byte, x, y, i int, res *bool) {
		if i == len(word) {
			*res = true
			return
		}
		if x < 0 || y < 0 || x >= m || y >= n || board[x][y] != word[i] {
			return
		}
		temp := board[x][y]
		board[x][y] = '0'
		for _, way := range ways {
			dfs(board, x+way[0], y+way[1], i+1, res)
		}
		board[x][y] = temp
	}
	res := false
	for j := 0; j < m; j++ {
		for k := 0; k < n; k++ {
			if board[j][k] == word[0] {
				dfs(board, j, k, 0, &res)
				if res == true {
					return res
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS"))
}
