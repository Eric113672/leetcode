/*
* @Author: lishuang
* @Date:   2022/4/18 17:54
 */

package main

import "fmt"

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

回溯算法 学习链接：
https://leetcode-cn.com/problems/permutations/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-hui-s-mfrp/

append(nums[:i],append([]int{cur},nums[i:]...)...)

[[5,4,2,6],[5,4,2,6],[5,6,2,4],[5,6,2,4],[5,2,6,4],[5,2,6,4],[4,5,2,6],[4,5,2,6],
[4,6,2,5],[4,6,2,5],[4,2,6,5],[4,2,6,5],[6,5,2,4],[6,5,2,4],[6,4,2,5],[6,4,2,5],
[6,2,4,5],[6,2,4,5],[2,5,6,4],[2,5,6,4],[2,4,6,5],[2,4,6,5],[2,6,4,5],[2,6,4,5]]

var res [][]int
func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums,len(nums),[]int{})
	return res
}
func backTrack(nums []int,numsLen int,path []int)  {
	if len(nums)==0{
		p:=make([]int,len(path))
		copy(p,path)
		res = append(res,p)
	}
	for i:=0;i<numsLen;i++{
		cur:=nums[i]
		path = append(path,cur)
		nums = append(nums[:i],nums[i+1:]...)//直接使用切片
		backTrack(nums,len(nums),path)
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)//回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}

*/
func permute(nums []int) [][]int {
	var res [][]int
	backTrack(nums, len(nums), []int{}, &res)
	return res
}

func backTrack(nums []int, numsLength int, path []int, res *[][]int) {
	if numsLength == 0 {
		// 这里的path 须得重新拷贝 因为path的值再变
		p := make([]int, len(path))
		copy(p, path)
		*res = append(*res, p)
		return
	}
	for i := 0; i < numsLength; i++ {
		curr := nums[i]
		path = append(path, curr)
		nums = append(nums[:i], nums[i+1:]...)
		backTrack(nums, len(nums), path, res)
		nums = append(nums[:i], append([]int{curr}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}

func makeTest(nums []int) {
	nums[0] = 2
	nums = append(nums, 0, 0, 0, 0, 0, 0)
}

func main() {
	aa := []int{0, 1}
	bb := [][]int{aa, {1, 1}}
	cc := [2]int{}
	makeTest(aa)
	fmt.Println(aa, bb, cc)
}
