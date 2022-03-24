/*
* @Author: lishuang
* @Date:   2022/3/24 10:20
 */

package main

import "sort"

/*
给你一个整数数组nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

示例 1：

输入：nums = [2,2,3,2]
输出：3
思路解析： 直接排序，然后找到只出现一次的元素
这样的做法 时间复杂度nlogn 空间复杂度O(1) 而hashmap 的则时空都是O(n)
*/

func singleNumber(nums []int) int {
	sort.Ints(nums)
	var cnt int
	cur := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == cur {
			cnt++
		} else {
			return nums[i-1]
		}

		if cnt == 3 {
			cur = nums[i+1]
			cnt = 0
		}
	}
	return nums[len(nums)-1]
}
