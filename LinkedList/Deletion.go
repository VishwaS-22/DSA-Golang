package main 

import "fmt"

func deletion(node *ListNode){
	if node == nil || node.next == nil {
		fmt.Println("Nothing to Delete.")
		return
	}
	node.next = node.next.next
}