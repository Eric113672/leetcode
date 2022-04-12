/*
* @Author: lishuang
* @Date:   2022/4/11 16:06
 */

package main

/*
给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母target，请你寻找在这一有序列表里比目标字母大的最小字母。

在比较时，字母是依序循环出现的。举个例子：

如果目标字母 target = 'z' 并且字符列表为letters = ['a', 'b']，则答案返回'a'

m = (l + h) / 2
m = l + (h - l) / 2
l + h 可能出现加法溢出，也就是说加法的结果大于整型能够表示的范围。但是 l 和 h 都为正数，因此 h - l 不会出现加法溢出问题。所以，最好使用第二种计算法方法。
示例 1：

输入: letters = ["c", "f", "j"]，target = "a"
输出: "c"
*/

func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters)-1
	for l <= h {
		m := l + (h-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	if l < len(letters) {
		return letters[l]
	} else {
		return letters[0]
	}
}
