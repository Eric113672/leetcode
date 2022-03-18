/*
* @Author: lishuang
* @Date:   2022/3/7 16:33
 */

package main

/*
从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，
而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
输入: [0,0,1,2,5]
输出: True

*/

func isStraight(nums []int) bool {

	var quickSort func(strArr []int, l, r int)
	quickSort = func(strArr []int, l, r int) {
		if l >= r {
			return
		}
		i, j, pivot := l, r, strArr[l]
		for i < j {
			for i < j && pivot <= strArr[j] {
				j--
			}
			for i < j && pivot > strArr[i] {
				i++
			}
			strArr[i], strArr[j] = strArr[j], strArr[i]
		}
		strArr[i], strArr[l] = strArr[l], strArr[i]
		quickSort(strArr, l, i-1)
		quickSort(strArr, i+1, r)
	}
	quickSort(nums, 0, len(nums)-1)
	zeroCount := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		if i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			continue
		}
		zeroCount = zeroCount - (nums[i+1] - nums[i] - 1)
		if zeroCount < 0 {
			return false
		}
		return true
	}
	return true
}
