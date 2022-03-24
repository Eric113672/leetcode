/*
* @Author: lishuang
* @Date:   2022/3/24 11:52
 */

package main

import "sort"

/*
给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。

示例1:

输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
*/

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	n := len(words)
	nums := make([]int, n)
	cnts := make([]int, n)
	// nums[i] 每个单词生成字母唯一位置对应的数字
	for i, word := range words {
		var num int
		for _, c := range word {
			num |= 1 << int(c-'a')
		}
		nums[i] = num
		cnts[i] = len(word)
	}
	var res int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if cnts[i]*cnts[j] < res {
				break
			}
			// 如果存在相同的字符 起与值肯定不为0 不存在的话则必为0
			if nums[i]&nums[j] == 0 {
				res = cnts[i] * cnts[j]
			}
		}
	}
	return res
}
