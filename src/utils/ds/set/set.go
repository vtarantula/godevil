package set

import (
	"sync"
)

// Contains unique keys and permits set operations
type Set[T comparable] struct {
	mu        sync.RWMutex
	m_integer map[T]bool
}

func New[T comparable]() *Set[T] {
	set := &Set[T]{
		m_integer: make(map[T]bool),
	}
	return set
}

func (set *Set[T]) Size() int {
	return len(set.m_integer)
}

func (set *Set[T]) Empty() bool {
	return len(set.m_integer) == 0
}

func (set *Set[T]) Contains(val T) bool {
	if _, ok := set.m_integer[val]; ok {
		return true
	}
	return false
}

func (set *Set[T]) Add(val T) {
	set.mu.Lock()
	set.m_integer[val] = true
	set.mu.Unlock()
}

func (set *Set[T]) Delete(val T) {
	if set.Contains(val) {
		set.mu.Lock()
		delete(set.m_integer, val)
		set.mu.Unlock()
	}
}

func (set *Set[T]) Union(otherset *Set[T]) *Set[T] {
	unionset := New[T]()
	for v := range set.m_integer {
		unionset.Add(v)
	}
	for v := range otherset.m_integer {
		unionset.Add(v)
	}
	return unionset
}

func (set *Set[T]) Intersect(otherset *Set[T]) *Set[T] {
	intersectset := New[T]()
	for v := range set.m_integer {
		if otherset.Contains(v) {
			intersectset.Add(v)
		}
	}
	return intersectset
}
