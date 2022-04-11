/*
* @Author: lishuang
* @Date:   2022/4/10 10:39
 */

package main

import (
	"fmt"
	"strconv"
)

/*
给你一个正整数 num 。你可以交换 num 中 奇偶性 相同的任意两位数字（即，都是奇数或者偶数）。

返回交换 任意 次之后 num 的 最大 可能值。

输入：num = 1234
输出：3412
解释：交换数字 3 和数字 1 ，结果得到 3214 。
交换数字 2 和数字 4 ，结果得到 3412 。
注意，可能存在其他交换序列，但是可以证明 3412 是最大可能值。
注意，不能交换数字 4 和数字 1 ，因为它们奇偶性不同。
*/
func largestInteger(num int) int {
	strNums := strconv.Itoa(num)
	bytesNums := []byte(strNums)
	for i := 0; i < len(bytesNums); i++ {
		if bytesNums[i]&1 == 0 {
			for j := i + 1; j < len(bytesNums); j++ {
				if bytesNums[j]&1 == 0 {
					if bytesNums[j] > bytesNums[i] {
						bytesNums[i], bytesNums[j] = bytesNums[j], bytesNums[i]
					}
				}
			}
		} else {
			for j := i + 1; j < len(bytesNums); j++ {
				if bytesNums[j]&1 != 0 {
					if bytesNums[j] > bytesNums[i] {
						bytesNums[i], bytesNums[j] = bytesNums[j], bytesNums[i]
					}
				}
			}
		}
	}
	res, _ := strconv.Atoi(string(bytesNums))
	return res
}
func main() {
	fmt.Println(maximumProduct([]int{24, 5, 64, 53, 26, 38}, 54))
}
