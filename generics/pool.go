package generics

import "sync"

// Pool is a generic wrapper around sync.Pool
type Pool[T any] struct {
	p sync.Pool
}

// PoolInstance creates a new generic pool with a constructor function.
func PoolInstance[T any](newFn func() T) *Pool[T] {
	return &Pool[T]{
		p: sync.Pool{
			New: func() any { return newFn() },
		},
	}
}

// Get retrieves an item from the pool.
func (gp *Pool[T]) Get() T {
	return gp.p.Get().(T)
}

// Put returns an item to the pool.
func (gp *Pool[T]) Put(v T) {
	gp.p.Put(v)
}
