package main

func reverseList(head *ListNode) *ListNode {
	var tmp *ListNode
	for head != nil {
		next := head.Next
		head.Next = tmp
		tmp = head
		head = next
	}
	return head
}

func main() {
}
