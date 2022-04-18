/*
* @Author: lishuang
* @Date:   2022/4/17 10:55
 */

package main

import "fmt"

func minimumRounds(tasks []int) int {
	recordMap := make(map[int]int)
	count := 0
	for i := 0; i < len(tasks); i++ {
		if _, ok := recordMap[tasks[i]]; !ok {
			count++
		}
		recordMap[tasks[i]]++
	}
	newCount := make([]int, count)
	j := 0
	for _, v := range recordMap {
		if v < 2 {
			return -1
		}
		newCount[j] = v
		j++
	}
	//fmt.Println(count, newCount, recordMap)
	res := 0
	for i := 0; i < len(newCount); {
		if newCount[i] < 3 {
			if newCount[i] <= 0 {
				i++
				continue
			}
			res++
			i++
		} else {
			newCount[i] -= 3
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
}
