package main

/*
示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
这题可以算是解得完美
虚拟的头节点  同leetcode21

为啥比21那个更快了
*/

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var n *ListNode
	if l1.Val < l2.Val {
		n = l1
		n.Next = mergeTwoLists(l1.Next, l2)
	} else {
		n = l2
		n.Next = mergeTwoLists(l1, l2.Next)
	}

	return n
}
