/*
* @Author: lishuang
* @Date:   2022/4/19 15:11
 */

package main

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/
/*
func combine(n int, k int) [][]int {
	var res [][]int
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	backTrace3(0, nums, n, []int{}, &res, k)
	return res
}

func backTrace3(startIndex int, nums []int, numsLength int, path []int, res *[][]int, k int) {
	if len(path) == k {
		p := make([]int, k)
		copy(p, path)
		*res = append(*res, p)
		return
	}
	for i := startIndex; i < numsLength; i++ {
		curr := nums[i]
		path = append(path, curr)
		backTrace3(i+1, nums, len(nums), path, res, k)
		path = path[:len(path)-1]
	}

}

func main() {
	fmt.Println(combine(4, 2))
}
*/
// 剪枝版

/*
图中每一个节点（图中为矩形），就代表本层的一个for循环，那么每一层的for循环从第二个数开始遍历的话，都没有意义，都是无效遍历。

所以，可以剪枝的地方就在递归中每一层的for循环所选择的起始位置。

如果for循环选择的起始位置之后的元素个数 已经不足 我们需要的元素个数了，那么就没有必要搜索了。

注意代码中i，就是for循环里选择的起始位置。


for (int i = startIndex; i <= n; i++) {
接下来看一下优化过程如下：

已经选择的元素个数：path.size();

还需要的元素个数为: k - path.size();

在集合n中至多要从该起始位置 : n - (k - path.size()) + 1，开始遍历

为什么有个+1呢，因为包括起始位置，我们要是一个左闭的集合。

举个例子，n = 4，k = 3， 目前已经选取的元素为0（path.size为0），n - (k - 0) + 1 即 4 - ( 3 - 0) + 1 = 2。

从2开始搜索都是合理的，可以是组合[2, 3, 4]。

这里大家想不懂的话，建议也举一个例子，就知道是不是要+1了。

所以优化之后的for循环是：

*/

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}
func backtrack(n, k, start int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
	}
	if len(track)+n-start+1 < k {
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
