/*
* @Author: lishuang
* @Date:   2022/3/15 14:25
 */

package main

/*
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]
解题思路:
整体思路，结果集中任何一个元素 = 其左边所有元素的乘积 * 其右边所有元素的乘积。
一轮循环构建左边的乘积并保存在结果集中，二轮循环 构建右边乘积的过程，乘以左边的乘积，并将最终结果保存


*/

func constructArr(a []int) []int {
	// initialization
	tmp := 1
	b := make([]int, len(a))
	for i := range b {
		b[i] = 1
	}
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1] // 下三角
	}
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1] // 上三角
		b[i] *= tmp
	}

	return b

}
