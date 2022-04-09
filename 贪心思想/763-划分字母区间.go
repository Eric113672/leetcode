/*
* @Author: lishuang
* @Date:   2022/4/9 16:49
 */

package main

/*
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

示例：

输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

记录每个字符出现的lastIndex,从0开始 看是否有大于当前的lastIndex 有则更新 无则append append之后更新firstIndex就行
*/

func partitionLabels(s string) []int {
	lastIndexsOfChar := [26]int{}
	for i := 0; i < len(s); i++ {
		lastIndexsOfChar[char2Index(s[i])] = i
	}
	partitions := make([]int, 0)
	firstIndex, lastIndex := 0, 0
	for firstIndex < len(s) {
		lastIndex = firstIndex
		for i := firstIndex; i < len(s) && i <= lastIndex; i++ {
			index := lastIndexsOfChar[char2Index(s[i])]
			if index > lastIndex {
				lastIndex = index
			}
		}
		partitions = append(partitions, lastIndex-firstIndex+1)
		firstIndex = lastIndex + 1
	}
	return partitions
}

func char2Index(a uint8) int {
	return int(a - 'a')
}
