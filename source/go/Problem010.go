package offer

import "errors"

type LinkedListNode struct {
	Value string
	Next * LinkedListNode
}

type LinkedList struct {
	Size int
	Head *LinkedListNode
	Tail *LinkedListNode
}

func DeleteNode(list *LinkedList, index int) error {
	if index < 0 || index > list.Size - 1 {
		return errors.New("please check index")
	} else {
		item := list.Head
		for i := 0; i < index - 1; i++ {
			item = item.Next
		}
		item.Next = item.Next.Next
		return nil
	}
}

