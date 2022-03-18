package main

/*
示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.


单指针删除
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	tempHead := head
	for head != nil {
		pre := head
		head = head.Next
		if head.Val == val {
			pre.Next = head.Next
			return tempHead
		}
	}
	return tempHead
}
*/

//以下是双指针版本  但我咋感觉俩的时间复杂度是一样的呢？？

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	pre, cur := head, head
	for cur != nil && cur.Val != val {
		pre, cur = cur, cur.Next
	}

	if cur != nil {
		pre.Next = cur.Next
	}

	return head
}
