package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entries: map[string]cacheEntry{},
		lock:    &sync.RWMutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
