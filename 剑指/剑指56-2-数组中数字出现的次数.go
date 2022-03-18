/*
* @Author: lishuang
* @Date:   2022/3/14 15:42
 */

package main

// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

func singleNumber(nums []int) int {
	h := map[int]int{}
	for _, v := range nums {
		h[v]++
	}
	for i, v := range h {
		if v == 1 {
			return i
		}
	}
	return 0
}
