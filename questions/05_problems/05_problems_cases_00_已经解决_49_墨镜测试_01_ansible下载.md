# q
`ansible` 命令找不到的典型根因是什么？
# a
通过 `pip3 install --user ansible` 安装后，`ansible` 实际安装在 `~/.local/bin/` 下，该目录默认不在 `PATH` 中。
解决方案：
1. 确认安装位置：`ls ~/.local/bin/ansible`
2. 直接使用完整路径运行：`~/.local/bin/ansible --version`
3. 或将目录加入 `PATH`：`export PATH=$PATH:~/.local/bin`（可写入 `~/.bashrc` 永久生效）

