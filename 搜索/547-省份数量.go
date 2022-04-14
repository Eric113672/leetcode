/*
* @Author: lishuang
* @Date:   2022/4/14 17:03
 */

package main

/*
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。

这题跟 岛屿数量是一样的 只是需要将连接的 城市重新置为1
*/

func findCircleNum(isConnected [][]int) int {

	var ans = 0
	for i, row := range isConnected {
		for j, _ := range row {
			ans += cit2(isConnected, i, j)
		}
	}
	return ans
}

func cit2(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 0
	if y < len(grid) {
		for j := 0; j < len(grid[0]); j++ {
			if grid[y][j] == 1 {
				cit2(grid, y, j)
			}
		}
	}
	return 1
}
