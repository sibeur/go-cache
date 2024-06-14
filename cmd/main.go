package main

import (
	"log"

	c "github.com/sibeur/go-cache"
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

	err = cache.Delete("key")
	if err != nil {
		panic(err)
	}

	err = cache.SetWithExpire("key", "value", 10)
	if err != nil {
		panic(err)
	}
}
