# GoCacheWisdom - 从零自研本地缓存库

基于 Go 从零实现的高可扩展本地缓存库，支持分步迭代式功能完善（内存限制、LRU淘汰、过期删除、时间轮优化），同时沉淀开发经验与最佳实践。

## 核心特性
- 线程安全：基于sync.RWMutex实现高并发安全访问
- 基础缓存：支持Set/Get/Delete/Batch核心操作
- 内存控制：可设置最大内存上限，内置LRU淘汰策略
- 过期管理：分步实现惰性删除→定时扫描→时间轮优化
- 高可扩展：接口化设计，支持自定义扩展
- 轻量无依赖：从零自研，无第三方依赖

## 项目结构
GoCacheWisdom/
├── cache/          # 核心缓存库（可导出）
│   ├── core/       # 核心逻辑
│   ├── policy/     # 淘汰策略
│   ├── expire/     # 过期处理
│   └── cache.go    # 对外导出入口
├── examples/       # 使用示例
├── docs/           # 开发文档
├── tests/          # 测试用例
├── internal/       # 内部工具
├── go.mod          # 依赖管理
├── CHANGELOG.md    # 版本记录
└── README.md       # 项目说明

## 快速开始
### 1. 安装
go get github.com/GoServePro/GoDevExperience/GoCacheWisdom

### 2. 基础使用
package main

import (
    "fmt"
    "github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache"
)

func main() {
    // 创建缓存实例
    c := cache.New()

    // 存储数据
    c.Set("user:1001", "张三")

    // 获取数据
    val, ok := c.Get("user:1001")
    if ok {
        fmt.Printf("命中缓存：%v\n", val)
    }
}

## 迭代路线图
| 阶段 | 核心功能 | 状态  |
|------|----------|-----|
| 阶段1 | 基础缓存（Set/Get/Delete） | 已完成 |
| 阶段2 | 内存限制 + LRU淘汰 | 开发中 |
| 阶段3 | 过期key删除 | 规划中 |
| 阶段4 | 时间轮优化 | 规划中 |

## 许可证
本项目基于MIT许可证开源。

## 联系作者
GitHub：@GoServePro
