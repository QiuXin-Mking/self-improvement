# q
在多个模块中对同名 logger 重复调用 `addHandler` 会导致什么问题？
# a
会导致日志消息被重复打印。因为 `logging.getLogger('st_tools')` 返回的是全局唯一的 logger 实例，每次 `addHandler` 都会向该 logger 追加新的 handler，同一日志条目会被多个 handler 重复输出。

# q
上述代码中（`utils.py` 与 `api.py` 各自定义 `st_log` 并实例化）为什么会导致日志重复？
# a
两个模块都通过 `logging.getLogger('st_tools')` 获取同一个 logger，然后在 `set_logger()` 中分别添加了 `StreamHandler` 和 `FileHandler`。两个模块都被导入时，logger 将累积 4 个 handler，每次 `logger.info(...)` 会被执行 4 遍，造成重复打印。

# q
如何防止因多次实例化或多次导入而重复添加 handler？
# a
在添加 handler 前检查 logger 是否已有 handler，例如：
```python
if not self.logger.handlers:
    self.logger.addHandler(sh)
    self.logger.addHandler(fh)
```
或者将 logger 配置提取到一次性初始化逻辑中，确保全局只添加一次 handler。

# q
`logger.handlers` 在避免重复添加 handler 中起什么作用？
# a
`logger.handlers` 是一个列表，存放当前 logger 已绑定的所有 handler。通过判断该列表是否为空，可以确定是否已经添加过 handler，从而避免重复添加。

