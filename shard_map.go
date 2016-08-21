package cmap

import (
	"hash/fnv"
	"sync"
)

const (
	// DefaultShartCount The default number of shard
	DefaultShartCount = 32
)

type shardItem struct {
	sync.RWMutex
	store map[string]interface{}
}

// NewShardMap create a new shard map
func NewShardMap(shardcount ...uint32) *ShardMap {
	var count uint32
	if len(shardcount) > 0 && shardcount[0] > 0 {
		count = shardcount[0]
	} else {
		count = DefaultShartCount
	}
	shardItems := make([]*shardItem, int(count))
	for i := 0; i < int(count); i++ {
		shardItems[i] = &shardItem{
			store: make(map[string]interface{}),
		}
	}

	return &ShardMap{
		shardItems: shardItems,
		count:      count,
	}
}

// ShardMap A thread-safe string to anything shard map,
// the distribution of keys used FNV-1a hash algorithm
type ShardMap struct {
	shardItems []*shardItem
	count      uint32
}

// find a shard
func (sm *ShardMap) locate(key string) (shard *shardItem) {
	h := fnv.New32a()
	h.Write([]byte(key))
	shard = sm.shardItems[h.Sum32()%sm.count]
	return
}

// Set the given value under the specified key
func (sm *ShardMap) Set(key string, value interface{}) {
	shard := sm.locate(key)
	shard.Lock()
	shard.store[key] = value
	shard.Unlock()
}

// Get retrieves an element from map under given key
func (sm *ShardMap) Get(key string) (value interface{}, ok bool) {
	shard := sm.locate(key)
	shard.RLock()
	value, ok = shard.store[key]
	shard.RUnlock()
	return
}

// Remove removes an element from the map
func (sm *ShardMap) Remove(key string) {
	shard := sm.locate(key)
	shard.Lock()
	delete(shard.store, key)
	shard.Unlock()
}

// Count the number of elements within the map
func (sm *ShardMap) Count() (count int) {
	for i := 0; i < int(sm.count); i++ {
		shard := sm.shardItems[i]
		shard.RLock()
		count += len(shard.store)
		shard.RUnlock()
	}
	return
}

// Clear removes all elements from the map
func (sm *ShardMap) Clear() {
	for i := 0; i < int(sm.count); i++ {
		shard := sm.shardItems[i]
		shard.Lock()
		shard.store = make(map[string]interface{})
		shard.Unlock()
	}
}

// Items returns all items as map[string]interface{}
func (sm *ShardMap) Items() (items map[string]interface{}) {
	items = make(map[string]interface{})
	for i := 0; i < int(sm.count); i++ {
		shard := sm.shardItems[i]
		shard.RLock()
		for k, v := range shard.store {
			items[k] = v
		}
		shard.RUnlock()
	}
	return
}
