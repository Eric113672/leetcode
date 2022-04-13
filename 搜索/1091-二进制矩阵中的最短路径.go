/*
* @Author: lishuang
* @Date:   2022/4/13 11:42
 */

package main

import "math"

/*
给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。

二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求：

路径途经的所有单元格都的值都是 0 。
路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
畅通路径的长度 是该路径途经的单元格总数。

BFS:
广度优先搜索一层一层地进行遍历，每层遍历都是以上一层遍历的结果作为起点，遍历一个距离能访问到的所有节点。需要注意的是，遍历过的节点不能再次被遍历。
第一层：
0 -> {6,2,1,5}
第二层：
6 -> {4}
2 -> {}
1 -> {}
5 -> {3}
第三层：
4 -> {}
3 -> {}
每一层遍历的节点都与根节点距离相同。设 di 表示第 i 个节点与根节点的距离，推导出一个结论：
对于先遍历的节点 i 与后遍历的节点 j，有 di <= dj。利用这个结论，可以求解最短路径等 最优解 问题：第一次遍历到目的节点，其所经过的路径为最短路径。
应该注意的是，使用 BFS 只能求解无权图的最短路径，无权图是指从一个节点到另一个节点的代价都记为 1。

在程序实现 BFS 时需要考虑以下问题：

队列：用来存储每一轮遍历得到的节点；
标记：对于遍历过的节点，应该将它标记，防止重复遍历。

第一遍是用DFS写的，不出所望TLE了=。=。改用BFS就过了，方法是用一个队列记录当前搜索的位置，
并用一个vis的二维数组记录grid对应位置上已经通过搜索得到的最短路径，只有当可以更新这个最短路径时才需要将下一个需要搜索的节点加入队列，代码如下

*/

var dx = []int{1, 1, 1, 0, 0, -1, -1, -1}
var dy = []int{1, 0, -1, 1, -1, 1, 0, -1}

type node struct {
	x, y, l int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	vis := make([][]int, len(grid))
	for i := range vis {
		vis[i] = make([]int, len(grid[i]))
		for j := range vis[i] {
			vis[i][j] = math.MaxInt64
		}
	}
	vis[0][0] = 1
	queue := []node{{0, 0, 1}}
	for len(queue) > 0 {
		l := len(queue)
		for _, q := range queue {
			for i := range dx {
				x := q.x + dx[i]
				y := q.y + dy[i]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 1 {
					continue
				}
				if vis[x][y] > q.l+1 {
					vis[x][y] = q.l + 1
					queue = append(queue, node{x, y, q.l + 1})
				}
			}
		}
		queue = queue[l:]
	}
	if vis[len(grid)-1][len(grid[0])-1] == math.MaxInt64 {
		return -1
	}
	return vis[len(grid)-1][len(grid[0])-1]
}
