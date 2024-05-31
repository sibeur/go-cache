package driver

import (
	"context"
	"log"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/sibeur/go-sibeur/cache/common"
)

type RedisCache struct {
	client      *redis.Client
	isAvailable bool
	driverName  string
}

func NewRedisCache(client *redis.Client) *RedisCache {
	driverName := "RedisCache"
	log.Printf("[%s] initiate cache", driverName)
	// ping redis
	pong, _ := client.Ping(context.Background()).Result()
	if pong == "" {
		log.Println(common.ErrCacheUnavailableMsg)
	}
	available := pong == "PONG"
	return &RedisCache{
		client:      client,
		isAvailable: available,
		driverName:  driverName,
	}
}

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

func (r *RedisCache) Set(key string, value string) error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s %s\n", r.driverName, common.SetCacheMsg, key)
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisCache) SetWithExpire(key string, value string, ttl uint64) error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s %s with TTL %d\n", r.driverName, common.SetCacheMsg, key, ttl)
	return r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (r *RedisCache) Delete(key string) error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s %s\n", r.driverName, common.DeleteCacheMsg, key)
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Flush() error {
	if !r.isAvailable {
		log.Println(common.ErrCacheUnavailableMsg)
		return nil
	}
	ctx := context.Background()
	log.Printf("[%s] %s\n", r.driverName, common.FlushCacheMsg)
	return r.client.FlushAll(ctx).Err()
}

func (r *RedisCache) IsCacheAvailable() bool {
	return r.isAvailable
}

func (r *RedisCache) SetCacheAvailable(available bool) {
	r.isAvailable = available
}

func (r *RedisCache) GetDriverName() string {
	return r.driverName
}
