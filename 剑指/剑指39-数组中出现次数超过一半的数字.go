/*
* @Author: lishuang
* @Date:   2022/3/15 11:40
 */

package main

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

投票法
*/

func majorityElement(nums []int) int {
	ans, votes := nums[0], 1
	for i := 1; i < len(nums); i++ {
		// fmt.Println(nums[i], votes)
		if nums[i] == ans {
			votes++
		} else { // first case
			votes--
		}

		if votes == 0 {
			ans = nums[i]
			votes = 1
		}
	}

	return ans
}
