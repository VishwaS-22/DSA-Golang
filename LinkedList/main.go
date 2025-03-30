package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

type DListNode struct {
	val  int
	prev *DListNode
	next *DListNode
}

func main() {
	//Creating three nodes
	// first := &ListNode{val: 1}
	// second := &ListNode{val: 2}
	// third := &ListNode{val: 3}

	// //Connecting all
	// first.next = second
	// second.next = third

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

	// second.prev = first
	// third.prev = second
	// fmt.Println("Forward Traversal:")
	// printList(first)
	// fmt.Println("Backward Traversal:")

	// curr := third
	// for curr != nil {
	// 	fmt.Printf("%d -> ", curr.val)
	// 	curr = curr.prev
	// }
	// fmt.Println("nil")

	//Doubly LinkedList
	//Creating three nodes
	// first := &DListNode{val: 1}
	// second := &DListNode{val: 2}
	// third := &DListNode{val: 3}

	// //Connecting all
	// first.next = second
	// second.prev = first
	// second.next = third
	// third.prev = second

	// fmt.Println("Original List:")
	// printDList(first)

	// insert_dll(first, 90)
	// fmt.Println("After inserting:")
	// printDList(first)

	// fmt.Println("Before deleting:")
	// printDList(first)

	// fmt.Println("After deleting:")
	// delete_dll(&first, first)
	// printDList(first)

	// newHead := reverseList_dll(first)

	// fmt.Println("Reversed list:")
	// printDList(newHead)
    
	list := createLinkedList([]int{1,2,3,4,5})
	mid := findMiddle(list)
	fmt.Println("Middle node is:")
	printList(mid)
}