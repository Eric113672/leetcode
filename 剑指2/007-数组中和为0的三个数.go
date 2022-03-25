/*
* @Author: lishuang
* @Date:   2022/3/25 16:40
 */

package main

import "sort"

/*
给定一个包含 n 个整数的数组nums，判断nums中是否存在三个元素a ，b ，c ，使得a + b + c = 0 ？请找出所有和为 0 且不重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

排序加双指针
*/

func threeSum(nums []int) [][]int {
	nLen := len(nums)
	if nLen < 3 {
		return nil
	}

	res := make([][]int, 0, len(nums))
	sort.Ints(nums)

	for i := range nums {
		// 解决相同的边界问题
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo, hi := i+1, nLen-1 // 初始化指针
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})

				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				lo++
				hi--
			} else if sum > 0 {
				hi--
			} else {
				lo++
			}
		}
	}
	return res

}
