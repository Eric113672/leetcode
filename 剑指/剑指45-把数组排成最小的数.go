/*
* @Author: lishuang
* @Date:   2022/3/7 14:57
 */

package main

import (
	"fmt"
	"strconv"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
快排调整 比较规则修改
输入: [10,2]
输出: "102"
*/
func minNumber(nums []int) string {
	// 快排实现排序，排序后转成string
	res := make([]string, len(nums))
	for i, v := range nums {
		res[i] = strconv.Itoa(v)
	}
	//fmt.Println("firstRes", res, reflect.TypeOf(res[0]))
	compare := func(str1, str2 string) bool {
		num1, _ := strconv.Atoi(str1 + str2)
		num2, _ := strconv.Atoi(str2 + str1)
		if num1 < num2 {
			return true
		}
		return false
	}
	var quickSort func(strArr []string, l, r int)
	quickSort = func(strArr []string, l, r int) {
		if l >= r {
			return
		}
		i, j, pivot := l, r, strArr[l]
		for i < j {
			for i < j && compare(pivot, strArr[j]) {
				j--
			}
			for i < j && !compare(pivot, strArr[i]) {
				i++
			}
			strArr[i], strArr[j] = strArr[j], strArr[i]
		}
		strArr[i], strArr[l] = strArr[l], strArr[i]
		quickSort(strArr, l, i-1)
		quickSort(strArr, i+1, r)
	}
	quickSort(res, 0, len(nums)-1)
	ans := ""
	fmt.Println("fin res", res)
	for _, s := range res {
		ans += s
	}
	return ans
}

func main() {
	nums := []int{10, 2}
	minNumber(nums)

}
