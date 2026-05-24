# q
在Windows上如何列出已安装的Node.js版本并切换到指定版本？
# a
使用nvm-windows工具：
- 列出已安装版本：```nvm list```
- 切换到指定版本（例如24.0.1）：```nvm use 24.0.1```

# q
Claude Code在Windows下推荐通过哪个配置文件设置API密钥和模型？该文件的路径是什么？
# a
推荐使用```~/.claude/settings.json```配置文件，而不是设置系统环境变量。在Windows上，```~```对应```C:\Users\Administrator```（或当前用户目录）。配置格式示例：
```json
{
    "env": {
        "ANTHROPIC_AUTH_TOKEN": "你的key",
        "ANTHROPIC_BASE_URL": "https://turingai.plus",
        "ANTHROPIC_MODEL": "claude-sonnet-4-5-20250929"
    }
}
```

# q
在PowerShell中如何查看当前会话的所有Anthropic和Claude Code相关环境变量？
# a
用管理员权限打开PowerShell，执行以下命令：
```powershell
Write-Host "ANTHROPIC_AUTH_TOKEN: $env:ANTHROPIC_AUTH_TOKEN"
Write-Host "ANTHROPIC_API_KEY: $env:ANTHROPIC_API_KEY"
Write-Host "ANTHROPIC_BASE_URL: $env:ANTHROPIC_BASE_URL"
Write-Host "CLAUDE_CODE_MAX_OUTPUT_TOKENS: $env:CLAUDE_CODE_MAX_OUTPUT_TOKENS"

Write-Host ""
Write-Host "=== 查看所有相关环境变量 ==="
Get-ChildItem Env: | Where-Object {$_.Name -like "*ANTHROPIC*" -or $_.Name -like "*CLAUDE_CODE*"}
```
输出示例会显示```ANTHROPIC_AUTH_TOKEN```、```ANTHROPIC_API_KEY```、```ANTHROPIC_BASE_URL```等已设置的环境变量及其值。

