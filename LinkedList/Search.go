package main

func search(node *ListNode, t int) bool {
	curr := node
	for curr != nil {
		if curr.val == t {
			return true
		}
		curr = curr.next
	}
	return false
}
