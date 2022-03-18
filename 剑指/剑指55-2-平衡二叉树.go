/*
* @Author: lishuang
* @Date:   2022/3/9 11:37
 */

package main

import "math"

/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。
*/

//后序遍历 自底向上

func isBalanced(node *TreeNode) bool {
	return find(node) != -1
}

func find(node *TreeNode) float64 {
	if node == nil {
		return 0
	}

	l := find(node.Left) //!左节点
	if l == -1 {         //剪枝，不平衡时直接返回，
		return -1
	}

	r := find(node.Right) //!右节点
	if r == -1 {          //剪枝，不平衡时直接返回
		return -1
	}

	if math.Abs(l-r) > 1 { //剪枝，不平衡时直接返回
		return -1
	}

	return math.Max(l, r) + 1 //计算深度 !根节点
}

//自顶向下

func isBalanced2(node *TreeNode) bool {
	return node == nil || isBalanced(node.Left) &&
		math.Abs(height(node.Left)-height(node.Right)) < 2 && //两边最大深度差  大于2
		isBalanced(node.Right)
}

//计算节点最大深度
func height(node *TreeNode) float64 {
	if node == nil {
		return 0
	}
	return math.Max(height(node.Left), height(node.Right)) + 1
}

// 计算深度->对比深度
