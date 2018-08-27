package offer

type LinkedList struct {
	size int
	head *LinkedListNode
	tail *LinkedListNode
}

func DeleteNode(list *LinkedList, node *LinkedListNode)  {
	if list.size == 1 && list.head == node {
		list = nil
		return
	}
	if list.head == node {
		node = list.head
		list.head = node.Next
		list.size--
		if list.size == 1 {
			list.tail = nil
		}
		return
	}
}
