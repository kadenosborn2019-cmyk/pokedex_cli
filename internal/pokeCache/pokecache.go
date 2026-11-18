package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)

	return c
}

func (c Cache) Add(key string, value []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	return nil
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	value, ok := c.Cache[key]
	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.Cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.Cache, k)
		}
	}
}
