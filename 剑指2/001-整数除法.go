/*
* @Author: lishuang
* @Date:   2022/3/23 10:21
 */

package main

import (
	"math"
)

/*
给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%'。

注意：
整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8以及truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,2^31−1]。本题中，如果除法结果溢出，则返回 2^31 − 1

示例 1：
输入：a = 15, b = 2
输出：7
解释：15/2 = truncate(7.5) = 7
*/

//快速相除法
/*
func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	res := 0

	// 处理边界，防止转正数溢出
	// 除数绝对值最大，结果必为 0 或 1
	if b == math.MinInt32 {
		if a == b {
			return 1
		} else {
			return 0
		}
	}

	// 被除数先减去一个除数
	if a == math.MinInt32 {
		a -= -abs(b)
		res += 1
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	a = abs(a)
	b = abs(b)

	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			a = a - (b << i)
			// 代码优化：这里控制 ans 大于等于 INT_MAX
			if res > math.MaxInt32-(1<<i) {
				return math.MinInt32
			}
			res += 1 << i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
*/

func divide(a int, b int) int {
	sign := 1
	// 两数是否异号 异或运算
	if a^b < 0 {
		sign = -1
	}
	//fmt.Println(1 ^ -2)
	// 两数分别取正值
	if a < 0 {
		a = ^a + 1
	}
	if b < 0 {
		b = ^b + 1
	}
	res := 0
	// 按位减去 b的值
	for i := 31; i >= 0; i-- {
		if a>>i >= b {
			if i == 31 && sign == 1 {
				return math.MaxInt32
			}
			res += 1 << i
			a -= b << i
		}
	}
	if sign == -1 {
		return ^res + 1
	}
	return res
}
