package main

import "fmt"

type ListNode struct {
	val  int
	prev *ListNode
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

	// //Traverse
	// curr := first
	// fmt.Println("Printing the LinkedList")
	// for curr != nil {
	// 	fmt.Println(curr.val)
	// 	curr = curr.next
	// }
	// fmt.Println("Before Insertion:")
	// printList(first)

	// insertion(first, 99)

	// fmt.Println("After Insertion:")
	// printList(first)

	// fmt.Println("Before Deletion:")
	// printList(first)

	// deletion(first)

	// fmt.Println("After Deletion:")
	// printList(first)

	// t := 2
	// found := search(first, t)

	// if found {
	// 	fmt.Printf("The val %d is found in the List.", t)
	// } else {
	// 	fmt.Printf("The val %d is not found in the List.", t)
	// }

	// fmt.Println("Before Reversing:")
	// printList(first)

	// newHead := reverseList(first)

	// fmt.Println("After Reversing:")
	// printList(newHead)

	// Cycle Detection
	// third.next = first

	// if cycle(first) {
	// 	fmt.Println("Cycle Detected.")
	// }else{
	// 	fmt.Println("Not a Cycle.")
	// }

	//Doubly LinkedList

	second.prev = first
	third.prev = second
	fmt.Println("Forward Traversal:")
	printList(first)
	fmt.Println("Backward Traversal:")

	curr := third
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.prev
	}
	fmt.Println("nil")
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}
