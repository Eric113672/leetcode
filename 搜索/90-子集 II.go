/*
* @Author: lishuang
* @Date:   2022/4/20 17:43
 */

package main

import "sort"

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。


示例 1：

输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

*/

var res3 [][]int

func subsetsWithDup(nums []int) [][]int {
	res3 = make([][]int, 0)
	sort.Ints(nums)
	dfs2([]int{}, nums, 0)
	return res3
}
func dfs2(temp, num []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)

	res3 = append(res3, tmp)
	for i := start; i < len(num); i++ {
		// 同一层的重复值 应被过滤掉
		if i > start && num[i] == num[i-1] {
			continue
		}
		temp = append(temp, num[i])
		dfs2(temp, num, i+1)
		temp = temp[:len(temp)-1]
	}
}
