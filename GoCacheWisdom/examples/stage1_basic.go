package main

import (
	"fmt"
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache"
)

func main() {
	// 创建基础缓存实例
	c := cache.New()

	// Set 数据
	c.Set("user:1001", "张三")
	c.Set("user:1002", map[string]int{"age": 25})

	// Get 数据
	val1, ok1 := c.Get("user:1001")
	fmt.Printf("user:1001 -> %v, exists: %v\n", val1, ok1) // 输出：张三, true

	val2, ok2 := c.Get("user:1002")
	fmt.Printf("user:1002 -> %v, exists: %v\n", val2, ok2) // 输出：map[age:25], true

	// Delete 数据
	c.Delete("user:1001")
	val3, ok3 := c.Get("user:1001")
	fmt.Printf("user:1001 -> %v, exists: %v\n", val3, ok3) // 输出：<nil>, false
}
