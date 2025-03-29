package main

import "fmt"

func insert_dll(node *DListNode, newval int) {
	if node == nil {
		fmt.Println("Cannot insert after a nil node!")
		return
	}

	newNode := &DListNode{val: newval}

	newNode.next = node.next
	newNode.prev = node

	if node.next != nil {
		node.next.prev = newNode
	}
	node.next = newNode
}
