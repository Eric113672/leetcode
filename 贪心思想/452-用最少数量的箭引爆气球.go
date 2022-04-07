/*
* @Author: lishuang
* @Date:   2022/4/7 10:53
 */

package main

import "sort"

/*
有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组points，
其中points[i] = [xstart, xend]表示水平直径在xstart和xend之间的气球。你不知道气球的确切 y 坐标。

一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足 xstart≤ x ≤ xend，
则该气球会被 引爆。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。

给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数。

示例 1：

输入：points = [[10,16],[2,8],[1,6],[7,12]]
输出：2
解释：气球可以用2支箭来爆破:
-在x = 6处射出箭，击破气球[2,8]和[1,6]。
-在x = 11处发射箭，击破气球[10,16]和[7,12]。
也是计算不重叠的区间个数，不过和 Non-overlapping Intervals 的区别在于，[1, 2] 和 [2, 3] 在本题中算是重叠区间。

*/

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { //将points按start大小递增排序
		return points[i][1] < points[j][1]
	})
	cnt := 1
	end := points[0][1]
	for i := 0; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		} else {
			cnt++
			end = points[i][1]
		}
	}
	return cnt
}
