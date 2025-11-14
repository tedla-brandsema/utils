package generics 

import (
	"maps"
	"sync"
)

// Registry is a thread-safe in-memory registry.
type Registry[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// New creates a new Registry.
func New[K comparable, V any]() *Registry[K, V] {
	return &Registry[K, V]{
		data: make(map[K]V),
	}
}

// Set registers or updates an item in the registry.
func (r *Registry[K, V]) Set(key K, value V) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[key] = value
}

// Get retrieves an item from the registry.
// It returns the value and a boolean indicating if it exists.
func (r *Registry[K, V]) Get(key K) (V, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.data[key]
	return val, ok
}

// Remove deletes an item from the registry.
func (r *Registry[K, V]) Remove(key K) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, key)
}

// All returns a copy of all items in the registry.
func (r *Registry[K, V]) All() map[K]V {
	r.mu.RLock()
	defer r.mu.RUnlock()
	copy := make(map[K]V, len(r.data))
	maps.Copy(copy, r.data)
	return copy
}
