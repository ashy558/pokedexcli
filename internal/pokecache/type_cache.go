package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	lock    *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: map[string]cacheEntry{},
		lock:    &sync.RWMutex{},
	}
	newCache.reapLoop(interval)
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) (val []byte, ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if entry, ok := c.entries[key]; !ok {
		return val, ok
	} else {
		return entry.val, ok
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	for {
		time.Tick(interval)
		now := time.Now()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > interval {
				c.lock.Lock()
				delete(c.entries, key)
				c.lock.Unlock()
			}
		}
	}
}
