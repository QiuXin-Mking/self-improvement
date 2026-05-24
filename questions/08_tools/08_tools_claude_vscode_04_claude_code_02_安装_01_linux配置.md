# q
如何在Linux上安装Claude Code？
# a
执行 `npm install -g pumpkinai-config` 安装配置工具，然后运行 `claude-config` 并跟随提示操作。启动前需设置环境变量 `export ANTHROPIC_BASE_URL=https://turingai.plus`，最后通过 `claude` 命令启动。

# q
使用Claude Code需要设置哪些关键环境变量？
# a
关键环境变量包括：
- `ANTHROPIC_AUTH_TOKEN` 和 `ANTHROPIC_API_KEY`（均为API密钥，示例值 `sk-tBOv77W3LSK592bUKxJ6hHHeqTMnuFvA1kBKdOXyI96v4ZP6`）
- `ANTHROPIC_BASE_URL`（API基础地址，示例 `https://turingai.plus`）
- `CLAUDE_CODE_MAX_OUTPUT_TOKENS`（最大输出token数，示例 `32000`）
- `ANTHROPIC_SMALL_FAST_MODEL`（快速模型名，示例 `claude-sonnet-4-5-20250929`）

# q
如何用curl测试Claude Code的API连接？
# a
使用以下命令模拟请求并查看响应：
```
curl -v -X POST https://turingai.plus/v1/messages \
  -H "Content-Type: application/json" \
  -H "x-api-key: sk-tBOv77W3LSK592bUKxJ6hHHeqTMnuFvA1kBKdOXyI96v4ZP6" \
  -d '{"model":"claude-3-5-haiku-20241022","max_tokens":10,"messages":[{"role":"user","content":"test"}]}'
```

# q
如何卸载Claude Code？
# a
执行 `npm uninstall -g @anthropic-ai/claude-code` 卸载Claude Code主包，如需重新配置可再运行 `npm install -g pumpkinai-config`。

# q
`/root/.bashrc` 文件有什么作用？
# a
`.bashrc` 是 Bash shell 的运行时配置文件，当启动一个新的非登录 Shell 时会自动执行其中的命令，常用于设置别名、函数和环境变量等。

