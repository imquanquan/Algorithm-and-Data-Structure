package offer

import "errors"

type MyQueue struct {
	Stack1 [] int
	Stack2 [] int
}

func Push(queue *MyQueue, val int) {
	queue.Stack1 = append(queue.Stack1, val)
}

func Pop(queue *MyQueue) (int, error) {
	if len(queue.Stack1) == 0 && len(queue.Stack2) == 0 {
		return -1, errors.New("queue is empty")
	}
	if len(queue.Stack2) != 0 {
		val := queue.Stack2[len(queue.Stack2) - 1]
		queue.Stack2 = queue.Stack2[:len(queue.Stack2) - 1]
		return val, nil
    }
	for i := len(queue.Stack1) - 1; i >= 0; i-- {
		queue.Stack2 = append(queue.Stack2, queue.Stack1[i])
		queue.Stack1 = queue.Stack1[:len(queue.Stack1) - 1]
	}
	val := queue.Stack2[len(queue.Stack2) - 1]
	queue.Stack2 = queue.Stack2[:len(queue.Stack2) - 1]
	return val, nil
}
