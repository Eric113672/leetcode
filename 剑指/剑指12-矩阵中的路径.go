/*
* @Author: lishuang
* @Date:   2022/3/7 10:58
 */

/*

给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
主体还是采用dfs

思路解析:对于每一个可能的起点，进行dfs匹配即可。每层对四个方向进行搜索同时标记(x, y)访问过。

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

*/

package main

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	if n*m < len(word) {
		return false
	}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		// 终止条件
		if k >= len(word) {
			return true
		}
		// 终止条件，坐标范围错误或者 i j 坐标对应 board 的值不等于 k 坐标对应 word 的值
		if i < 0 || j < 0 || i >= n || j >= m || board[i][j] != word[k] {
			return false
		}
		// 如果往回前找，不会找到相同字符，如 ABAB，k = 2 时，往前往后都是 B ；
		// 将 B 修改（剪枝）为不存在的字符，杜绝往回找成功的可能性。
		board[i][j] = '0'
		res := dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i+1, j, k+1) || dfs(i-1, j, k+1)
		// 找完了就改回来
		board[i][j] = word[k]
		return res
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
