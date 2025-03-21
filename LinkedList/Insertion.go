package main

import "fmt"

func insertion(node *ListNode, newVal int) {
	if node == nil {
		fmt.Println("Can't insert after a nil node!")
		return
	}

	newNode := &ListNode{val: newVal}
	newNode.next = node.next
	node.next = newNode
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}
