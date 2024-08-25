package cache

import (
	"container/list"
	"sync"
	"time"
)
//item :to unmarshal the json data
type Item struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration string `json:"expiration"`
}

// CacheItem is the type of the value that is stored in the cache
type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration time.Time
}

// LRUCache is a LRU cache
type LRUCache struct {
	mu        sync.Mutex
	capacity  int
	cache     map[string]*list.Element
	evictList *list.List
}

// NewLRUCache creates a new LRUCache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:  capacity,
		cache:     make(map[string]*list.Element),
		evictList: list.New(),
	}
}

// Set adds a value to the cache
func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.evictList.MoveToFront(ele)
		ele.Value.(*CacheItem).Value = value
		ele.Value.(*CacheItem).Expiration = time.Now().Add(expiration)
	} else {
		item := &CacheItem{
			Key:        key,
			Value:      value,
			Expiration: time.Now().Add(expiration),
		}
		ele := c.evictList.PushFront(item)
		c.cache[key] = ele

		if c.evictList.Len() > c.capacity {
			c.removeOldest()
		}
	}
}

// Get returns a value from the cache
func (c *LRUCache) Get(key string) (*CacheItem, bool) {
	// Lock the cache
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if the key exists in the cache
	if ele, ok := c.cache[key]; ok {
		// Check if the item has expired
		if ele.Value.(*CacheItem).Expiration.After(time.Now()) {
			// Move the item to the front of the list
			c.evictList.MoveToFront(ele)
			return ele.Value.(*CacheItem), true
		}
		// If the item has expired, remove it from the cache
		c.removeElement(ele)
		return nil, false
	}
	return nil, false
}

// Delete removes a value from the cache
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ele, ok := c.cache[key]; ok {
		c.removeElement(ele)
	}
}

// removeOldest removes the oldest item from the cache
func (c *LRUCache) removeOldest() {
	if ele := c.evictList.Back(); ele != nil {
		c.removeElement(ele)
	}
}

// removeElement removes a given element from the cache
func (c *LRUCache) removeElement(ele *list.Element) {
	c.evictList.Remove(ele)
	delete(c.cache, ele.Value.(*CacheItem).Key)
}

// CleanupExpiredItems removes all expired items from the cache
func (c *LRUCache) CleanupExpiredItems() {
	for {
		time.Sleep(time.Second)
		c.mu.Lock()
		for _, ele := range c.cache {
			if ele.Value.(*CacheItem).Expiration.Before(time.Now()) {
				c.removeElement(ele)
			}
		}
		c.mu.Unlock()
	}
}

// GetCacheState returns the current state of the cache
func (c *LRUCache) GetCacheState() []map[string]interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	state := []map[string]interface{}{}
	for _, v := range c.cache {
		item := v.Value.(*CacheItem)
		if item.Expiration.After(time.Now()) {
			state = append(state, map[string]interface{}{
				"key":        item.Key,
				"value":      item.Value,
				"expiration": item.Expiration.Format(time.RFC3339),
			})
		}
	}
	return state
}
