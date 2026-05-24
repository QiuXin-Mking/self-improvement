# q
如何使用Docker在Windows上启动一个配置了Claude Code的Node.js开发容器？
# a
使用`docker run`命令，映射项目目录（`C:\Users\Administrator\Desktop\vue3-demo`到`/app`）和备份目录（`D:\qiuxin_aliyun_back`到`/qiuxin_aliyun_back`），并设置工作目录为`/app`，暴露端口5173。示例命令：
```
docker run -itd \
  --name vue3-dev \
  -p 5173:5173 \
  -v "C:\Users\Administrator\Desktop\vue3-demo:/app" \
  -v "D:\qiuxin_aliyun_back:/qiuxin_aliyun_back" \
  -w /app \
  node:20
```

# q
为了让Claude Code在Docker容器内正常工作，如何预配置环境变量和设置文件？
# a
需要修改Dockerfile，实现两项配置：
1. 在`/root/.bashrc`中追加环境变量：
   ```
   export ANTHROPIC_API_KEY="sk-..."
   export ANTHROPIC_BASE_URL="https://turingai.plus/"
   export CLAUDE_CODE_DISABLE_EXPERIMENTAL_BETAS="1"
   export ANTHROPIC_SMALL_FAST_MODEL="claude-3-5-haiku-20241022"
   export CLAUDE_CODE_MAX_OUTPUT_TOKENS=32000
   ```
2. 创建`/root/.claude/settings.json`文件，内容如下：
   ```
   {
     "env": {
       "ANTHROPIC_AUTH_TOKEN": "sk-...",
       "ANTHROPIC_BASE_URL": "https://turingai.plus",
       "ANTHROPIC_SMALL_FAST_MODEL": "claude-3-5-haiku-20241022"
     }
   }
   ```

# q
如何通过一键脚本快速安装并配置Claude Code？
# a
执行以下命令即可完成安装和交互式配置：
```
npm install -g pumpkinai-config
claude-config
```
若`claude-config`运行报错，可先手动安装`@anthropic-ai/claude-code`（`npm install -g @anthropic-ai/claude-code`），再重新执行`claude-config`。

# q
手动安装Claude Code并配置其环境变量的完整步骤是什么？
# a
1. 设置npm淘宝镜像源加速：
   ```
   npm config set registry https://registry.npmmirror.com
   ```
2. 全局安装Claude Code：
   ```
   npm install -g @anthropic-ai/claude-code
   ```
3. 验证安装：
   ```
   claude --version
   ```
4. 在shell配置文件（如`~/.zshrc`）中追加所需环境变量（请将`sk-...`替换为实际密钥）：
   ```
   export ANTHROPIC_AUTH_TOKEN="sk-..."
   export ANTHROPIC_API_KEY="sk-..."
   export ANTHROPIC_BASE_URL="https://turingai.plus"
   export ANTHROPIC_SMALL_FAST_MODEL="claude-3-5-haiku-20241022"
   ```
5. 使配置生效：
   ```
   source ~/.zshrc
   ```

