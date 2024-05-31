package tests

import (
	"testing"
	"time"

	"github.com/sibeur/go-sibeur/cache/driver"
)

func TestMemoryCache_Set_Get(t *testing.T) {
	expire := time.Minute
	cache := driver.NewMemoryCache(expire)

	// Test case 1: Set a value and retrieve it
	key := "key1"
	value := "value1"
	err := cache.Set(key, value)
	if err != nil {
		t.Errorf("Failed to set value: %v", err)
	}

	retrievedValue, err := cache.Get(key)
	if err != nil {
		t.Errorf("Failed to retrieve value: %v", err)
	}
	if retrievedValue != value {
		t.Errorf("Retrieved value does not match: expected %s, got %s", value, retrievedValue)
	}

	// Test case 2: Set a value with the same key and verify it's updated
	newValue := "new value"
	err = cache.Set(key, newValue)
	if err != nil {
		t.Errorf("Failed to set updated value: %v", err)
	}

	retrievedValue, err = cache.Get(key)
	if err != nil {
		t.Errorf("Failed to retrieve updated value: %v", err)
	}
	if retrievedValue != newValue {
		t.Errorf("Retrieved updated value does not match: expected %s, got %s", newValue, retrievedValue)
	}
}

func TestMemoryCache_Delete(t *testing.T) {
	expire := time.Minute
	cache := driver.NewMemoryCache(expire)

	// Set a value
	key := "key1"
	value := "value1"
	err := cache.Set(key, value)
	if err != nil {
		t.Errorf("Failed to set value: %v", err)
	}

	t.Logf("Set value: %s", value)

	get_value, _ := cache.Get(key)

	t.Logf("Get value 1 : %s", get_value)

	if get_value != value {
		t.Errorf("Retrieved value does not match: expected %s, got %s", value, get_value)
	}

	// Delete the value
	err = cache.Delete(key)
	if err != nil {
		t.Errorf("Failed to delete value: %v", err)
	}

	// Verify that the value is deleted
	get_value, _ = cache.Get(key)
	t.Logf("Get value 2 : %s", get_value)
	if get_value != "" {
		t.Errorf("Value is not deleted: expected empty string, got %s", get_value)
	}
}

func TestMemoryCache_Flush(t *testing.T) {
	expire := time.Minute
	cache := driver.NewMemoryCache(expire)

	// Set a value
	key := "key1"
	value := "value1"
	err := cache.Set(key, value)
	if err != nil {
		t.Errorf("Failed to set value: %v", err)
	}

	t.Logf("Set value: %s", value)

	get_value, _ := cache.Get(key)

	t.Logf("Get value 1 : %s", get_value)

	if get_value != value {
		t.Errorf("Retrieved value does not match: expected %s, got %s", value, get_value)
	}

	// Flush the cache
	err = cache.Flush()
	if err != nil {
		t.Errorf("Failed to flush cache: %v", err)
	}

	// Verify that the value is deleted
	get_value, _ = cache.Get(key)
	if get_value != "" {
		t.Errorf("Value is not deleted: expected empty string, got %s", get_value)
	}
}

func TestMemoryCache_SetWithExpire(t *testing.T) {
	expire := time.Second
	cache := driver.NewMemoryCache(expire)

	// Set a value with
	key := "key1"
	value := "value1"
	err := cache.SetWithExpire(key, value, 1)
	if err != nil {
		t.Errorf("Failed to set value with expiry: %v", err)
	}

	// Wait for the value to expire
	// Sleep for 2 seconds to ensure the value has expired
	time.Sleep(2 * time.Second)

	// Get the value from the cache
	retrievedValue, err := cache.Get(key)
	if err != nil {
		t.Errorf("Failed to retrieve value: %v", err)
	}

	// Check if the retrieved value is empty
	if retrievedValue != "" {
		t.Errorf("Expected value to be empty, but got %s", retrievedValue)
	}

	// Test case 2: Set a value with expiry and verify it's updated
	newValue := "new value"
	err = cache.SetWithExpire(key, newValue, 1)
	if err != nil {
		t.Errorf("Failed to set updated value with expiry: %v", err)
	}

	// Wait for the value to expire
	// Sleep for 2 seconds to ensure the value has expired
	time.Sleep(2 * time.Second)

	// Get the value from the cache
	retrievedValue, err = cache.Get(key)
	if err != nil {
		t.Errorf("Failed to retrieve updated value: %v", err)
	}

	// Check if the retrieved value is empty
	if retrievedValue != "" {
		t.Errorf("Expected value to be empty, but got %s", retrievedValue)
	}

}