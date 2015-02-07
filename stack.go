package stack

import (
	"sync"
)

type Stack struct {
	list []interface{}
	sync.Mutex
}

func NewStack() *Stack {
	return &Stack{list: make([]interface{}, 0)}
}

// Push, Pop, Peek, Size

func (stack *Stack) Push(obj interface{}) {
	stack.Lock()
	defer stack.Unlock()
	tmp := []interface{}{obj}
	stack.list = append(tmp, stack.list...)
}

func (stack *Stack) Pop() interface{} {
	stack.Lock()
	defer stack.Unlock()
	if len(stack.list) <= 0 {
		return nil
	}
	obj := stack.list[0]
	stack.list = append([]interface{}{}, stack.list[1:]...)
	return obj
}

func (stack *Stack) Peek() interface{} {
	stack.Lock()
	defer stack.Unlock()
	if len(stack.list) <= 0 {
		return nil
	}
	return stack.list[0]
}

func (stack *Stack) Size() int {
	return len(stack.list)
}
