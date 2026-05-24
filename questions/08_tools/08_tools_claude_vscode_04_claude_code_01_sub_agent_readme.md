# q
Sub-agents适合什么场景？
# a
需要并行执行多个复杂任务的时候。

# q
启动Claude Code主agent并跳过权限检查的命令是什么？
# a
```
claude --dangerously-skip-permissons --versbose
```

# q
如何使用子代理（sub-agent）？
# a
先用 `/agent` 创建一个执行子任务的超级专家，然后用一个父亲agent去调用他。

