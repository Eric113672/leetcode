/*
* @Author: lishuang
* @Date:   2022/3/4 10:32
 */

package main

import "fmt"

/*
输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
*/

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1

	for i < j {
		if nums[i]&1 == 1 {
			i++
		}
		if nums[j]&1 == 0 {
			j--
		}
		if nums[i]&1 == 0 && nums[j]&1 == 1 && i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	return nums
}

func main() {
	a := []int{11, 9, 3, 7, 16, 4, 2, 0}
	fmt.Println(exchange(a))
}
