/*
* @Author: lishuang
* @Date:   2022/4/18 17:14
 */

package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
*/

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	buildPath(root, "", &res)
	return res
}

func buildPath(root *TreeNode, pathStr string, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		pathStr += strconv.Itoa(root.Val)
		*res = append(*res, pathStr)
		return
	}
	pathStr += strconv.Itoa(root.Val) + "->"
	buildPath(root.Left, pathStr, res)
	buildPath(root.Right, pathStr, res)
}
