package queue

import (
	"sync"
)

type ItemType interface{}

type Queue struct {
	items  []ItemType
	rwLock sync.RWMutex
}

func (queue *Queue) PushTop(t ItemType) {
	if queue.items == nil {
		queue.items = []ItemType{}
	}

	queue.rwLock.Lock()

	queue.items = append(queue.items, t)

	queue.rwLock.Unlock()
}

func (queue *Queue) PushBottom(t ItemType) {
	if queue.items == nil {
		queue.items = []ItemType{}
	}

	queue.rwLock.Lock()

	queue.items = append(queue.items, t)

	var tmp_item []ItemType
	tmp_item = append(tmp_item, t)

	queue.items = append(tmp_item, queue.items...)

	queue.rwLock.Unlock()
}

func (queue *Queue) PopTop() *ItemType {
	if len(queue.items) == 0 {
		return nil
	}

	queue.rwLock.Lock()

	item := queue.items[len(queue.items)-1]
	queue.items = queue.items[0 : len(queue.items)-1]

	queue.rwLock.Unlock()

	return &item
}

func (queue *Queue) PopBottom() *ItemType {
	if len(queue.items) == 0 {
		return nil
	}

	queue.rwLock.Lock()

	item := queue.items[0]
	queue.items = queue.items[1 : len(queue.items)]

	queue.rwLock.Unlock()

	return &item
}

func (queue *Queue) Size() int {
	queue.rwLock.RLock()

	defer queue.rwLock.RUnlock()

	return len(queue.items)
}

func (queue *Queue) All() []ItemType {

	queue.rwLock.RLock()

	defer queue.rwLock.RUnlock()

	return queue.items
}

func (queue *Queue) IsEmpty() bool {
	queue.rwLock.RLock()

	defer queue.rwLock.RUnlock()

	return len(queue.items) == 0
}

