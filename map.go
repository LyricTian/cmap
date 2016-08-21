package cmap

import (
	"sync"
)

// NewMap create a new map
func NewMap() *Map {
	return &Map{
		store: make(map[interface{}]interface{}),
	}
}

// Map A thread-safe anything to anything map
type Map struct {
	sync.RWMutex
	store map[interface{}]interface{}
}

// Set the given value under the specified key
func (m *Map) Set(key, value interface{}) {
	m.Lock()
	m.store[key] = value
	m.Unlock()
}

// Get retrieves an element from map under given key
func (m *Map) Get(key interface{}) (value interface{}, ok bool) {
	m.RLock()
	value, ok = m.store[key]
	m.RUnlock()
	return
}

// Remove removes an element from the map
func (m *Map) Remove(key interface{}) {
	m.Lock()
	delete(m.store, key)
	m.Unlock()
}

// Count the number of elements within the map
func (m *Map) Count() (count int) {
	m.RLock()
	count = len(m.store)
	m.RUnlock()
	return
}

// Clear removes all elements from the map
func (m *Map) Clear() {
	m.Lock()
	m.store = make(map[interface{}]interface{})
	m.Unlock()
}

// Items returns all items as map[interface{}]interface{}
func (m *Map) Items() (items map[interface{}]interface{}) {
	m.RLock()
	items = make(map[interface{}]interface{})
	for k, v := range m.store {
		items[k] = v
	}
	m.RUnlock()
	return
}
