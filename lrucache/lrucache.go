package lrucache

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

type LruCache struct {
	cache *expirable.LRU[string, any]
}

// size: size of array, e.g: 100 for max 100 key
// ttl time to life cache
func New(size int, ttl time.Duration) *LruCache {
	return &LruCache{
		cache: expirable.NewLRU[string, any](size, nil, ttl),
	}
}

func (lru *LruCache) Set(key string, value any) {
	lru.cache.Add(key, value)
}

func (lru *LruCache) Get(key string) any {
	r, ok := lru.cache.Get(key)

	if !ok {
		return nil
	}

	return r
}
