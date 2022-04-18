/*
* @Author: lishuang
* @Date:   2022/4/15 11:34
 */

package main

/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

// 排列组合的 回溯 题目 也是比较经典的题目
/*
func letterCombinations(digits string) (ans []string) {
	var m = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'h', 'i', 'g'},
		'5': {'k', 'l', 'j'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	ld := len(digits)
	if ld == 0 {
		return
	}

	var trace func(k int, tmp []byte)
	trace = func(k int, tmp []byte) {
		if k == ld {
			ans = append(ans, string(tmp))
			return
		}
		for i := 0; i < len(m[digits[k]]); i++ {
			tmp = append(tmp, m[digits[k]][i])
			trace(k+1, tmp)
			// 回溯
			tmp = tmp[:len(tmp)-1]
		}
	}

	trace(0, []byte{})

	return
}*/
func letterCombinations(digits string) []string {
	// 回溯法
	length := len(digits)
	// 题目规定最大是4位数字，考虑边界条件
	if length == 0 || length > 4 {
		return nil
	}
	letterMap := [10]string{ // index-string
		"",     //0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	var res []string

	var backtracking func(tmpstring, digits string, index int, letterMap [10]string, res *[]string)
	// 回溯三部曲
	// 1.回溯函数及参数
	backtracking = func(tmpstring, digits string, index int, letterMap [10]string, res *[]string) {
		// 2.回溯终止条件
		if len(tmpstring) == len(digits) {
			*res = append(*res, tmpstring)
			return
		}
		// 3.单层回溯逻辑
		tmpK := digits[index] - '0' // index指向的数字转为int [0,9]
		letters := letterMap[tmpK]  // 诸如"abc", "def"
		for i := 0; i < len(letters); i++ {
			tmpstring = tmpstring + string(letters[i])               // 处理节点
			backtracking(tmpstring, digits, index+1, letterMap, res) // 递归下一层
			tmpstring = tmpstring[:len(tmpstring)-1]                 // 撤销节点
		}
	}
	backtracking("", digits, 0, letterMap, &res)
	return res
}
