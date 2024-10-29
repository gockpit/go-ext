package sync

import "sync"

func NewMap[M ~map[K]V, K comparable, V any](m M) *Map[M, K, V] {
	if m == nil {
		return nil
	}
	return &Map[M, K, V]{
		m: m,
	}
}

type Map[M ~map[K]V, K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (m *Map[M, K, V]) Delete(k K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, k)
}

func (m *Map[M, K, V]) Load(k K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m[k]
}

func (m *Map[M, K, V]) Store(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[k] = v
}

func (m *Map[M, K, V]) Swap(k K, v V) (previous V, loaded bool) {
	m.mu.RLock()
	previous, loaded = m.m[k]
	m.mu.RUnlock()
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[k] = v
	return previous, loaded
}
