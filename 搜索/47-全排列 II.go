/*
* @Author: lishuang
* @Date:   2022/4/19 11:41
 */

package main

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

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
	used := [21]int{}//跟前一题唯一的区别，同一层不使用重复的数。关于used的思想carl在递增子序列那一题中提到过
	for i:=0;i<numsLen;i++{
		if used[nums[i]+10]==1{
			continue
		}
		cur:=nums[i]
		path = append(path,cur)
		used[nums[i]+10]=1
		nums = append(nums[:i],nums[i+1:]...)
		backTrack(nums,len(nums),path)
		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
		path = path[:len(path)-1]

	}

}
*/

func permuteUnique(nums []int) [][]int {
	var res [][]int
	backTrack2(nums, len(nums), []int{}, &res)
	return res
}

func backTrack2(nums []int, numsLength int, path []int, res *[][]int) {
	if numsLength == 0 {
		// 这里的path 须得重新拷贝 因为path的值再变
		p := make([]int, len(path))
		copy(p, path)
		*res = append(*res, p)
		return
	}
	used := [21]int{} //跟前一题唯一的区别，同一层不使用重复的数。关于used的思想carl在递增子序列那一题中提到过  每一层的递归开始时都是新的数组 只在这一层不会再重复此值
	for i := 0; i < numsLength; i++ {
		if used[nums[i]+10] == 1 {
			continue
		}
		curr := nums[i]
		path = append(path, curr)
		used[nums[i]+10] = 1
		nums = append(nums[:i], nums[i+1:]...)
		backTrack2(nums, len(nums), path, res)
		nums = append(nums[:i], append([]int{curr}, nums[i:]...)...)
		path = path[:len(path)-1]
	}
}
