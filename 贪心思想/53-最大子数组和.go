/*
* @Author: lishuang
* @Date:   2022/4/8 14:54
 */

package main

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
dp
*/

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		if i == 0 {
			maxSum = dp[i]
			continue
		}
		if dp[i-1]+dp[i] > dp[i] {
			dp[i] += dp[i-1]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}
