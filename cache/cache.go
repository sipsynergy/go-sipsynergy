package cache

import "time"

// Cache is the default interface for our app caching.
// This interface follows the redis api, with a same difference. That
// difference is the expansion of words, example: del > delete.
type Cache interface {
	// Get returns the item
	Get(key string) (*Item, error)
	// Set saves the given value to the cache.
	Set(key string, value interface{}, ttl time.Duration) (bool, error)
	// Delete deletes the given key
	Delete(key string) (bool, error)
	// Exists returns true or false depending on if it can find the key.
	Exists(key string) (bool, error)
	// FlushAll with delete all keys
	FlushAll() (bool, error)
	// ExpireKeys loops over each key indefinetly and checks to see if it has expired.
	// If it has expired it will be removed from the cache.
	ExpireKeys()
}

// Item is used within the cache to cleanly store data.
type Item struct {
	Key    string
	Value  interface{}
	Expiry time.Time
}
