/*
* @Author: lishuang
* @Date:   2022/4/19 17:03
 */

package main

import (
	"fmt"
)

/*
给你一个 无重复元素 的整数数组candidates 和一个目标整数target，
找出candidates中可以使数字和为目标数target 的 所有不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为target 的不同组合数少于 150 个。

示例1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backTrack4(0, candidates, target, &res, 0, []int{})
	return res
}

func backTrack4(startIndex int, candidates []int, target int, res *[][]int, total int, path []int) {
	if total > target {
		return
	}
	if target == total {
		p := make([]int, len(path))
		copy(p, path)
		*res = append(*res, p)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if total > target {
			return
		}
		curr := candidates[i]
		path = append(path, curr)
		total += curr
		backTrack4(i, candidates, target, res, total, path)
		total -= curr
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
