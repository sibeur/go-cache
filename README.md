# Go Cache Library

## Description
The Go Cache Library is an amazing caching solution designed for Go applications. It's super simple and easy to use, making it the perfect choice for developers of all skill levels. The library supports two types of drivers: an in-memory driver for simple, straightforward caching needs, and a Redis driver for more robust, distributed caching scenarios. You can easily switch between the two drivers, giving you the flexibility and adaptability to meet your application's caching needs.

## Features
* Simple
* Two option of driver (memory or redis)

## Installation
```bash
go get github.com/sibeur/go-sibeur/cache
```

## Usage

```go
package main

import (
	"log"

	c "github.com/sibeur/go-sibeur/cache"
)

func main() {
	cache := c.NewCache()
	err := cache.Set("key", "value")
	if err != nil {
		panic(err)
	}

	value, err := cache.Get("key")
	if err != nil {
		panic(err)
	}

	log.Printf("Value: %s", value)
}
```

The `NewCache` function is used to create a new cache instance. It doesn't take any parameters because it's designed to read from environment variables to configure the cache. Example of environment variables

| Environment Variable | Description               | Default |
|----------------------|---------------------------|---------|
| `CACHE_TYPE`         | Type of cache to use      | "memory"|
| `REDIS_ADDR`         | Redis server address      | "localhost:6379"      |
| `REDIS_PASSWORD`     | Password for Redis server |       |
| `REDIS_DB`           | Redis database name       | 0     |



### Get cache data

``` go
value, err := cache.Get("key")
if err != nil {
    panic(err)
}
log.Printf("Value: %s", value)
```

### Set cache data

``` go
err := cache.Set("key", "value")
if err != nil {
    panic(err)
}
```

### Set cache data with expirity

Expirity TTL unit in seconds
``` go
err := cache.SetWithExpire("key", "value", 10)
if err != nil {
    panic(err)
}
```

### Delete cache data
``` go	
err := cache.Delete("key")
if err != nil {
    panic(err)
}
```

#### Flush cache data
Clear all cache data
``` go	
err := cache.Flush()
if err != nil {
    panic(err)
}
```


## Memory vs Redis driver

If you're using this library on a single Go app instance, the memory driver is a great option for simplicity. But if your app is deployed in multiple instances (a distributed system or microservice), I highly recommend the Redis driver!