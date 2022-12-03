package set

import "sync"

// Contains unique keys and permits set operations
type Set[T comparable] struct {
	mu     sync.RWMutex
	m_data map[T]bool
}

func New[T comparable]() *Set[T] {
	set := &Set[T]{
		m_data: make(map[T]bool),
	}
	return set
}

func (set *Set[T]) Size() int {
	return len(set.m_data)
}

func (set *Set[T]) Empty() bool {
	return len(set.m_data) == 0
}

func (set *Set[T]) Contains(val T) bool {
	if _, ok := set.m_data[val]; ok {
		return true
	}
	return false
}

func (set *Set[T]) Add(val T) {
	set.mu.Lock()
	set.m_data[val] = true
	set.mu.Unlock()
}

func (set *Set[T]) Delete(val T) {
	if set.Contains(val) {
		set.mu.Lock()
		delete(set.m_data, val)
		set.mu.Unlock()
	}
}

func (set *Set[T]) Union(otherset *Set[T]) []T {
	l_union := make([]T, 0)
	for v := range set.m_data {
		l_union = append(l_union, v)
	}
	for v := range otherset.m_data {
		if !set.Contains(v) {
			l_union = append(l_union, v)
		}
	}
	return l_union
}

func (set *Set[T]) Intersect(otherset *Set[T]) []T {
	l_intersect := make([]T, 0)
	for v := range set.m_data {
		if otherset.Contains(v) {
			l_intersect = append(l_intersect, v)
		}
	}
	return l_intersect
}
