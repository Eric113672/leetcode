/*
* @Author: lishuang
* @Date:   2022/4/5 22:11
 */

package main

import (
	"fmt"
)

/*
给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回dictionary 中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。

如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。

示例 1：

输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
输出："apple"

时间复杂度高 空间低
通过删除字符串 s 中的一个字符能得到字符串 t，可以认为 t 是 s 的子序列，我们可以使用双指针来判断一个字符串是否为另一个字符串的子序列。

可以使用排序 复杂度会降下来
*/

func findLongestWord(s string, dictionary []string) string {
	var longestWord string
	for _, target := range dictionary {
		if len(longestWord) > len(target) || (len(longestWord) == len(target) && stringCompare(longestWord, target) < 0) {
			continue
		}
		//fmt.Println(longestWord, target)
		if isSubStr(s, target) {
			longestWord = target
		}
	}
	return longestWord

}

func isSubStr(string1 string, string2 string) bool {
	i, j := 0, 0
	for i < len(string1) && j < len(string2) {
		if string1[i] == string2[j] {
			j++
		}
		i++
	}
	return j == len(string2)
}

func stringCompare(string1 string, string2 string) int {
	// 相等长度字符串 返回字符串字母序较小值 0 的话就是两个字符串完全一样 -1 则string1 < string2 反之
	for i := 0; i < len(string1); i++ {
		if string1[i] < string2[i] {
			return -1
		} else if string1[i] > string2[i] {
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
