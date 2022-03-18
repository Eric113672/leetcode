package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	numList := []int{}
	for head != nil {
		numList = append(numList, head.Val)
		head = head.Next
	}
	for i, j := 0, len(numList)-1; i < j; {
		numList[i], numList[j] = numList[j], numList[i]
		i++
		j--
	}
	return numList
}

func main() {

}
