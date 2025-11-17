package core

import (
	"sync"
	"time"
)

// Core 核心缓存结构（封装存储、锁、内存统计）
type Core struct {
	data      map[string]*Item // 核心存储 map
	mu        sync.RWMutex     // 读写锁（保证并发安全）
	totalSize int64            // 缓存总内存占用
}

// NewCore 创建核心缓存实例
func NewCore() *Core {
	return &Core{
		data: make(map[string]*Item),
	}
}

// Set 存储缓存（线程安全）
func (c *Core) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item := NewItem(key, value)
	// 若 key 已存在，先清理旧条目内存
	if oldItem, exists := c.data[key]; exists {
		c.totalSize -= oldItem.Size
	}
	c.data[key] = item
	c.totalSize += item.Size
}

// Get 获取缓存（线程安全，更新访问时间）
func (c *Core) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, exists := c.data[key]
	c.mu.RUnlock()

	if !exists {
		return nil, false
	}

	// 更新访问时间（加写锁）
	c.mu.Lock()
	item.AccessAt = time.Now()
	c.mu.Unlock()

	return item.Value, true
}

// Delete 删除缓存（线程安全）
func (c *Core) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, exists := c.data[key]; exists {
		c.totalSize -= item.Size
		delete(c.data, key)
	}
}
