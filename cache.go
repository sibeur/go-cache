package cache

import (
	"os"
	"strconv"

	redis "github.com/redis/go-redis/v9"
	"github.com/sibeur/go-sibeur/cache/driver"
)

// Cache represents the interface for interacting with a cache.
type Cache interface {
	// Get retrieves the value associated with the given key from the cache.
	// It returns the value as a string and an error if the key is not found or an error occurs.
	Get(key string) (string, error)

	// Set sets the value associated with the given key in the cache.
	// It returns an error if an error occurs while setting the value.
	Set(key string, value string) error

	// SetWithExpire sets the value associated with the given key in the cache with a specified time-to-live (TTL).
	// The value will expire and be automatically deleted from the cache after the specified TTL (in seconds).
	// It returns an error if an error occurs while setting the value.
	SetWithExpire(key string, value string, ttl uint64) error

	// Delete deletes the value associated with the given key from the cache.
	// It returns an error if an error occurs while deleting the value.
	Delete(key string) error

	// Flush deletes all the values stored in the cache.
	// It returns an error if an error occurs while flushing the cache.
	Flush() error

	// IsCacheAvailable checks if the cache is available for use.
	// It returns true if the cache is available, false otherwise.
	IsCacheAvailable() bool

	// SetCacheAvailable sets the availability of the cache.
	// It takes a boolean value indicating whether the cache is available or not.
	SetCacheAvailable(available bool)

	// GetDriverName returns the name of the cache driver being used.
	GetDriverName() string
}

// NewCache creates a new cache based on the value of the CACHE_TYPE environment variable.
// If CACHE_TYPE is not set, it defaults to "redis".
// The function returns a Cache interface that can be used to interact with the cache.
// The supported cache types are "redis" and "memory".
// For "redis" cache type, it uses the REDIS_ADDR and REDIS_PASSWORD environment variables to connect to the Redis server.
// If REDIS_DB is set, it uses that value as the Redis database number, otherwise it uses the default database.
// For "memory" cache type, it takes a time duration as a parameter to specify the expiration time for the cache entries.
// If an unsupported cache type is specified, the function panics with an error message.
func NewCache() Cache {
	var err error
	var cache Cache
	// discover cache type
	cacheType := os.Getenv("CACHE_TYPE")
	if cacheType == "" {
		cacheType = "memory"
	}

	switch cacheType {
	case "redis":
		// load redis
		redisDb := 0
		if os.Getenv("REDIS_DB") != "" {
			redisDb, err = strconv.Atoi(os.Getenv("REDIS_DB"))
			if err != nil {
				panic(err)
			}
		}
		redisClient := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       redisDb,                     // use default DB
		})

		// load cache
		cache = driver.NewRedisCache(redisClient)

	case "memory":
		// load memory cache
		cache = driver.NewMemoryCache()
	default:
		panic("Cache type not supported")
	}
	return cache
}
