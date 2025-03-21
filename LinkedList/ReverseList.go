package main

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := node

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
