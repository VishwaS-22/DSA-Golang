package main

import "fmt"

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}

func printDList(head *DListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}
