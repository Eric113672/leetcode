/*
* @Author: lishuang
* @Date:   2022/4/20 18:04
 */

package main

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串。 返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
*/

func partition(s string) [][]string {
	var tmpString []string //切割字符串集合
	var res [][]string     //结果集合
	backTracking(s, tmpString, 0, &res)
	return res
}
func backTracking(s string, tmpString []string, startIndex int, res *[][]string) {
	if startIndex == len(s) { //到达字符串末尾了
		//进行一次切片拷贝，怕之后的操作影响tmpString切片内的值
		t := make([]string, len(tmpString))
		copy(t, tmpString)
		*res = append(*res, t)
	}
	for i := startIndex; i < len(s); i++ {
		//处理（首先通过startIndex和i判断切割的区间，进而判断该区间的字符串是否为回文，若为回文，则加入到tmpString，否则继续后移，找到回文区间）（这里为一层处理）
		if isPartition(s, startIndex, i) {
			tmpString = append(tmpString, s[startIndex:i+1])
		} else {
			continue
		}
		//递归
		backTracking(s, tmpString, i+1, res)
		//回溯
		tmpString = tmpString[:len(tmpString)-1]
	}
}

//判断是否为回文
func isPartition(s string, startIndex, end int) bool {
	left := startIndex
	right := end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		//移动左右指针
		left++
		right--
	}
	return true
}
