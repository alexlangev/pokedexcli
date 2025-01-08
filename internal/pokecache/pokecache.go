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

func (*cache) Add(key string, val []byte) {
	
}

func (*cache) Get(key string) ([]byte bool) {
}

func (*cache) reapLoop() {
}

func NewCache(t time.Time) {
}
