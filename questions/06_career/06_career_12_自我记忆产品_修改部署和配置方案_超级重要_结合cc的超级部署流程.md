# q
Docker 运行 vue3-dev 容器时，如何正确挂载主机目录并设置工作目录？
# a
执行以下命令，其中映射两个路径：
```bash
docker run -itd \
  --name vue3-dev \
  -p 5173:5173 \
  -v "C:\Users\Administrator\Desktop\vue3-demo:/app" \
  -v "D:\qiuxin_aliyun_back:/qiuxin_aliyun_back" \
  -w /app \
  node:20
```
端口 5173 映射到容器内 5173，工作目录设为 `/app`。

# q
在容器内为 Claude Code 配置环境变量，需要在 `/root/.bashrc` 中追加哪些内容？
# a
追加以下环境变量（具体 key 需替换为实际值）：
```bash
export ANTHROPIC_API_KEY="sk-tBOv77W3LSK592bUKxJ6hHHeqTMnuFvA1kBKdOXyI96v4ZP6"
export ANTHROPIC_BASE_URL="https://turingai.plus/"
export CLAUDE_CODE_DISABLE_EXPERIMENTAL_BETAS="1"
export ANTHROPIC_SMALL_FAST_MODEL="claude-3-5-haiku-20241022"
export CLAUDE_CODE_MAX_OUTPUT_TOKENS=32000
```

# q
容器内用于 Claude Code 的 `/root/.claude/settings.json` 应写入什么配置？
# a
```json
{
  "env": {
    "ANTHROPIC_AUTH_TOKEN": "sk-tBOv77W3LSK592bUKxJ6hHHeqTMnuFvA1kBKdOXyI96v4ZP6",
    "ANTHROPIC_BASE_URL": "https://turingai.plus",
    "ANTHROPIC_SMALL_FAST_MODEL": "claude-3-5-haiku-20241022"
  }
}
```
此文件设置 Claude Code 运行时的 API 身份、基础 URL 和默认小快速模型。

