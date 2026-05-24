# q
什么是Python pip包镜像源
# a
Python pip包镜像源是PyPI（Python Package Index）的镜像站点，用于加速Python包的下载和安装，特别是在网络访问官方源缓慢或不稳定的情况下。常见的镜像源有清华源（`https://pypi.tuna.tsinghua.edu.cn/simple`）等。

# q
如何使用清华源临时安装Python包
# a
使用`pip install`命令并指定`-i`参数，后跟清华源的URL：
```
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple 包名
```
例如临时使用清华源安装某个包而不修改全局配置。

# q
清华源的完整URL是什么
# a
`https://pypi.tuna.tsinghua.edu.cn/simple`

