/*
* @Author: lishuang
* @Date:   2022/3/25 18:30
 */

package main

import "math"

/*
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。
示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组[4,3]是该条件下的长度最小的子数组。

题意理解错了。。
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	// 判断总和是否大于target
	for _, v := range nums {
		total += v
	}
	if total < target {
		return 0
	}
	if total == target {
		return len(nums)
	}
	// 先排序 逆序寻找首个小于target的值
	sort.Ints(nums)
	fmt.Println(nums)
	j := len(nums) - 1
	for ; j >= 0; j-- {
		if nums[j] <= target {
			break
		}
	}
	total = 0
	totalValue := 0
	for ; j >= 0; j-- {
		totalValue += nums[j]
		total++
		if totalValue >= target {
			return total
		}
	}
	return 0
}
*/
//思路解析: 注意题目所说的连续子数组并非是一个等差数列， 只需要保证索引连续即可。所以很明显的做法是滑动窗口。
//分为 lo 和 hi 下标，左坐标右移，sum 变小；右坐标右移， sum 变大。默认右坐标 hi 移动，当 sum==target 时，取最小索引长度差值

func minSubArrayLen(target int, nums []int) int {
	lo, hi := 0, 0
	ret := math.MaxInt32
	sum := 0

	for hi < len(nums) {
		sum += nums[hi]
		for sum >= target {
			ret = min(ret, hi-lo+1)
			sum -= nums[lo]
			lo++
		}
		hi++
	}
	if ret == math.MaxInt32 {
		return 0
	}

	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
