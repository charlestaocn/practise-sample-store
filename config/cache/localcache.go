package cache

import (
	"sync"
	"time"
)

// Cache 是一个简单的本地缓存实现
type Cache struct {
	data sync.Map
}

// NewCache 创建一个新的缓存实例
func NewCache() *Cache {
	return &Cache{data: sync.Map{}}
}

// Get 从缓存中获取值
func (c *Cache) Get(key string) (interface{}, bool) {
	value, ok := c.data.Load(key)
	return value, ok
}

// Set 向缓存中设置值
func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

// Delete 从缓存中删除一个键
func (c *Cache) Delete(key string) {
	c.data.Delete(key)
}

// ExpireAfter 设置缓存项的有效期
func (c *Cache) ExpireAfter(key string, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		c.Delete(key)
	}()
}
