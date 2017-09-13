package cache

import (
	"fmt"
	"sync"
	"time"
)

// NewDefaultCache returns an instance of the default cache.
// Also creates a go routine to periodically check for expired keys.
func NewDefaultCache() *DefaultCache {
	c := &DefaultCache{
		items: make(map[string]*Item),
		m:     &sync.RWMutex{},
	}

	go func(c Cache) {
		for {
			time.Sleep(time.Minute)
			c.ExpireKeys()
		}
	}(c)

	return c
}

// DefaultCache is the default implementation of the cache interface.
type DefaultCache struct {
	items map[string]*Item
	m     *sync.RWMutex
}

// Get returns the item
func (c *DefaultCache) Get(key string) (i *Item, err error) {
	c.m.RLock()
	i, ok := c.items[key]
	c.m.RUnlock()

	if !ok {
		err = fmt.Errorf("item not found for key '%s'", key)
	}

	if c.hasExpired(i) {
		c.Delete(key)

		err = fmt.Errorf("item not found for key '%s'", key)
	}

	return
}

// Set saves the given value to the cache.
func (c *DefaultCache) Set(key string, value interface{}, ttl time.Duration) (saved bool, err error) {
	c.m.Lock()
	c.items[key] = &Item{
		Key:    key,
		Value:  value,
		Expiry: time.Now().Add(ttl),
	}
	c.m.Unlock()

	saved = true

	return
}

// Delete deletes the given key
func (c *DefaultCache) Delete(key string) (bool, error) {
	if ok, _ := c.Exists(key); !ok {
		return true, nil
	}

	c.m.Lock()
	delete(c.items, key)
	c.m.Unlock()

	return true, nil
}

// Exists returns true or false depending on if it can find the key.
func (c *DefaultCache) Exists(key string) (bool, error) {
	c.m.RLock()
	i, ok := c.items[key]
	c.m.RUnlock()

	if !ok {
		return false, nil
	}

	if c.hasExpired(i) {
		return false, nil
	}

	return true, nil
}

// FlushAll with delete all keys
func (c *DefaultCache) FlushAll() (bool, error) {
	c.m.Lock()
	c.items = make(map[string]*Item)
	c.m.Unlock()

	return true, nil
}

// ExpireKeys loops over each key indefinetly and checks to see if it has expired.
// If it has expired it will be removed from the cache.
func (c *DefaultCache) ExpireKeys() {
	for i := range c.items {
		item := c.items[i]

		if c.hasExpired(item) {
			c.m.Lock()
			c.Delete(item.Key)
			c.m.Unlock()
		}
	}
}

// HasExpired returns true if the item has expired.
func (c *DefaultCache) hasExpired(i *Item) bool {
	return time.Now().After(i.Expiry)
}
