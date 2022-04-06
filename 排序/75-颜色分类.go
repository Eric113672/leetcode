/*
* @Author: lishuang
* @Date:   2022/4/6 17:35
 */

package main

/*
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题。

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

荷兰国旗问题
荷兰国旗包含三种颜色：红、白、蓝。

有三种颜色的球，算法的目标是将这三种球按颜色顺序正确地排列。它其实是三向切分快速排序的一种变种，
在三向切分快速排序中，每次切分都将数组分成三个区间：小于切分元素、等于切分元素、大于切分元素，而该算法是将数组分成三个区间：等于红色、等于白色、等于蓝色。

*/

func sortColors(nums []int) {
	zero, one, two := -1, 0, len(nums)
	for one < two {
		if nums[one] == 0 {
			zero++
			nums[zero], nums[one] = nums[one], nums[zero]
			one++
		} else if nums[one] == 2 {
			two--
			nums[two], nums[one] = nums[one], nums[two]
		} else {
			one++
		}

	}
}
