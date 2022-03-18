package main

/*
现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target=5，返回true。

给定target=20，返回false。

思路解析 先判断每一行的区间 看当前值是否在这个个区间 在的话再用二分查找 不在就直接下一行

*/

func findNumberIn2DArray(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		currLength := len(matrix[i])
		if currLength == 0 {
			continue
		}
		if currLength == 1 && matrix[i][0] != target {
			continue
		}
		if currLength == 1 && matrix[i][0] == target {
			return true
		}

		if matrix[i][0] <= target && target <= matrix[i][currLength-1] {
			low, high := 0, len(matrix[i])-1
			for low <= high {
				mid := (low + high) >> 1
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] < target {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}

		}
	}
	return false
}
