package cache

import (
	"fmt"
	"sync"

	"github.com/KnlnKS/list"
)

type kv[Value any] struct {
	key   string
	value Value
}

type LRUCache[Value any] struct {
	cache    map[string]*list.Element[kv[Value]]
	capacity int
	list     *list.List[kv[Value]]
	lock     sync.RWMutex
}

// Global Variables
var (
	ApiKeyCache *LRUCache[bool]
	once        sync.Once
)

/*
creates API key LRUCache with the given capacity.
*/
func InitApiKeyCache(capacity int) error {
	once.Do(func() {
		ApiKeyCache = New[bool](capacity)
	})
	return nil
}

/*
creates a new LRUCache with the given capacity.
*/
func New[Value any](capacity int) *LRUCache[Value] {
	return &LRUCache[Value]{
		cache:    make(map[string]*list.Element[kv[Value]], capacity),
		capacity: capacity,
		list:     list.New[kv[Value]](),
	}
}

/*
returns the value associated with the given key.
*/
func (lru *LRUCache[Value]) Get(key string) (Value, bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	elem, ok := lru.cache[key]
	if ok {
		lru.list.MoveToFront(elem)
		return elem.Value.value, ok
	}
	// *new(T) idiom -> zero value for a generic
	return *new(Value), false
}

/*
adds the given key to the cache.
*/
func (lru *LRUCache[Value]) Add(key string, value Value) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return
	}

	lru.cache[key] = lru.list.PushFront(kv[Value]{key, value})

	if lru.list.Len() > lru.capacity {
		if elem := lru.list.Back(); elem != nil {
			lru.list.Remove(elem)
			delete(lru.cache, elem.Value.key)
		}
	}
}

/*
removes the given key from the cache.
*/
func (lru *LRUCache[Value]) Remove(key string) error {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	if elem, ok := lru.cache[key]; ok {
		lru.list.Remove(elem)
		delete(lru.cache, key)
		return nil
	}
	return fmt.Errorf("%s is not in the cache", key)
}
