package driver

import (
	"log"
	"sync"
	"time"

	"github.com/sibeur/go-sibeur/cache/common"
)

type MemoryCache struct {
	isAvailable bool
	data        map[string]interface{}
	mutex       sync.RWMutex
	expire      time.Duration
	cleanup     *time.Timer
	cleanupM    sync.Mutex
	driverName  string
}

func NewMemoryCache(expire time.Duration) *MemoryCache {
	driverName := "MemoryCache"
	log.Printf("[%s] initiate cache", driverName)
	cache := &MemoryCache{
		data:        make(map[string]interface{}),
		expire:      expire,
		cleanup:     time.NewTimer(expire),
		isAvailable: true,
		driverName:  driverName,
	}
	go cache.startCleanup()
	return cache
}

func (c *MemoryCache) startCleanup() {
	for {
		select {
		case <-c.cleanup.C:
			c.cleanupExpired()
			c.cleanup.Reset(c.expire)
		}
	}
}

func (c *MemoryCache) cleanupExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key := range c.data {
		if c.isExpired(key, now) {
			delete(c.data, key)
		}
	}
}

func (c *MemoryCache) isExpired(key string, now time.Time) bool {
	expiration, ok := c.data[key+"_expiration"].(time.Time)
	return ok && expiration.Before(now)
}

func (c *MemoryCache) Set(key string, value string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	log.Printf("[%s] %s %s\n", c.driverName, common.SetCacheMsg, key)
	return nil
}

func (c *MemoryCache) SetWithExpire(key string, value string, ttl uint64) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	c.data[key+"_expiration"] = time.Now().Add(time.Duration(ttl) * time.Second)
	log.Printf("[%s] %s %s with TTL %d\n", c.driverName, common.SetCacheMsg, key, ttl)
	return nil
}

func (c *MemoryCache) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.data[key].(string)
	if !ok {
		return "", nil
	}
	log.Printf("[%s] %s %s\n", c.driverName, common.GetCacheMsg, key)
	return value, nil
}

func (c *MemoryCache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
	log.Printf("[%s] %s %s\n", c.driverName, common.DeleteCacheMsg, key)
	return nil
}

func (c *MemoryCache) Flush() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data = make(map[string]interface{})
	log.Printf("[%s] %s\n", c.driverName, common.FlushCacheMsg)
	return nil
}

func (c *MemoryCache) IsCacheAvailable() bool {
	return c.isAvailable
}

func (c *MemoryCache) SetCacheAvailable(available bool) {
	c.isAvailable = available
}

func (c *MemoryCache) GetDriverName() string {
	return c.driverName
}
