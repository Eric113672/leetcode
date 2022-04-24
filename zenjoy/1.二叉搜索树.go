/*
* @Author: lishuang
* @Date:   2022/4/22 20:32
 */

package main

/*
思路解析: 二叉搜索树 重要特性 左子树节点值小于根节点值小于右子树节点值 那么可以直接中序遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBST(root *TreeNode) bool {
	var pre *TreeNode // 上一个节点的指针
	var LDR func(node *TreeNode) bool
	LDR = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		leftResult := LDR(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		rightResult := LDR(node.Right)
		return leftResult && rightResult
	}
	return LDR(root)
}
