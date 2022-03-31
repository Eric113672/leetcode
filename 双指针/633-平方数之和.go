/*
* @Author: lishuang
* @Date:   2022/3/31 21:49
 */

package main

import (
	"fmt"
	"math"
)

/*
给定一个非负整数c，你要判断是否存在两个整数 a 和 b，使得a^2 + b^2 = c 。

示例 1：

输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5

还是使用双指针 但右边界是小于c的平方根的 勾股定理。。
*/

func judgeSquareSum(c int) bool {
	temp := math.Sqrt(float64(c))
	left := 0
	right := int(math.Floor(temp))

	for left <= right {
		if left*left+right*right == c {
			return true
		} else if left*left+right*right > c {
			right -= 1
		} else {
			left += 1
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(2))
}
