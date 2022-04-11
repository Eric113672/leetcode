/*
* @Author: lishuang
* @Date:   2022/4/10 10:58
 */

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
给你一个下标从 0 开始的字符串 expression ，格式为 "<num1>+<num2>" ，其中 <num1> 和 <num2> 表示正整数。

请你向 expression 中添加一对括号，使得在添加之后， expression 仍然是一个有效的数学表达式，并且计算后可以得到 最小 可能值。左括号 必须 添加在 '+' 的左侧，而右括号必须添加在 '+' 的右侧。

返回添加一对括号后形成的表达式 expression ，且满足 expression 计算得到 最小 可能值。如果存在多个答案都能产生相同结果，返回任意一个答案。

生成的输入满足：expression 的原始值和添加满足要求的任一对括号之后 expression 的值，都符合 32-bit 带符号整数范围。

输入：expression = "247+38"
输出："2(47+38)"
解释：表达式计算得到 2 * (47 + 38) = 2 * 85 = 170 。
注意 "2(4)7+38" 不是有效的结果，因为右括号必须添加在 '+' 的右侧。
可以证明 170 是最小可能值。

*/

func minimizeResult(expression string) (ans string) {
	// 差不多跟py一个思路 只是将a b c d 分开拆除了 来计算乘积  最开始 我还理解错题目了 以为是返回计算的最小值
	sp := strings.Split(expression, "+")
	l, r := sp[0], sp[1]
	for i, min := 0, math.MaxInt64; i < len(l); i++ { // 枚举左括号插入位置
		a := 1
		if i > 0 {
			a, _ = strconv.Atoi(l[:i]) // 左括号左边的数（不存在时为 1）
		}
		b, _ := strconv.Atoi(l[i:])    // 左括号右边的数
		for j := 1; j <= len(r); j++ { // 枚举右括号插入位置
			c, _ := strconv.Atoi(r[:j]) // 右括号左边的数
			d := 1
			if j < len(r) {
				d, _ = strconv.Atoi(r[j:]) // 右括号右边的数（不存在时为 1）
			}
			res := a * (b + c) * d
			if res < min {
				min = res
				ans = fmt.Sprintf("%s(%s+%s)%s", l[:i], l[i:], r[:j], r[j:])
			}
		}
	}
	return
}

/*
py的解法。
def minimizeResult(expression: str) -> str:
    expression = expression.split('+')
    min_num = 1 << 32
    # print(expression)
    for i in range(len(expression[0])):
        if i == 0:
            tempNum = "(" + expression[0]
            for j in range(len(expression[1])):
                #print("j", j)
                if j == len(expression[1]) - 1:
                    tempNum2 = expression[1] + ")"
                else:
                    tempNum2 = expression[1][:j+1] + ")*" + expression[1][j+1:]
                totalNum = tempNum + "+" + tempNum2
                #print(totalNum)
                if eval(totalNum) < min_num:
                    min_res = totalNum
                    min_num = eval(totalNum)

        else:
            tempNum = expression[0][:i] + "*(" + expression[0][i:]
            for j in range(len(expression[1])):
                #print("j", j)
                if j == len(expression[1]) - 1:
                    tempNum2 = expression[1] + ")"
                else:
                    tempNum2 = expression[1][:j+1] + ")*" + expression[1][j+1:]
                totalNum = tempNum + "+" + tempNum2
                #print(totalNum)
                if eval(totalNum) < min_num:
                    min_res = totalNum
                    min_num = eval(totalNum)
    min_res = min_res.split("*")

    return ''.join(min_res)
*/
