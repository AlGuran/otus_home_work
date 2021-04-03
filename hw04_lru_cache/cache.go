package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mx       sync.RWMutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	item, ok := lc.items[key]

	if !ok {
		if lc.queue.Len() == lc.capacity {
			last := lc.queue.Back()
			lc.queue.Remove(last)
			delete(lc.items, last.Value.(*cacheItem).key)
		}

		cacheItem := &cacheItem{key, value}
		item = lc.queue.PushFront(cacheItem)
		lc.items[cacheItem.key] = item
		return false
	}

	item.Value.(*cacheItem).value = value
	lc.queue.MoveToFront(item)
	return true
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	item, ok := lc.items[key]

	if ok {
		lc.queue.MoveToFront(item)
		return item.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mx.RLock()
	defer lc.mx.RUnlock()

	lc.queue, lc.items = NewList(), make(map[Key]*ListItem)
}
