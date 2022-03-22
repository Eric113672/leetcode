/*
* @Author: lishuang
* @Date:   2022/3/22 11:59
 */

package main

import (
	"math"
	"strconv"
)

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]

py 无脑解法：
class Solution:
    def printNumbers(self, n: int) -> List[int]:
        return list(range(1, 10**n))
*/

//解题思路
//1.固定第一位为1-9
//2.第二位以后可以是0-9
//3.递归剩下的n-1位
//出口：如果当前byte数组的下标为digit下标，表示该个数组已完成目的

func printNumbers2(n int) []int {
	res := make([]int, int(math.Pow10(n)-1))
	count := 0
	var dfs func(index int, num []byte, digit int)
	dfs = func(index int, num []byte, digit int) {
		if index == digit {
			tmp, _ := strconv.Atoi(string(num))
			res[count] = tmp
			count++
			return
		}
		for i := '0'; i <= '9'; i++ {
			num[index] = byte(i)
			dfs(index+1, num, digit)
		}
	}
	for digit := 1; digit <= n; digit++ {
		for first := '1'; first <= '9'; first++ {
			num := make([]byte, digit)
			num[0] = byte(first)
			dfs(1, num, digit)
		}
	}
	return res
}
