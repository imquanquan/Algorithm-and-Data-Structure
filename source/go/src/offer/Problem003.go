package offer

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Value string
	Next * LinkedList
}

func PrintLinks(links *LinkedList) (int, error) {
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
