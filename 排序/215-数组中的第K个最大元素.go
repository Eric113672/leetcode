/*
* @Author: lishuang
* @Date:   2022/4/6 11:22
 */

package main

import (
	"container/heap"
	"sort"
)

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

*/
// 方法1 直接排序 返回第k个  时间复杂度 O(NlogN)，空间复杂度 O(1)
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 方法2  堆 ：时间复杂度 O(NlogK)，空间复杂度 O(K)。

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// 大顶堆版
func findKthLargest2(nums []int, k int) int {
	a := IntHeap(nums)
	h := &a
	heap.Init(h)
	var res interface{}
	for k > 0 {
		res = heap.Pop(h)
		k--
	}
	return res.(int)
}
