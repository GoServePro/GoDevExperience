package core

import (
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/internal"
	"time"
)

// Item 缓存条目（存储key-value及元信息）
type Item struct {
	Key      string      // 缓存key
	Value    interface{} // 缓存值
	AccessAt time.Time   // 最后访问时间（用于LRU淘汰）
	Size     int64       // 条目占用内存大小（近似值）
}

// NewItem 创建缓存条目（计算初始大小）
func NewItem(key string, value interface{}) *Item {
	return &Item{
		Key:      key,
		Value:    value,
		AccessAt: time.Now(),
		Size:     internal.CalculateSize(value),
	}
}

// UpdateAccess 更新访问时间（LRU淘汰时用到）
func (i *Item) UpdateAccess() {
	i.AccessAt = time.Now()
}
