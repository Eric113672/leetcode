package main

/*
示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
*/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	nodeList := make([]*ListNode, 10)
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	return nodeList[len(nodeList)-k]
}
