# q
在 Debian/Ubuntu 上执行 `apt install python` 因“Package 'python' has no installation candidate”失败，应如何解决？
# a
首先运行 `apt update` 更新软件包索引，然后安装 `python-is-python3` 来提供 `python` 命令，再安装 `python3 python3-pip python3-venv` 以获得完整的 Python 3 环境。
```sh
apt update
apt install python-is-python3
python --version
apt install python3 python3-pip python3-venv
```

