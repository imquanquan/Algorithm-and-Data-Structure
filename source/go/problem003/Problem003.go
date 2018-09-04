package problem003

import (
	"errors"
	"fmt"
)

type LinkedListNode struct {
	Value string
	Next * LinkedListNode
}

type LinkedList struct {
	Size int
	Head *LinkedListNode
	Tail *LinkedListNode
}

func PrintLinks(links *LinkedListNode) (int, error) {
	var stack [] string
	if links == nil {
		return 1, errors.New("links nil")
	}
	for links != nil {
		stack = append(stack, links.Value)
		links = links.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
	return 0, nil
}
