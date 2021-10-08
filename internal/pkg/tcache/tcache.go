package tcache

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type TCache struct {
	_cache *ristretto.Cache
}

func New() *TCache {
	c, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e4,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 20, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	return &TCache{
		_cache: c,
	}
}

func (c *TCache) Get(key string) interface{} {
	val, ok := c._cache.Get(key)
	if !ok {
		return nil
	}
	return val
}

func (c *TCache) GetAndFlush(key string, ttl time.Duration) interface{} {
	val, ok := c._cache.Get(key)
	if !ok {
		return nil
	}
	c._cache.SetWithTTL(key, val, 1, ttl)
	return val
}

func (c *TCache) Set(key string, val interface{}) bool {
	return c._cache.Set(key, val, 1)
}

func (c *TCache) SetWithTTL(key string, val interface{}, ttl time.Duration) bool {
	return c._cache.SetWithTTL(key, val, 1, ttl)
}

func (c *TCache) Del(key string) {
	c._cache.Del(key)
}
