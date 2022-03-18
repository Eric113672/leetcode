/*
* @Author: lishuang
* @Date:   2022/3/7 14:20
 */

package main

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]

两种思路 DFS 与 BFS
都实现一下

*/

func pathSum(root *TreeNode, target int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, path []int) {
		if node.Left == nil && node.Right == nil && sumOfArray(path) == target {
			res = append(res, path)
		}
		temp := make([]int, len(path))
		copy(temp, path)
		if node.Left != nil {
			leftPath := append(temp, node.Left.Val)
			dfs(node.Left, leftPath)
		}
		if node.Right != nil {
			rightPath := append(temp, node.Right.Val)
			dfs(node.Right, rightPath)
		}
	}
	dfs(root, []int{root.Val})
	return res
}

func sumOfArray(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}

	return sum
}

type Element struct {
	Node *TreeNode
	Path []int
}

// BFS
func pathSum2(root *TreeNode, target int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []Element{Element{root, []int{root.Val}}}
	for len(queue) != 0 {
		element := queue[0]
		queue = queue[1:]
		node, path := element.Node, element.Path
		if node.Left == nil && node.Right == nil && sumOfArray(path) == target {
			res = append(res, path)
		}
		temp := make([]int, len(path))
		copy(temp, path)
		if node.Left != nil {
			leftPath := append(temp, node.Left.Val)
			queue = append(queue, Element{node.Left, leftPath})
		}
		if node.Right != nil {
			rightPath := append(temp, node.Right.Val)
			queue = append(queue, Element{node.Right, rightPath})
		}
	}
	return res
}
