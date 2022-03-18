package main

/*
示例 1：
输入：[3,4,5,1,2]
输出：1

示例 2：
输入：[2,2,2,0,1]
输出：0

思路解析 双指针
*/

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
		if numbers[j] < numbers[j-1] {
			return numbers[j]
		}
		i++
		j--
	}
	return numbers[0]
}
