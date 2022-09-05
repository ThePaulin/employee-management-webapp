package cache

import (
	"errors"
	"sync"
	"time"
)

var ErrItemNotFound = errors.New("cache: item not found")

type item struct {
	value     interface{}
	createdAt int64
	ttl       int64
}

type MemoryCache struct {
	cache map[interface{}]*item
	sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{cache: make(map[interface{}]*item)}

	go c.setTtlTimer()

	return c
}

func (cache *MemoryCache) setTtlTimer() {
	for {
		cache.Lock()
		for k, v := range cache.cache {
			if time.Now().Unix()-v.createdAt > v.ttl {
				delete(cache.cache, k)
			}
		}

		cache.Unlock()
		<-time.After(time.Second)
	}
}

func (cache *MemoryCache) Set(key, value interface{}, ttl int64) error {
	cache.Lock()
	cache.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}

	cache.Unlock()

	return nil
}

func (cache *MemoryCache) Get(key interface{}) (interface{}, error) {
	cache.RLock()
	item, ex := cache.cache[key]

	cache.RUnlock()

	if !ex {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}
