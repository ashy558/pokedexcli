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
