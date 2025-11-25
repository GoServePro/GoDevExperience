package policy

import "container/list"

// LRU 最近最少使用淘汰策略
type LRU struct {
	list  *list.List               // 双向链表：存储key，按访问顺序排序（队首=最久未使用）
	index map[string]*list.Element // 哈希表：key->链表节点，O(1)查找
}

func NewLRU() *LRU {
	return &LRU{
		list:  list.New(),
		index: make(map[string]*list.Element),
	}
}

// OnAccess 访问key时，将节点移到队尾（表示最近使用）
func (l *LRU) OnAccess(key string) {
	if elem, exists := l.index[key]; exists {
		l.list.MoveToBack(elem)
	}
}

// OnAdd 新增key时，添加到队尾
func (l *LRU) OnAdd(key string) {
	// 若已存在，先删除旧节点
	if elem, exists := l.index[key]; exists {
		l.list.Remove(elem)
	}
	elem := l.list.PushBack(key)
	l.index[key] = elem
}

// OnDelete 删除key时，从链表和哈希表中移除
func (l *LRU) OnDelete(key string) {
	if elem, exists := l.index[key]; exists {
		l.list.Remove(elem)
		delete(l.index, key)
	}
}

// Evict 淘汰队首节点（最久未使用）
func (l *LRU) Evict() string {
	elem := l.list.Front()
	if elem == nil {
		return ""
	}
	l.list.Remove(elem)
	key := elem.Value.(string)
	delete(l.index, key)
	return key
}
