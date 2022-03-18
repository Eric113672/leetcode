/*
* @Author: lishuang
* @Date:   2022/3/7 14:45
 */

package main

/*
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
先序遍历 建立记录数组
*/
func kthLargest(root *TreeNode, k int) int {
	numsList := make([]int, k)
	var firstOrder func(*TreeNode)
	firstOrder = func(node *TreeNode) {
		if node.Left != nil {
			firstOrder(node.Left)
		}
		numsList = append(numsList, node.Val)
		if node.Right != nil {
			firstOrder(node.Right)
		}
	}
	firstOrder(root)
	return numsList[len(numsList)-k]
}
