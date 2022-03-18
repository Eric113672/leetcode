package main

/*
输入:
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
*/
//  状态转移方程 f(x, y) = max(f(x-1,y)+grid[x][y],f(x,y-1)+grid[x][y])
func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && j > 0 {
				grid[i][j] += max(grid[i-1][j], grid[i][j-1])
			} else if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 && i > 0 {
				grid[i][j] += grid[i-1][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
