/*
* @Author: lishuang
* @Date:   2022/4/7 10:29
 */

package main

import "sort"

/*
给定一个区间的集合intervals，其中 intervals[i] = [starti, endi]。返回 需要移除区间的最小数量，使剩余区间互不重叠。

示例 1:

输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

思路解析：贪心留出后面的区间
先计算最多能组成的不重叠区间个数，然后用区间总个数减去不重叠区间的个数。

在每次选择中，区间的结尾最为重要，选择的区间结尾越小，留给后面的区间的空间越大，那么后面能够选择的区间个数也就越大。

按区间的结尾进行排序，每次选择结尾最小，并且和前一个区间不重叠的区间。

*/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { //将intervals按end大小递增排序
		return intervals[i][1] < intervals[j][1]
	})

	end := intervals[0][1]
	cnt := 1 //互相不重叠分组的数量
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end { //如果起始点 大于等于 end时，不重叠
			cnt++
			end = intervals[i][1] //更新end
		}
	}
	return len(intervals) - cnt //总数量 - 不重叠组数量 等于需要移除的数量
}
