package main

func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{val: arr[0]}
	curr := head

	for _, v := range arr[1:] {
		curr.next = &ListNode{val: v}
		curr = curr.next
	}

	return head
}
