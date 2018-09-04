package problem012

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

func LastKTH(head *LinkedListNode, k int) (string, error) {
	if head == nil || k <= 0{
		return "", errors.New("error")
	}
	fast := head
	for fast != nil && k - 1 >= 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		head = head.Next
	}
	if k == 0 {
		return head.Value, nil
	}
	return "", errors.New("error")
}