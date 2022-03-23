/*
* @Author: lishuang
* @Date:   2022/3/23 10:54
 */

package main

import "strconv"

/*
给定两个 01 字符串a和b，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字1和0。

示例1:

输入: a = "11", b = "10"
输出: "101"

思路解析:
模拟二进制的计算过程，从最后位加起，有进位时向前加1。
注意运算逻辑2个点：

当前位 %2， 进位 /2
只需要一个 carry 作为 cursor, 以确保每次相同位置做计算
如果算法实现上先向后进位，最后逆序字符串的时候，可以采用双指针方法降低逆序的时间复杂度

*/

func addBinary(a string, b string) string {
	var ans string
	carry := 0 // cursor
	lenA, lenB := len(a), len(b)
	n := lenA
	if n < lenB {
		n = lenB
	}
	// 从最低位加起来  %2=1 则是1+0 %2=1 则是1+1 这位就是0 进位保留1继续运算 最后carry大于0则是在存在进位
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}

		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}
