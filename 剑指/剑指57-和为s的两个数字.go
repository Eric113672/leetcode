/*
* @Author: lishuang
* @Date:   2022/3/4 11:08
 */

package main

/*
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
*/
func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	targetNums := make([]int, 2)
	for i < j {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			targetNums[0] = nums[i]
			targetNums[1] = nums[j]
			return targetNums
		}
	}
	return targetNums
}
