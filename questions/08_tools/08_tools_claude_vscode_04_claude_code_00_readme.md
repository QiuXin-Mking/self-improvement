# q
如何克隆 everything-claude-code 仓库并复制其配置文件到 Claude 的配置目录？
# a
使用以下命令：
```sh
git clone https://github.com/affaan-m/everything-claude-code.git
cp everything-claude-code/agents/*.md ~/.claude/agents/
cp everything-claude-code/rules/*.md ~/.claude/rules/
cp everything-claude-code/commands/*.md ~/.claude/commands/
cp -r everything-claude-code/skills/* ~/.claude/skills/
```

# q
如何创建 appuser 用户并赋予其对 `/app/rocksdb-9.7.3` 目录的完整权限？
# a
先创建用户，然后改变目录所有权，最后切换用户或以其身份执行命令：
```sh
useradd -m appuser
chown -R appuser:appuser /app/rocksdb-9.7.3
su - appuser
```

# q
如何以 appuser 身份运行 claude 并跳过权限检查、执行 `/init` 命令？
# a
使用 `su - appuser -c` 执行带选项的 claude 命令：
```sh
su - appuser -c 'claude --allow-dangerously-skip-permissions -p "/init"'
```

