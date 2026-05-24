# q
安装Node.js时为什么建议使用默认路径不修改？
# a
建议默认安装在C盘，不要修改路径，以防后续出现报错。

# q
在Windows PowerShell中运行npm命令时出现“无法加载文件 ... 禁止运行脚本”的错误如何解决？
# a
临时绕过执行策略：
```powershell
powershell -ExecutionPolicy Bypass -Command "npm -v"
```
永久修改当前用户执行策略：
```powershell
Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned
```

# q
如何安装并配置Claude Code（基于pumpkinai-config）？
# a
安装配置工具并初始化：
```sh
npm install -g pumpkinai-config
claude-config
```
若`claude-config`报错，可先手动安装依赖再执行：
```sh
npm install -g @anthropic-ai/claude-code
claude-config
```
根据需要设置API相关环境变量，例如：
```powershell
$env:ANTHROPIC_BASE_URL = "https://turingai.plus"
```

