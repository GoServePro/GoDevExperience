package core

import (
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache/policy"
	"sync"
)

// Core 核心缓存结构（封装存储、锁、内存统计）
type Core struct {
	data      map[string]*Item      // 核心存储 map
	mu        sync.RWMutex          // 读写锁（保证并发安全）
	totalSize int64                 // 缓存总内存占用
	maxSize   int64                 // 最大内存限制（0=无限制）
	eviction  policy.EvictionPolicy // 淘汰策略接口
}

// Option 配置选项（建造者模式，支持后续扩展更多配置）
type Option func(*Core)

// WithMaxSize 设置最大内存限制（字节）
func WithMaxSize(maxSize int64) Option {
	return func(c *Core) {
		c.maxSize = maxSize
	}
}

// WithEvictionPolicy 设置淘汰策略（默认LRU）
func WithEvictionPolicy(ep policy.EvictionPolicy) Option {
	return func(c *Core) {
		c.eviction = ep
	}
}

// NewCore 创建核心缓存实例(支持配置选项)
func NewCore(opts ...Option) *Core {
	c := &Core{
		data:     make(map[string]*Item),
		eviction: policy.NewLRU(), // 默认LRU淘汰策略
	}
	// 应用配置选项
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Set 存储缓存（线程安全）
func (c *Core) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item := NewItem(key, value)
	// 若 key 已存在，先清理旧条目内存
	if oldItem, exists := c.data[key]; exists {
		c.totalSize -= oldItem.Size
		if c.maxSize > 0 {
			c.eviction.OnDelete(key) // 从淘汰策略中移除旧key
		}
	}

	// 检查内存是否超限，若超限则淘汰key直到内存足够
	for c.maxSize > 0 && c.totalSize+item.Size > c.maxSize {
		evictKey := c.eviction.Evict() //获取需要淘汰key
		if evictKey == "" {
			break // 无key可淘汰，直接存储（允许超内存，避免死循环）
		}
		if evictItem, exists := c.data[evictKey]; exists {
			c.totalSize -= evictItem.Size
			delete(c.data, evictKey)
		}
	}

	// 存储新条目
	c.data[key] = item
	c.totalSize += item.Size
	if c.maxSize > 0 {
		c.eviction.OnAdd(key) // 将新key添加到淘汰策略
	}
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
	// 更新访问时间和淘汰策略状态
	item.UpdateAccess()
	if c.maxSize > 0 {
		c.eviction.OnAccess(key)
	}
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
		if c.maxSize > 0 {
			c.eviction.OnDelete(key) // 从淘汰策略中移除
		}
	}
}
