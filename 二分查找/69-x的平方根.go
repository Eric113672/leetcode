/*
* @Author: lishuang
* @Date:   2022/4/11 15:46
 */

package main

/*
给你一个非负整数 x ，计算并返回x的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。

示例 1：

输入：x = 4
输出：2

一个数 x 的开方 sqrt 一定在 0 ~ x 之间，并且满足 sqrt == x / sqrt。可以利用二分查找在 0 ~ x 之间查找 sqrt。

对于 x = 8，它的开方是 2.82842...，最后应该返回 2 而不是 3。
在循环条件为 l <= h 并且循环退出时，h 总是比 l 小 1，也就是说 h = 2，l = 3，因此最后的返回值应该为 h 而不是 l。
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 0, x
	for l <= h {
		mid := l + (h-l)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if mid > sqrt {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return h
}
