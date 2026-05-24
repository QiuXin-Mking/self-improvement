# q
在树莓派上使用pip安装ultralytics时，提示“Could not find a version that satisfies the requirement numpy>=1.22.2”应如何解决？
# a
升级pip并更换为国内镜像源（如清华源）：
```bash
pip install --upgrade pip -i https://pypi.tuna.tsinghua.edu.cn/simple
```
此问题通常是因为树莓派自带的pip版本过低，无法从默认源正确解析高版本numpy的依赖关系。

