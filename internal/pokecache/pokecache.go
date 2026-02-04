package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mut *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mut: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reap(now time.Time, last time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}
