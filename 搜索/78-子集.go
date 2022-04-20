/*
* @Author: lishuang
* @Date:   2022/4/20 16:37
 */

package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/

var res2 [][]int

func subset(nums []int) [][]int {
	res2 = make([][]int, 0)
	sort.Ints(nums)
	Dfs([]int{}, nums, 0)
	return res2
}
func Dfs(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res2 = append(res2, tmp)
	for i := start; i < len(nums); i++ {
		//if i>start&&nums[i]==nums[i-1]{
		//	continue
		//}

		temp = append(temp, nums[i])
		Dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subset([]int{1, 2, 2}))
}
