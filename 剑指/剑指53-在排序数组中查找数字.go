package main

/*
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && target == nums[0] {
		return 1
	}

	lo, hi := 0, len(nums)-1
	index := -1

	for lo <= hi {
		i := (lo + hi) >> 1
		if nums[i] == target {
			index = i
			break
		} else if nums[i] > target {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}

	if index == -1 {
		return 0
	}

	cnt := 0
	for i := index; i >= 0 && nums[i] == target; i-- {
		cnt++
	}
	for i := index; i < len(nums) && nums[i] == target; i++ {
		cnt++
	}
	return cnt - 1
}

func main() {

}
