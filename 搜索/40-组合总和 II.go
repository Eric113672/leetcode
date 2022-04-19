/*
* @Author: lishuang
* @Date:   2022/4/19 17:59
 */

package main

import (
	"fmt"
	"sort"
)

/*
给定一个候选人编号的集合candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。

candidates中的每个数字在每个组合中只能使用一次。

注意：解集不能包含重复的组合。

示例1:

输入: candidates =[10,1,2,7,6,1,5], target =8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
*/

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	backTrack5(0, candidates, target, &res, 0, []int{})
	return res
}

func backTrack5(startIndex int, candidates []int, target int, res *[][]int, total int, path []int) {
	if total > target {
		return
	}
	if target == total {
		p := make([]int, len(path))
		copy(p, path)
		*res = append(*res, p)
		return
	}
	used := [51]int{}
	for i := startIndex; i < len(candidates); i++ {
		if total > target {
			return
		}
		if used[candidates[i]] == 1 {
			continue
		}
		curr := candidates[i]
		path = append(path, curr)
		total += curr
		used[candidates[i]] = 1
		backTrack5(i+1, candidates, target, res, total, path)
		total -= curr
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
