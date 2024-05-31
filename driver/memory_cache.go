package driver

import (
	"log"
	"sync"
	"time"

	"github.com/sibeur/go-sibeur/cache/common"
)

// MemoryCache represents an in-memory cache implementation.
type MemoryCache struct {
	isAvailable bool                   // Flag indicating if the cache is available.
	data        map[string]interface{} // The actual cache data stored as key-value pairs.
	mutex       sync.RWMutex           // Mutex for concurrent access to the cache.
	expire      time.Duration          // The duration after which cache entries expire.
	cleanup     *time.Timer            // Timer for periodic cache cleanup.
	cleanupM    sync.Mutex             // Mutex for synchronizing cache cleanup.
	driverName  string                 // The name of the cache driver.
}

// NewMemoryCache creates a new instance of the MemoryCache with the specified expiration time.
// The cache is initialized with an empty data map, a cleanup timer, and the driver name set to "memory".
// The cleanup timer is started in a separate goroutine to periodically remove expired entries from the cache.
func NewMemoryCache() *MemoryCache {
	driverName := "memory"
	log.Printf("[%s] initiate cache", driverName)
	cache := &MemoryCache{
		data:        make(map[string]interface{}),
		expire:      time.Second,
		cleanup:     time.NewTimer(time.Second),
		isAvailable: true,
		driverName:  driverName,
	}
	go cache.startCleanup()
	return cache
}

// startCleanup starts the cleanup routine for the MemoryCache.
// It continuously listens for the cleanup channel and performs
// cache cleanup by calling the cleanupExpired method and resetting
// the cleanup timer.
func (c *MemoryCache) startCleanup() {
	for {
		select {
		case <-c.cleanup.C:
			c.cleanupExpired()
			c.cleanup.Reset(c.expire)
		}
	}
}

// cleanupExpired removes all expired entries from the memory cache.
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

// isExpired checks if the cache entry with the given key has expired.
// It compares the expiration time of the cache entry with the current time.
// Returns true if the cache entry has expired, false otherwise.
func (c *MemoryCache) isExpired(key string, now time.Time) bool {
	expiration, ok := c.data[key+"_expiration"].(time.Time)
	return ok && expiration.Before(now)
}

// Set sets the value for the given key in the memory cache.
// If the key already exists, its value will be overwritten.
// The method is thread-safe.
func (c *MemoryCache) Set(key string, value string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	log.Printf("[%s] %s %s\n", c.driverName, common.SetCacheMsg, key)
	return nil
}

// SetWithExpire sets a key-value pair in the memory cache with an expiration time.
// The key-value pair will be stored in the cache for the specified TTL (time to live in seconds) duration.
// After the TTL duration has passed, the key-value pair will be automatically evicted from the cache.
// The method acquires a lock on the cache to ensure thread safety during the operation.
// The key-value pair is stored in the `data` map, and the expiration time is stored in a separate entry in the map.
// The method logs the cache operation with the driver name, the key, and the TTL.
// If an error occurs during the operation, it will be returned.
func (c *MemoryCache) SetWithExpire(key string, value string, ttl uint64) error {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	c.data[key+"_expiration"] = time.Now().Add(time.Duration(ttl) * time.Second)
	log.Printf("[%s] %s %s with TTL %d\n", c.driverName, common.SetCacheMsg, key, ttl)
	return nil
}

// Get retrieves the value associated with the given key from the memory cache.
// If the key does not exist or the value is not of type string, it returns an empty string and no error.
// It also logs the cache message with the driver name and key.
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

// Delete removes the cache entry with the specified key from the memory cache.
// It acquires a lock to ensure thread safety and then deletes the entry from the cache.
// Finally, it logs the deletion operation.
func (c *MemoryCache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
	log.Printf("[%s] %s %s\n", c.driverName, common.DeleteCacheMsg, key)
	return nil
}

// Flush clears the cache by resetting the data map to an empty map.
// It also logs a flush cache message.
func (c *MemoryCache) Flush() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data = make(map[string]interface{})
	log.Printf("[%s] %s\n", c.driverName, common.FlushCacheMsg)
	return nil
}

// IsCacheAvailable checks if the cache is available.
// It returns true if the cache is available, otherwise false.
func (c *MemoryCache) IsCacheAvailable() bool {
	return c.isAvailable
}

// SetCacheAvailable sets the availability of the memory cache.
// If available is true, the cache is marked as available. Otherwise, it is marked as unavailable.
func (c *MemoryCache) SetCacheAvailable(available bool) {
	c.isAvailable = available
}

// GetDriverName returns the name of the driver used by the MemoryCache.
func (c *MemoryCache) GetDriverName() string {
	return c.driverName
}
