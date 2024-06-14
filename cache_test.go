package cache_test

import (
	"os"
	"testing"

	"github.com/sibeur/go-sibeur/cache"
)

func TestNewCache(t *testing.T) {
	os.Setenv("CACHE_TYPE", "memory")
	cache := cache.NewCache()
	if cache.GetDriverName() != "memory" {
		t.Errorf("Expected driver name %s, but got %s", "memory", cache.GetDriverName())
	}
}

func TestNewCacheWithRedis(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "0")
	cache := cache.NewCache()
	if cache.GetDriverName() != "redis" {
		t.Errorf("Expected driver name %s, but got %s", "redis", cache.GetDriverName())
	}
}

func TestNewCacheWithInvalidType(t *testing.T) {
	os.Setenv("CACHE_TYPE", "invalid")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got nil")
		}
	}()
	cache.NewCache()
}

func TestNewCacheWithInvalidRedisDB(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "invalid")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got nil")
		}
	}()
	cache.NewCache()
}

func TestNewCacheWithInvalidRedisAddr(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "invalid")
	os.Setenv("REDIS_DB", "0")
	c := cache.NewCache()
	if c.IsCacheAvailable() {
		t.Errorf("Expected cache available %v, but got %v", false, c.IsCacheAvailable())
	}
}

func TestNewCacheWithInvalidRedisPassword(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "0")
	os.Setenv("REDIS_PASSWORD", "invalid")
	c := cache.NewCache()
	if !c.IsCacheAvailable() {
		t.Errorf("Expected cache available %v, but got %v", true, c.IsCacheAvailable())
	}
}

func TestNewCacheWithEmptyRedisPassword(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "0")
	os.Setenv("REDIS_PASSWORD", "")
	cache := cache.NewCache()
	if cache.GetDriverName() != "redis" {
		t.Errorf("Expected driver name %s, but got %s", "redis", cache.GetDriverName())
	}
}

func TestNewCacheWithEmptyRedisDB(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_DB", "")
	cache := cache.NewCache()
	if cache.GetDriverName() != "redis" {
		t.Errorf("Expected driver name %s, but got %s", "redis", cache.GetDriverName())
	}
}

func TestNewCacheWithEmptyRedisAddr(t *testing.T) {
	os.Setenv("CACHE_TYPE", "redis")
	os.Setenv("REDIS_ADDR", "")
	os.Setenv("REDIS_DB", "0")
	c := cache.NewCache()
	if !c.IsCacheAvailable() {
		t.Errorf("Expected cache available %v, but got %v", true, c.IsCacheAvailable())
	}
}

func TestNewCacheWithEmptyType(t *testing.T) {
	os.Setenv("CACHE_TYPE", "")
	cache := cache.NewCache()
	if cache.GetDriverName() != "memory" {
		t.Errorf("Expected driver name %s, but got %s", "memory", cache.GetDriverName())
	}
}
