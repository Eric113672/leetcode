/*
* @Author: lishuang
* @Date:   2022/4/12 15:47
 */

package main

/*
给你一个由数字和运算符组成的字符串expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。

示例 1：

输入：expression = "2-1-1"
输出：[0,2]
解释：
((2-1)-1) = 0
(2-(1-1)) = 2
*/

// 大佬思路  分支算法初探
/*
对于一个形如 x op y（op 为运算符，x 和 y 为数） 的算式而言，它的结果组合取决于 x 和 y 的结果组合数，而 x 和 y 又可以写成形如 x op y 的算式。

因此，该问题的子问题就是 x op y 中的 x 和 y：以运算符分隔的左右两侧算式解。

然后我们来进行 分治算法三步走：

分解：按运算符分成左右两部分，分别求解
解决：实现一个递归函数，输入算式，返回算式解
合并：根据运算符合并左右两部分的解，得出最终解

*/

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	// 如果是数字，直接返回
	if isDigit(input) {
		tmp, _ := strconv.Atoi(input)
		return []int{tmp}
	}
	// 空切片
	var res []int
	// 遍历字符串
	for index, c := range input {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的算式
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}

	return res
}

// 判断是否为全数字
func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
