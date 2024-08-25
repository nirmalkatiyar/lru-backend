package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirmalkatiyar/cache"
)

var LruCache = cache.NewLRUCache(5)

// GetCacheItem returns a cache item
func GetCacheItem(c *gin.Context) {
	
	key := c.Request.URL.Query().Get("key")
	fmt.Println("key", key)
	item, found := LruCache.Get(key)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid key Id", "message": "Item not found for key:" + key})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": item.Key, "value": item.Value, "expirationTime": item.Expiration})
}

// SetCacheItem sets a cache item
func SetCacheItem(c *gin.Context) {
	// Get the key, value and expiration time from the request
	item := cache.Item{}
	r, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("unable to read", err)
	}
	err=json.Unmarshal(r, &item)
	if err!=nil{
		fmt.Print("unable to unmarshal", err)
	}
	key := item.Key
	_, found := LruCache.Get(key)
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}
	if found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "message": "Item already exists for key:" + key})
		return
	}
	value := item.Value
	if value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value"})
		return
	}
	expirationStr := item.Expiration
	if expirationStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration time"})
		return
	}

	expiration, err := strconv.Atoi(expirationStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration time"})
		return
	}
	// Set the cache item
	LruCache.Set(key, value, time.Duration(expiration)*time.Second)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Key set successfully"})
}

// DeleteCacheItem deletes a cache item
func DeleteCacheItem(c *gin.Context) {
	key := c.Param("key")
	_, found := LruCache.Get(key)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"key": key, "error": "Invalid key", "message": "Item doest not exist."})
		return
	}
	LruCache.Delete(key)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "key": key, "message": "Key deleted successfully"})
}
