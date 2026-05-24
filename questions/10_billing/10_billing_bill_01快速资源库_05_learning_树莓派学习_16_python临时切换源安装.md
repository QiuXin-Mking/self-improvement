# q
如何使用 pip 临时切换到清华大学 PyPI 镜像源来安装 matplotlib？
# a
使用 `-i` 参数指定镜像地址：
```
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple matplotlib
```

# q
上述 pip 命令中 `-i` 参数的作用是什么？
# a
`-i` 或 `--index-url` 用于指定 Python 包的索引源 URL，此次安装仅本次使用该源，不会修改全局配置。

