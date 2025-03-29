package main 

func delete_dll(head **DListNode, node *DListNode){
	if node == nil {
		return
	}

	if *head == node {
		*head = node.next
		if *head != nil {
			(*head).prev = nil
		}
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev 
	}
}