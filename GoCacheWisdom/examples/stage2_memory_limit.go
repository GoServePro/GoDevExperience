package main

import (
	"fmt"
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache"
)

func main() {
	// 创建缓存：最大内存1KB，使用LRU淘汰
	c := cache.New(
		cache.WithMaxSize(1 * 1024), // 1KB
	)

	// 存储10个大字符串（每个约0.2KB，总2KB，超过1KB）
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := string(make([]byte, 200))
		c.Set(key, value)
	}

	// 验证：保留的key
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		_, ok := c.Get(key)
		fmt.Printf("Get %s: %v\n", key, ok)
	}
}
