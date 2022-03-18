/*
* @Author: lishuang
* @Date:   2022/3/7 11:24
 */

package main

import "fmt"

/*
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

输入：m = 2, n = 3, k = 1
输出：3

*/

func movingCount(m int, n int, k int) int {
	if k == 0 {
		return 1
	}
	total := 0
	var exist = make([][]bool, 0)
	for i := 0; i < m; i++ {
		var line = make([]bool, n)
		exist = append(exist, line)
	}
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		// 终止条件，坐标范围错误或者 i j 坐标对应 数位之后大于行坐标和列坐标
		if i < 0 || j < 0 || i >= m || j >= n || exist[i][j] || !isAvail(i, j, k) {
			return
		}
		// 如果往回前找，不会找到相同字符，如 ABAB，k = 2 时，往前往后都是 B ；
		// 将 B 修改（剪枝）为不存在的字符，杜绝往回找成功的可能性。
		total++
		fmt.Println(i, j, total)
		exist[i][j] = true
		dfs(i, j-1, k)
		dfs(i, j+1, k)
		dfs(i+1, j, k)
		dfs(i-1, j, k)

	}
	dfs(0, 0, k)
	return total
}

func isAvail(x int, y int, k int) bool {
	var sum = 0
	for x > 0 {
		sum += x % 10
		x = x / 10
	}
	for y > 0 {
		sum += y % 10
		y = y / 10
	}
	fmt.Println("isAvail", x, y, sum, k)
	if sum <= k {
		return true
	}
	return false
}

func main() {
	fmt.Println(movingCount(7, 2, 3))

}
