/*
* @Author: lishuang
* @Date:   2022/4/6 16:47
 */

package main

import "container/heap"

/*
给定一个字符串 s ，根据字符出现的 频率 对其进行 降序排序 。一个字符出现的 频率 是它出现在字符串中的次数。

返回 已排序的字符串。如果有多个答案，返回其中任何一个。

示例 1:

输入: s = "tree"
输出: "eert"
解释: 'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。

*/

func frequencySort(s string) string {
	map_num := map[int32]int32{}
	//记录每个元素出现的次数
	for _, item := range s {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	//所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int32{key, value})
	}
	res := make([]rune, 0, len(s))
	//按顺序返回堆中的元素
	for h.Len() > 0 {
		temp := heap.Pop(h)
		for i := 0; i < int(temp.([2]int32)[1]); i++ {
			res = append(res, temp.([2]int32)[0])
		}
	}
	return string(res)
}

/*
type IHeap [][2]int32

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int32))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
*/
