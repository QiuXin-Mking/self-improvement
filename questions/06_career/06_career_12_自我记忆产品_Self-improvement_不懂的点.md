# q
在 Debian/Ubuntu 容器中运行 `apt install python` 为什么提示包不存在？
# a
因为官方仓库中没有独立的 `python` 包，该名称被 `python-is-python3` 和 `2to3` 包替代。现代 Debian/Ubuntu 默认将所有 Python 相关命令指向 Python 3，所以需要安装 `python-is-python3` 来创建 `/usr/bin/python` 符号链接。

# q
在 Debian/Ubuntu 中安装 Python 3 的完整步骤是什么？
# a
先更新包列表，然后安装必要的 Python 3 组件：
```sh
apt update
apt install python3 python3-pip python3-venv
```
如果需要兼容旧的 `python` 命令，额外安装 `python-is-python3`：
```sh
apt install python-is-python3
```
安装后可通过 `python --version` 或 `python3 --version` 验证。

# q
`python-is-python3` 包的作用是什么？
# a
它会在系统中创建一个 `/usr/bin/python` 的符号链接，指向 Python 3 解释器，使得直接执行 `python` 命令时实际调用的是 Python 3，从而兼容那些仍使用 `python` 命令的旧脚本或工具。

