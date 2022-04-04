/*
* @Author: lishuang
* @Date:   2022/4/4 16:15
 */

package main

/*
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
示例 1:

输入: s = "aba"
输出: true

双指针:算法逻辑里有一个分叉，允许删除一个元素，那么在出现分叉的时候，贪心地计算左推进和右减小这两种方案，其中一种满足回文串即可。
即出现第一个不相等的地方时既可以向左进一 又可以右减一 集中一个满足即可
*/
func validPalindrome(s string) bool {
	isPalindrome := func(lo, hi int) bool {
		for lo < hi {
			if s[lo] != s[hi] {
				return false
			}

			lo++
			hi--
		}
		return true
	}

	for lo, hi := 0, len(s)-1; lo < hi; lo, hi = lo+1, hi-1 {
		if s[lo] != s[hi] {
			return isPalindrome(lo+1, hi) || isPalindrome(lo, hi-1)
		}

	}
	return true

}
