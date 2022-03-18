package main

/*
例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]
*/

func levelOrder3(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodeList := []*TreeNode{root}
	var recordTime int
	for len(nodeList) > 0 {
		var tempList []int
		currLength := len(nodeList)
		for i := 0; i < currLength; i++ {
			node := nodeList[0]
			nodeList = nodeList[1:]

			if node != nil {
				tempList = append(tempList, node.Val)
			}
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
		}
		recordTime++
		if recordTime%2 == 0 {
			reverse(tempList)
		}
		res = append(res, tempList)
	}
	return res
}

func reverse(tempList []int) []int {
	i, j := 0, len(tempList)-1
	for i < j {
		tempList[i], tempList[j] = tempList[j], tempList[i]
		i++
		j--
	}
	return tempList
}
