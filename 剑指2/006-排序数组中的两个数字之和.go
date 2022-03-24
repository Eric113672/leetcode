/*
* @Author: lishuang
* @Date:   2022/3/24 15:30
 */

package main

/*
给定一个已按照 升序排列的整数数组numbers ，请你从数组中找出两个数满足相加之和等于目标数target 。

函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 0开始计数 ，所以答案数组应当满足 0<= answer[0] < answer[1] <numbers.length。

假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。

示例 1：

输入：numbers = [1,2,4,6,10], target = 8
输出：[1,3]
解释：2 与 6 之和等于目标数 8 。因此 index1 = 1, index2 = 3 。
思路解析： 很简单的双指针

*/

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	lo, hi := 0, n-1

	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		if sum == target {
			return []int{lo, hi}
		} else if sum > target {
			hi--
		} else {
			lo++
		}
	}
	return nil
}
