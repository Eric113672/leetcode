/*
* @Author: lishuang
* @Date:   2022/4/12 15:15
 */

package main

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
进阶：

你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

可以用二分查找找出第一个位置和最后一个位置，但是寻找的方法有所不同，需要实现两个二分查找。
我们将寻找 target 最后一个位置，转换成寻找 target+1 第一个位置，再往前移动一个位置。这样我们只需要实现一个二分查找代码即可。

在寻找第一个位置的二分查找代码中，需要注意 h 的取值为 nums.length，而不是 nums.length - 1。先看以下示例：
nums = [2,2], target = 2

如果 h 的取值为 nums.length - 1，那么 last = findFirst(nums, target + 1) - 1 = 1 - 1 = 0。
这是因为 findLeft 只会返回 [0, nums.length - 1] 范围的值，对于 findFirst([2,2], 3) ，我们希望返回 3 插入 nums 中的位置，也就是数组最后一个位置再往后一个位置，
即 nums.length。所以我们需要将 h 取值为 nums.length，从而使得 findFirst返回的区间更大，能够覆盖 target 大于 nums 最后一个元素的情况。

*/

func searchRange(nums []int, target int) []int {
	first, last := findFirst(nums, target), findFirst(nums, target+1)-1
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	return []int{first, max(first, last)}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func findFirst(nums []int, target int) int {
	l, h := 0, len(nums) // h的初始值
	for l < h {
		m := l + (h-l)>>1
		if nums[m] >= target {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}
