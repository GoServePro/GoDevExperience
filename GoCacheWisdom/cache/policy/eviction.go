package policy

// EvictionPolicy 淘汰策略接口（定义统一方法，支持扩展）
type EvictionPolicy interface {
	// OnAccess 当key被访问时，更新策略状态（如LRU更新访问顺序）
	OnAccess(key string)
	// OnAdd 当新增key时，添加到策略中
	OnAdd(key string)
	// OnDelete 当key被删除时，从策略中移除
	OnDelete(key string)
	// Evict 淘汰一个key（返回要淘汰的key）
	Evict() string
}
