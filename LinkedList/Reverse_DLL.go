package main 

func reverseList_dll(head *DListNode) *DListNode {
	var temp *DListNode
	curr := head

	for curr != nil {
		temp = curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}
	if temp != nil {
        return temp.prev
    }
    return head
}