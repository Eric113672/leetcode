/*
* @Author: lishuang
* @Date:   2022/4/14 16:51
 */

package main

/*
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/
/*
type pair struct {
	x, y int
}

var dirs = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

*/

func numIslands(grid [][]byte) int {
	var ans = 0
	for i, row := range grid {
		for j, _ := range row {
			ans += cit(grid, i, j)
		}
	}
	return ans
}

func cit(grid [][]byte, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != '1' {
		return 0
	}
	grid[x][y] = '0'
	for _, dir := range dirs {
		cit(grid, x+dir.x, y+dir.y)
	}
	return 1
}
