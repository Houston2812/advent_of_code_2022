// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package set

import "sync"

type ByteSet struct {
	items map[byte]bool
	lock  sync.RWMutex
}

func (s *ByteSet) Add(t byte) *ByteSet {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.items == nil {
		s.items = make(map[byte]bool)
	}

	_, ok := s.items[t]

	if !ok {
		s.items[t] = true
	}
	return s
}

func (s *ByteSet) Clear() {

	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = make(map[byte]bool)
}

func (s *ByteSet) Delete(item byte) bool {

	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.items[item]

	if ok {
		delete(s.items, item)
	}
	return ok
}

func (s *ByteSet) Has(item byte) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.items[item]

	return ok
}

func (s *ByteSet) Bytes() []byte {

	s.lock.Lock()
	defer s.lock.Unlock()

	items := []byte{}

	for i := range s.items {
		items = append(items, i)
	}

	return items
}

func (s *ByteSet) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items)
}
