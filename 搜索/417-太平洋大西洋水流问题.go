/*
* @Author: lishuang
* @Date:   2022/4/15 11:14
 */

package main

/*
有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。“太平洋”处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。

这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵heights，heights[r][c]表示坐标 (r, c) 上单元格 高于海平面的高度 。

岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。

返回 网格坐标 result的 2D列表 ，其中result[i] = [ri, ci]表示雨水可以从单元格 (ri, ci) 流向 太平洋和大西洋 。

看不懂。。
*/

func pacificAtlantic(matrix [][]int) [][]int {

	res := make([][]int, 0)
	if nil == matrix || len(matrix) == 0 {
		return res
	}
	// 行 高 流向方向
	row, col, dirs := len(matrix), len(matrix[0]), [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 行 高校验
	inArea := func(x, y int) bool {
		return x >= 0 && x < row && y >= 0 && y < col
	}

	var dfs func(x, y int, matrix [][]int, travelled [][]int)
	dfs = func(x, y int, matrix [][]int, travelled [][]int) {
		if travelled[x][y] == 1 {
			return
		}
		travelled[x][y] = 1
		for _, d := range dirs {
			newx := x + d[0]
			newy := y + d[1]
			if !inArea(newx, newy) || matrix[x][y] > matrix[newx][newy] || travelled[newx][newy] == 1 {
				continue
			}
			dfs(newx, newy, matrix, travelled)
		}
	}

	pcf, atl := make([][]int, row), make([][]int, row)
	for i := 0; i < row; i++ {
		pcf[i] = make([]int, col)
		atl[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		dfs(i, 0, matrix, pcf)
		dfs(i, col-1, matrix, atl)
	}
	for i := 0; i < col; i++ {
		dfs(0, i, matrix, pcf)
		dfs(row-1, i, matrix, atl)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pcf[i][j] == 1 && atl[i][j] == 1 {
				res = append(res, [][]int{{i, j}}...)
			}
		}
	}
	return res
}
