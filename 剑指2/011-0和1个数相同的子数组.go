/*
* @Author: lishuang
* @Date:   2022/3/28 14:36
 */

package main

/*
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。

示例 1:

输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。

同样是前缀和
*/

func findMaxLength(nums []int) int {
	var res int
	var cnt int
	hashMap := make(map[int]int)
	hashMap[0]--
	for i, num := range nums {
		if num == 0 {
			cnt--
		} else {
			cnt++
		}
		if v, has := hashMap[cnt]; has {
			res = max(res, i-v)
		} else {
			hashMap[cnt] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
