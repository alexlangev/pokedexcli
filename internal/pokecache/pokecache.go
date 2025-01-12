package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte // raw bytes were caching
}

type cache struct {
	ca map[string]cacheEntry
	mu *sync.Mutex
}

func (c *cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.c {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}

func NewCache(interval time.Time) {
	c := cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.readLoop(interval)

	return c
}
