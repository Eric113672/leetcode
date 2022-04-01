/*
* @Author: lishuang
* @Date:   2022/4/1 22:19
 */

package main

/*
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。

示例 1：

输入：s = "hello"
输出："holle"

双指针
*/

func reverseVowels(s string) string {
	b := []byte(s)
	l, r := 0, len(b)-1
	for l < r {
		for l < r && !isVowels(b[l]) {
			l++
		}
		for l < r && !isVowels(b[r]) {
			r--
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func isVowels(ch byte) bool {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
