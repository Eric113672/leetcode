/*
* @Author: lishuang
* @Date:   2022/4/22 15:06
 */

package main

import "fmt"

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
*/

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	res1 := dp1(0, len(nums)-2, nums)
	res2 := dp1(1, len(nums)-1, nums)
	return max(res1, res2)

}

func dp1(start, end int, nums []int) int {
	if end == start {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])

	for i := start + 2; i < end+1; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	//fmt.Println(dp)
	return dp[end]
}

func main() {
	fmt.Println(rob2([]int{1, 2, 1, 1}))
}
