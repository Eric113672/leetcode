/*
* @Author: lishuang
* @Date:   2022/3/22 14:38
 */

package main

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
示例 1:

输入: [7,5,6,4]
输出: 5

精髓就是利用归并排序在合并过程中子数组的有序性。
至于如何利用，第 3 个题解有动画。
*/

var cnt int // 第一行

func reversePairs(nums []int) int {
	cnt = 0 // 第二行
	mergeSort(nums)
	return cnt
}

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	mid := n / 2

	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
}

func merge(leftNums, rightNums []int) []int {
	n1, n2 := len(leftNums), len(rightNums)
	result := make([]int, n1+n2)
	i, j, k := 0, 0, 0
	for i < n1 && j < n2 {
		if leftNums[i] <= rightNums[j] {
			result[k] = leftNums[i]
			i++
		} else {
			result[k] = rightNums[j]
			j++
			cnt += (n1 - i) // 第三行
		}
		k++
	}

	for i < n1 {
		result[k] = leftNums[i]
		i, k = i+1, k+1
	}

	for j < n2 {
		result[k] = rightNums[j]
		j, k = j+1, k+1
	}

	return result
}
