package stack

import (
	"sync"
)

type ItemType interface{}

type Stack struct {
	items  []ItemType
	rwLock sync.RWMutex
}

func (stack *Stack) Push(t ItemType) {
	if stack.items == nil {
		stack.items = []ItemType{}
	}

	stack.rwLock.Lock()

	stack.items = append(stack.items, t)

	stack.rwLock.Unlock()
}

func (stack *Stack) Pop() *ItemType {
	if len(stack.items) == 0 {
		return nil
	}

	stack.rwLock.Lock()

	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[0 : len(stack.items)-1]

	stack.rwLock.Unlock()

	return &item
}

func (stack *Stack) Size() int {
	stack.rwLock.RLock()

	defer stack.rwLock.RUnlock()

	return len(stack.items)
}

func (stack *Stack) All() []ItemType {

	stack.rwLock.RLock()

	defer stack.rwLock.RUnlock()

	return stack.items
}

func (stack *Stack) IsEmpty() bool {
	stack.rwLock.RLock()

	defer stack.rwLock.RUnlock()

	return len(stack.items) == 0
}

