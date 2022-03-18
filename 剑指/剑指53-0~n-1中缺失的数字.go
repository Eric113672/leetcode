package main

/*
输入: [0,1,3]
输出: 2

解法三:总和的差值
计算当数组不缺数字时的总和 count
count减去该数组nums的总和
得到的差值即是缺失的数字
--执行用时：20 ms --内存消耗：6 MB

等差数列
func missingNumber(nums []int) int {
    n:=len(nums)
    count:=(n+n*n) >> 1
    for _,v:=range nums{
        count-=v
    }
    return count
}

*/

func missingNumber(nums []int) int {
	if len(nums) == 1 {
		if nums[0] == 0 {
			return 1
		}
		if nums[0] == 1 {
			return 0
		}
	}
	if nums[0] != 0 {
		return 0
	}
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
		if nums[j]-1 != nums[j-1] {
			return nums[j] - 1
		}
		i++
		j--
	}
	return 1
}
