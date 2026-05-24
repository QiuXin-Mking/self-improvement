# q
在树莓派上使用 pip 安装 ultralytics 时，为什么会出现 `Could not find a version that satisfies the requirement numpy>=1.22.2` 的错误？
# a
因为树莓派的默认 pip 源中可用的 numpy 最高版本仅为 1.21.6，而 ultralytics 需要 numpy>=1.22.2，导致无法找到满足版本要求的包。

# q
如何解决树莓派安装 ultralytics 时因 numpy 版本不足导致的依赖错误？
# a
首先升级 pip 并切换到国内镜像源（如清华源），命令为：
```bash
pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple
```
然后再执行 `pip install ultralytics` 即可正常安装。

