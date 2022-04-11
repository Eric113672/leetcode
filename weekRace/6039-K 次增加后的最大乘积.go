/*
* @Author: lishuang
* @Date:   2022/4/11 14:31
 */

package main

import (
	"container/heap"
	"sort"
)

/*
给你一个非负整数数组nums和一个整数k。每次操作，你可以选择nums中 任一元素并将它 增加1。

请你返回 至多k次操作后，能得到的nums的最大乘积。由于答案可能很大，请你将答案对10^9 + 7取余后返回。

示例 1：

输入：nums = [0,4], k = 5
输出：20
解释：将第一个数增加 5 次。
得到 nums = [5, 4] ，乘积为 5 * 4 = 20 。
可以证明 20 是能得到的最大乘积，所以我们返回 20 。
存在其他增加 nums 的方法，也能得到最大乘积。
*/

// 最大乘积可以理解为减小数组中最大值最小值的差异 即可能得到数组的乘积最大 每次都给最小值加1
//原地堆化可以做到 O(1) 的空间复杂度。

func maximumProduct(nums []int, k int) int {
	h := hp{nums}
	for heap.Init(&h); k > 0; k-- {
		h.IntSlice[0]++ // 每次给最小的加一
		heap.Fix(&h, 0)
	}
	ans := 1
	for _, num := range nums {
		ans = ans * num % (1e9 + 7)
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
