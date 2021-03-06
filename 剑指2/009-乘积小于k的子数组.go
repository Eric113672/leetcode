/*
* @Author: lishuang
* @Date:   2022/3/25 21:32
 */

package main

/*
给定一个正整数数组nums和整数 k，请找出该数组内乘积小于k的连续的子数组的个数。

示例 1:

输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。


还是滑动窗口 一个经典的滑动窗口实现，需要注意 ans+=hi-lo+1 的意义
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	lo, hi := 0, 0
	var ans int
	product := 1

	for hi < len(nums) {
		product *= nums[hi]
		// 乘到大于k
		for lo <= hi && product >= k {
			product /= nums[lo]
			lo++
		}
		if lo <= hi {
			ans += hi - lo + 1
		}
		hi++
	}

	return ans
}
