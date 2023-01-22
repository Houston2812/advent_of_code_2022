package set

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=set.go -out=set-string.go gen "Item=string Value=int"
//go:generate genny -in=set.go -out=set-int.go gen "Item=int"
//go:generate genny -in=set.go -out=set-byte.go gen "Item=byte"
//go:generate genny -in=set.go -out=set-rune.go gen "Item=rune"


type Item generic.Type

type ItemSet struct {
	items map[Item]bool
	lock  sync.RWMutex
}

func (s *ItemSet) Add(t Item) *ItemSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[Item]bool)
	}

	_, ok := s.items[t]

	if !ok {
		s.items[t] = true
	}
	return s
}

func (s *ItemSet) Clear() {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = make(map[Item]bool)
}

func (s *ItemSet) Delete(item Item) bool {

	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.items[item]

	if ok {
		delete(s.items, item)
	}
	return ok
}

func (s *ItemSet) Has(item Item) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.items[item]

	return ok
}

func (s *ItemSet) Items() []Item {

	s.lock.Lock()
	defer s.lock.Unlock()

	items := []Item{}

	for i := range s.items {
		items = append(items, i)
	}

	return items
}

func (s *ItemSet) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items)
}
