# q
如何通过指定国内镜像源在线安装 open-webui？
# a
使用以下命令指定清华源安装：
```bash
python3.11 -m pip install open-webui -i https://pypi.tuna.tsinghua.edu.cn/simple
```

# q
离线安装 open-webui 需要执行哪两步命令？
# a
先下载离线包到指定目录：
```bash
pip download open-webui -d ./offline-packages -i https://pypi.tuna.tsinghua.edu.cn/simple
```
再通过本地包安装：
```bash
python3.11 -m pip install --no-index --find-links ./offline-packages open-webui
```

# q
安装 open-webui 时遇到 “Cannot uninstall blinker 1.4” 错误如何解决？
# a
该错误由系统安装的 blinker 包冲突引起，需先移除系统包：
```bash
apt-get remove python3-blinker
```

# q
安装 open-webui 时提示 pypika 构建 wheel 失败，应该如何处理？
# a
先升级构建工具链并强制重装 pip，再安装 pypika：
```bash
python3.11 -m pip install --upgrade --force-reinstall pip setuptools wheel build cython
python3.11 -m pip install pypika
```

