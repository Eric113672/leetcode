package main

/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

/*
思路：

	对于对称的二叉树的任意两个对称节点有：
	1. Left.Val = Right.Val
	2. Left.Left.Val = Right.Right.Val 即 Left 的 左子节点 和 Right 的 右子节点 对称；
	3. Left.Right.Val = Right.Left.Val 即 Left 的 右子节点 和 Right 的 左子节点 对称；


	递归处理：

	1. 若 Left == Right == nil 同时为空时，则无子节点，不需递归，直接返回 true
	2. 以下任意条件成立时，返回 false
		1. 若 Left.Val != Right.Val
		2. Left 和 Right 其中一个为nil，另外一个不为nil
	3. 以下两个条件同时成立时，返回 true
		1. Left 的 左子节点 和 Right 的 右子节点 对称（进行递归）
		2. Left 的 右子节点 和 Right 的 左子节点 对称（进行递归）

*/

func dfs(left *TreeNode, right *TreeNode) bool {
	// 同时为空时，则无子节点，不需递归，直接返回 true
	if left == nil && right == nil {
		return true
	}
	// 左、右其中一个为空时 或者 左右两个结点 不同时，返回false
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	// 递归 Left 的 左子节点 和 Right 的 右子节点 对称，同时 Left 的 右子节点 和 Right 的 左子节点 对称 ，返回true
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}
