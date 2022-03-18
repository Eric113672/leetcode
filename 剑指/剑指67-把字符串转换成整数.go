/*
* @Author: lishuang
* @Date:   2022/3/17 11:40
 */

package main

import "math"

func strToInt(str string) int {
	//前导空格
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	str = str[i:]

	ans := 0      //64位数来存储好比较
	flag := false //符号位

	for i, v := range str {
		if v >= '0' && v <= '9' {
			ans = ans*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			flag = true
		} else if v == '+' && i == 0 {
		} else {
			break
		}

		if ans > math.MaxInt32 {
			if flag {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if flag {
		return -1 * ans
	}
	return ans
}
