/*
* @Author: lishuang
* @Date:   2022/4/14 16:25
 */

package main

/*
给你一个大小为 m x n 的二进制矩阵 grid 。

岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。

岛屿的面积是岛上值为 1 的单元格的数目。

计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
*/
// 每个点都dfs一下 递归去检查上下左右是否时1

type pair struct {
	x, y int
}

var dirs = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func maxAreaOfIsland(grid [][]int) int {
	var ans = 0
	for i, row := range grid {
		for j, _ := range row {
			ans = max(ans, getCnt(grid, i, j))
		}
	}
	return ans
}

func getCnt(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 0
	var res = 1
	for _, dir := range dirs {
		res += getCnt(grid, x+dir.x, y+dir.y)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
