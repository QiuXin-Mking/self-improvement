# q
使用阿里云百炼模型时，Moltbot 的 API 调用 base_url 是什么？
# a
```text
https://dashscope.aliyuncs.com/compatible-mode/v1
```

# q
Moltbot 的配置文件通常存放在哪个路径？
# a
`~/.moltbot/moltbot.json` 或 `~/.clawdbot/clawdbot.json`

# q
在 Moltbot 配置中添加完百炼模型后，如何让配置生效？
# a
停止服务后重新启动，或直接执行：
```bash
clawdbot gateway restart
```

# q
如何测试所配置的百炼模型是否正常可用？
# a
- 检查模型连通性：`clawdbot models status --probe`
- 发送一条测试消息：`clawdbot agent --agent main --message "介绍下阿里云百炼"`

# q
使用百炼模型时，API 密钥应设置在哪个环境变量中？
# a
`DASHCOSPE_API_KEY`，通常写入 `~/.bash_profile` 文件中（bash 环境）。

