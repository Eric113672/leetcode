/*
* @Author: lishuang
* @Date:   2022/3/9 10:46
 */

package main

/*
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。

层次遍历 为啥执行不过 报找不到left的地址  我是真的搞不懂
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	treeNodeList, res := []*TreeNode{root}, 0
	for len(treeNodeList) > 0 {
		tmp := make([]*TreeNode, 2)
		for _, node := range treeNodeList {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		treeNodeList = tmp
		res++
	}
	return res
}
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	treeNodeList, res := []*TreeNode{root}, 0
	for len(treeNodeList) > 0 {
		lens := len(treeNodeList)
		for i := 0; i < lens; i++ {
			if treeNodeList[i].Left != nil {
				treeNodeList = append(treeNodeList, treeNodeList[i].Left)
			}
			if treeNodeList[i].Right != nil {
				treeNodeList = append(treeNodeList, treeNodeList[i].Right)
			}
		}
		treeNodeList = treeNodeList[lens:]
		res++
	}
	return res
}
