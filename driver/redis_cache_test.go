package driver_test

import (
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sibeur/go-cache/driver"
)

func TestRedisCache_Get(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value matches the expected value
	expectedValue := "value"
	if value != expectedValue {
		t.Errorf("Expected value %s, but got %s", expectedValue, value)
	}
}

func TestRedisCache_Set(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value matches the expected value
	expectedValue := "value"
	if value != expectedValue {
		t.Errorf("Expected value %s, but got %s", expectedValue, value)
	}
}

func TestRedisCache_SetWithExpire(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	key := "setwithexpireredis"
	value := "value1"
	// Set a value in the cache with an expiry time of 1 second
	err := cache.SetWithExpire(key, value, 1)
	if err != nil {
		t.Errorf("Failed to set value in cache with expiry: %v", err)
	}

	// Wait for the value to expire
	// Sleep for 2 seconds to ensure the value has expired
	time.Sleep(2 * time.Second)

	// Get the value from the cache
	value, err = cache.Get(key)
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_Delete(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Delete the value from the cache
	err = cache.Delete("key")
	if err != nil {
		t.Errorf("Failed to delete value from cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_Flush(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}

	// Flush all keys in the cache
	err = cache.Flush()
	if err != nil {
		t.Errorf("Failed to flush cache: %v", err)
	}

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil && err != redis.Nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_IsCacheAvailable(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Check if the cache is available
	if !cache.IsCacheAvailable() {
		t.Errorf("Expected cache to be available, but got unavailable")
	}
}

func TestRedisCache_SetCacheAvailable(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Set the cache to be unavailable
	cache.SetCacheAvailable(false)

	// Check if the cache is unavailable
	if cache.IsCacheAvailable() {
		t.Errorf("Expected cache to be unavailable, but got available")
	}
}

func TestRedisCache_GetDriverName(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Get the driver name
	driverName := cache.GetDriverName()

	// Check if the driver name matches the expected value
	expectedDriverName := "redis"
	if driverName != expectedDriverName {
		t.Errorf("Expected driver name %s, but got %s", expectedDriverName, driverName)
	}
}

func TestRedisCache_Set_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Set a value in the cache
	err := cache.Set("key", "value")
	if err != nil {
		t.Errorf("Failed to set value in cache: %v", err)
	}
}

func TestRedisCache_Get_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Get the value from the cache
	value, err := cache.Get("key")
	if err != nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_SetWithExpire_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	key := "setwithexpireredis"
	value := "value1"
	// Set a value in the cache with an expiry time of 1 second
	err := cache.SetWithExpire(key, value, 1)
	if err != nil {
		t.Errorf("Failed to set value in cache with expiry: %v", err)
	}

	// Wait for the value to expire
	// Sleep for 2 seconds to ensure the value has expired
	time.Sleep(2 * time.Second)

	// Get the value from the cache
	value, err = cache.Get(key)
	if err != nil {
		t.Errorf("Failed to get value from cache: %v", err)
	}

	// Check if the retrieved value is empty
	if value != "" {
		t.Errorf("Expected value to be empty, but got %s", value)
	}
}

func TestRedisCache_Delete_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Delete the value from the cache
	err := cache.Delete("key")
	if err != nil {
		t.Errorf("Failed to delete value from cache: %v", err)
	}
}

func TestRedisCache_Flush_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Flush all keys in the cache
	err := cache.Flush()
	if err != nil {
		t.Errorf("Failed to flush cache: %v", err)
	}
}

func TestRedisCache_GetDriverName_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Get the driver name
	driverName := cache.GetDriverName()

	// Check if the driver name matches the expected value
	expectedDriverName := "redis"
	if driverName != expectedDriverName {
		t.Errorf("Expected driver name %s, but got %s", expectedDriverName, driverName)
	}
}

func TestRedisCache_IsCacheAvailable_UnavailableCache(t *testing.T) {
	// Create a Redis client for testing
	client := redis.NewClient(&redis.Options{
		Addr: "invalid",
	})

	// Create a RedisCache instance
	cache := driver.NewRedisCache(client)

	// Check if the cache is available
	if cache.IsCacheAvailable() {
		t.Errorf("Expected cache to be unavailable, but got available")
	}
}
