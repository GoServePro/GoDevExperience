package cache

import (
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache/core"
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache/policy"
)

// Cache 是 cache 包的本地类型，封装 core.Core（对外隐藏 core 实现）
type Cache struct {
	core *core.Core
}

// WithMaxSize 设置缓存最大内存限制（字节）
// 对外导出，用户通过 cache.WithMaxSize() 调用，无需关心 core 包
func WithMaxSize(maxSize int64) core.Option {
	return core.WithMaxSize(maxSize)
}

// WithEvictionPolicy 设置缓存淘汰策略（默认 LRU）
// 支持用户传入自定义策略（需实现 policy.EvictionPolicy 接口）
func WithEvictionPolicy(ep policy.EvictionPolicy) core.Option {
	return core.WithEvictionPolicy(ep)
}

// New 创建缓存实例（对外暴露的入口）
func New(opts ...core.Option) *Cache {
	coreInstance := core.NewCore(opts...)
	return &Cache{
		core: coreInstance,
	}
}

// Set 存储缓存（封装 core.Core.Set）
func (c *Cache) Set(key string, value interface{}) {
	c.core.Set(key, value)
}

// Get 获取缓存（封装 core.Core.Get）
func (c *Cache) Get(key string) (interface{}, bool) {
	return c.core.Get(key)
}

// Delete 删除缓存（封装 core.Core.Delete）
func (c *Cache) Delete(key string) {
	c.core.Delete(key)
}
