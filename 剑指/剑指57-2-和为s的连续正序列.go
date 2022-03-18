/*
* @Author: lishuang
* @Date:   2022/3/15 15:00
 */

package main

/*
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]

*/
//双指针
//解题思路
//注意题目要求至少两个数字组成，所以左指针的边界是 [1,target/2]

func findContinuousSequence(target int) [][]int {
	lo, hi, sum := 1, 1, 0
	res := make([][]int, 0)

	for lo <= target>>1 {
		if sum < target {
			// move right cursor and increase sum
			sum += hi
			hi++
		} else if sum > target {
			// move left cursor and reduce sum
			sum -= lo
			lo++
		} else {
			get := make([]int, 0)
			for i := lo; i < hi; i++ {
				get = append(get, i)
			}
			res = append(res, get)

			// move left cursor and reduce got situation
			sum -= lo
			lo++
		}

	}

	return res

}
