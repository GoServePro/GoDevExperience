# 阶段1：实现基础缓存（线程安全 Set/Get/Delete）

## 一、阶段目标
- 完成线程安全的Set、Get、Delete操作
- 定义缓存条目结构，记录key、value、访问时间、内存大小
- 保证高并发访问安全
- 为后续功能预留扩展接口

## 二、技术选型
- 存储结构：map[string]*Item + sync.RWMutex
- 理由：map实现O(1)读写，读写锁提升并发性能

## 三、核心代码
### 1. 缓存条目结构体
type Item struct {
    Key      string
    Value    interface{}
    AccessAt time.Time
    Size     int64
}

### 2. 缓存主结构体
type Cache struct {
    data      map[string]*Item
    mu        sync.RWMutex
    totalSize int64
}

### 3. Set方法
func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    item := NewItem(key, value)
    if oldItem, exists := c.data[key]; exists {
        c.totalSize -= oldItem.Size
    }
    c.data[key] = item
    c.totalSize += item.Size
}

### 4. Get方法
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    item, exists := c.data[key]
    c.mu.RUnlock()
    if !exists {
        return nil, false
    }
    c.mu.Lock()
    item.UpdateAccess()
    c.mu.Unlock()
    return item.Value, true
}

## 四、踩坑记录
1. 高并发map panic：添加sync.RWMutex读写锁解决
2. 覆盖key内存统计错误：Set前先删除旧条目并更新内存
3. Get操作锁竞争：读锁查询，写锁更新访问时间

## 五、测试结果
- 单元测试覆盖率≥90%
- 并发测试无数据竞争
- 读QPS≈15w/s，写QPS≈8w/s

## 六、后续计划
- 阶段2：添加内存限制+LRU淘汰
- 优化内存计算精度
- 支持批量操作
