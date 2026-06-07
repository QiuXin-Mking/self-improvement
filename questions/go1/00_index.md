# Go 训练题目索引入口

> C 高级工程师 → Go 高级开发工程师，按场景和知识点整理的专项 quiz 训练。

## 学习顺序

按依赖关系排列，建议从上往下刷：

| # | 文件 | 专题 | 题数 | C→Go 重要度 |
|---|------|------|------|------------|
| 1 | [02_slice.md](02_slice.md) | Slice 切片 | 10 | ★★★★★ Go 独有，C 最不适应的 |
| 2 | [03_map.md](03_map.md) | Map 映射 | 10 | ★★★★★ 内置 hash，C 没有 |
| 3 | [04_defer.md](04_defer.md) | Defer 延迟执行 | 10 | ★★★★★ Go 独有内存管理机制 |
| 4 | [05_goroutine.md](05_goroutine.md) | Goroutine 协程 | 9 | ★★★★★ Go 的并发基石 |
| 5 | [06_channel.md](06_channel.md) | Channel 通道 | 12 | ★★★★★ Go 并发通信核心 |
| 6 | [07_interface.md](07_interface.md) | Interface 接口 | 10 | ★★★★☆ 和 C++ 虚函数完全不同 |
| 7 | [08_sync.md](08_sync.md) | 同步原语 | 10 | ★★★★☆ Mutex/WaitGroup/Once |
| 8 | [09_context.md](09_context.md) | Context 上下文 | 10 | ★★★★☆ 超时/取消/传值 |
| 9 | [10_error.md](10_error.md) | Error 错误处理 | 10 | ★★★★☆ 没有 exception 的 Go 哲学 |
| 10 | [11_pointer_escape.md](11_pointer_escape.md) | 指针与逃逸分析 | 10 | ★★★★☆ GC 三色标记/栈堆 |
| 11 | [12_struct.md](12_struct.md) | 结构体与方法 | 10 | ★★★☆☆ 嵌入/接收者/tag |
| 12 | [13_json.md](13_json.md) | JSON 序列化 | 9 | ★★★☆☆ Marshal/Unmarshal/Decoder |
| 13 | [14_net_http.md](14_net_http.md) | HTTP 服务 | 9 | ★★★☆☆ 标准库 HTTP/middleware |
| 14 | [15_testing.md](15_testing.md) | 测试 | 8 | ★★★☆☆ 单元/基准/Table-driven |
| 15 | [16_generics.md](16_generics.md) | 泛型 | 7 | ★★★☆☆ Go 1.18+ 新特性 |
| 16 | [17_patterns.md](17_patterns.md) | 常见模式 | 9 | ★★★☆☆ Option/Worker/Pipeline/项目结构 |

## 使用方式

每个文件里是 `# q` / `# a` 交替的题目和答案：

```
# q
问题描述

# a
答案和详细解析
```

建议：**先看 q，自己默答一遍，再看 a 对照**。每题不要超过 2 分钟。

## 刷题路线

- **第 1 天**：Slice + Map + Defer（Go 三大核心差异）
- **第 2 天**：Goroutine + Channel（Go 并发灵魂）
- **第 3 天**：Interface + Sync（面向接口 + 同步）
- **第 4 天**：Context + Error（工程化必备）
- **第 5 天**：Pointer/Escape + Struct（内存模型 + 面向对象）
- **第 6 天**：JSON + HTTP + Testing（标准库实战）
- **第 7 天**：Generics + Patterns（进阶 + 最佳实践）
