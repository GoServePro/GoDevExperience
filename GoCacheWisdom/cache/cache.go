package cache

import "github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache/core"

// Cache 是 cache 包的本地类型，封装 core.Core
type Cache struct {
	core *core.Core
}

// New 创建缓存实例（对外暴露的入口）
func New() *Cache {
	return &Cache{
		core: core.NewCore(),
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
