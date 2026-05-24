# q
如何在Windows中激活Python虚拟环境？
# a
在Windows命令行中，执行以下命令激活指定虚拟环境（例如venv_pycharm）：
```sh
.\venv_pycharm\Scripts\activate
```

# q
如何在Windows中临时使用pip升级自身而不依赖环境变量中的pip命令？
# a
通过调用`python.exe`直接运行pip模块：
```sh
python.exe -m pip install --upgrade pip
```

# q
如何在本地存在所有wheel包的情况下离线安装git-review？
# a
使用`--no-index`和`--find-links`参数，从本地目录安装：
```sh
pip install --no-index --find-links=./ git-review
```

# q
git-review已安装但命令行仍无法使用，可能的原因是什么？
# a
系统环境变量没有配置。git-review的可执行文件所在目录（通常是Python的Scripts目录）未被添加到系统`PATH`中，导致命令行无法找到该命令。

