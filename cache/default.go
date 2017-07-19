package cache

import (
	"fmt"
	"time"
)

// NewDefaultCache returns an instance of the default cache.
func NewDefaultCache() *DefaultCache {
	return &DefaultCache{items: make(map[string]*Item)}
}

// DefaultCache is the default implementation of the cache interface.
type DefaultCache struct {
	items map[string]*Item
}

// Get returns the item
func (c *DefaultCache) Get(key string) (item *Item, err error) {
	item, ok := c.items[key]

	if !ok {
		err = fmt.Errorf("item not found for key '%s'", key)
	}

	return
}

// Set saves the given value to the cache.
func (c *DefaultCache) Set(key string, value interface{}, ttl time.Duration) (saved bool, err error) {

	c.items[key] = &Item{
		Key:    key,
		Value:  value,
		Expiry: time.Now().Add(ttl),
	}

	saved = true

	return
}

// Delete deletes the given key
func (c *DefaultCache) Delete(key string) (bool, error) {
	if ok, _ := c.Exists(key); !ok {
		return true, nil
	}

	delete(c.items, key)

	return true, nil
}

// Exists returns true or false depending on if it can find the key.
func (c *DefaultCache) Exists(key string) (bool, error) {
	_, ok := c.items[key]

	return ok, nil
}

// FlushAll with delete all keys
func (c *DefaultCache) FlushAll() (bool, error) {
	c.items = make(map[string]*Item)

	return true, nil
}

// ExpireKeys loops over each key indefinetly and checks to see if it has expired.
// If it has expired it will be removed from the cache.
func (c *DefaultCache) ExpireKeys() {
	for _, i := range c.items {
		if i.Expiry.After(time.Now()) {
			c.Delete(i.Key)
		}
	}
}
