package driver

import (
	"context"
	"log"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/sibeur/go-sibeur/cache/common"
)

// RedisCache represents a cache driver that uses Redis as the underlying storage.
type RedisCache struct {
	client      *redis.Client
	isAvailable bool
	driverName  string
}

// NewRedisCache creates a new instance of RedisCache using the provided Redis client.
// It initializes the cache and checks if the Redis server is available by pinging it.
// If the server is unavailable, the cache will be marked as unavailable.
// The function returns a pointer to the created RedisCache instance.
func NewRedisCache(client *redis.Client) *RedisCache {
	driverName := "redis"
	log.Printf("[%s] initiate cache", driverName)
	// ping redis
	pong, _ := client.Ping(context.Background()).Result()
	available := true
	if pong == "" {
		log.Println(common.ErrCacheUnavailableMsg)
		available = false
	}

	return &RedisCache{
		client:      client,
		isAvailable: available,
		driverName:  driverName,
	}
}

// Get retrieves the value associated with the given key from the Redis cache.
// If the cache is unavailable, it logs an error message and returns an empty string.
// It returns the value and any error encountered during the retrieval process.
func (r *RedisCache) Get(key string) (string, error) {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return "", nil
	}
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	log.Printf("[%s] %s %s\n", r.driverName, common.GetCacheMsg, key)
	return val, nil
}

// Set sets the value for the given key in the Redis cache.
// If the cache is unavailable, it logs an error message and returns nil.
// It returns an error if there was a problem setting the value in the cache.
func (r *RedisCache) Set(key string, value string) error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s %s\n", r.driverName, common.SetCacheMsg, key)
	return r.client.Set(ctx, key, value, 0).Err()
}

// SetWithExpire sets a key-value pair in the Redis cache with an expiration time.
// If the cache is unavailable, it logs an error message and returns nil.
// It takes the key, value, and time-to-live (TTL in seconds) as parameters.
// The TTL specifies the duration for which the key-value pair should be stored in the cache.
// It returns an error if there was a problem setting the key-value pair in the cache.
func (r *RedisCache) SetWithExpire(key string, value string, ttl uint64) error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s %s with TTL %d\n", r.driverName, common.SetCacheMsg, key, ttl)
	return r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

// Delete removes the cache entry with the specified key from the Redis cache.
// If the cache is unavailable, it logs an error message and returns nil.
// It returns an error if there was a problem deleting the cache entry.
func (r *RedisCache) Delete(key string) error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s %s\n", r.driverName, common.DeleteCacheMsg, key)
	return r.client.Del(ctx, key).Err()
}

// Flush deletes all the keys in the cache.
func (r *RedisCache) Flush() error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s\n", r.driverName, common.FlushCacheMsg)
	return r.client.FlushAll(ctx).Err()
}

// IsCacheAvailable checks if the Redis cache is available.
// It returns true if the cache is available, otherwise false.
func (r *RedisCache) IsCacheAvailable() bool {
	return r.isAvailable
}

// SetCacheAvailable sets the availability status of the Redis cache.
func (r *RedisCache) SetCacheAvailable(available bool) {
	r.isAvailable = available
}

// GetDriverName returns the name of the Redis cache driver.
func (r *RedisCache) GetDriverName() string {
	return r.driverName
}
