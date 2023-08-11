package utils

import (
	"sync"
)

type Set[T comparable] struct {
	m map[T]struct{}
	sync.RWMutex
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

//func main() {
//	// Initialize our Set
//	s := New()
//
//	// Add example items
//	s.Add("item1")
//	s.Add("item1") // duplicate item
//	s.Add("item2")
//	fmt.Printf("%d items\n", s.Len())
//
//	// Clear all items
//	s.Clear()
//	if s.IsEmpty() {
//		fmt.Printf("0 items\n")
//	}
//
//	s.Add("item2")
//	s.Add("item3")
//	s.Add("item4")
//
//	// Check for existence
//	if s.Has("item2") {
//		fmt.Println("item2 does exist")
//	}
//
//	// Remove some of our items
//	s.Remove("item2")
//	s.Remove("item4")
//	fmt.Println("list of all items:", s.List())
//}

// Add add
func (s *Set[T]) Add(item T) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
}

func (s *Set[T]) AddList(items []T) {
	for _, item := range items {
		s.Add(item)
	}
}

// Remove deletes the specified item from the map
func (s *Set[T]) Remove(item T) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

// Has looks for the existence of an item
func (s *Set[T]) Has(item T) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Len returns the number of items in a set.
func (s *Set[T]) Len() int {
	return len(s.List())
}

// Clear removes all items from the set
func (s *Set[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[T]struct{})
}

// IsEmpty checks for emptiness
func (s *Set[T]) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Set returns a slice of all items
func (s *Set[T]) List() []T {
	s.RLock()
	defer s.RUnlock()
	list := make([]T, 0)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
