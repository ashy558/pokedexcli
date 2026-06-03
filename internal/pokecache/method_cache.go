package pokecache

import "time"

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
	ticker := time.Tick(interval)
	for tick := range ticker {
		for key, entry := range c.entries {
			if tick.Sub(entry.createdAt) > interval {
				c.lock.Lock()
				delete(c.entries, key)
				c.lock.Unlock()
			}
		}
	}
}
