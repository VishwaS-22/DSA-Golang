package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	//Creating three nodes
	first := &ListNode{val: 1}
	second := &ListNode{val: 2}
	third := &ListNode{val: 3}

	//Connecting all
	first.next = second
	second.next = third

	//Traverse
	curr := first
	fmt.Println("Printing the LinkedList")
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}

}
