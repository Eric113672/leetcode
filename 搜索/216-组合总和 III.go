/*
* @Author: lishuang
* @Date:   2022/4/20 15:24
 */

package main

import "fmt"

/*
找出所有相加之和为n 的k个数的组合，且满足下列条件：

只使用数字1到9
每个数字最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了
*/

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	backTrack6(k, n, 0, 1, []int{}, &res)
	return res

}

func backTrack6(k, n, total, startIndex int, path []int, res *[][]int) {
	if total > n {
		return
	}
	if total == n && len(path) == k {
		p := make([]int, k)
		copy(p, path)
		*res = append(*res, p)
		return
	}
	//fmt.Println("curr startIndex", startIndex, path, total)
	for i := startIndex; i <= 9; i++ {
		path = append(path, i)
		total += i
		backTrack6(k, n, total, i+1, path, res)
		total -= i
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(9, 45))
}
