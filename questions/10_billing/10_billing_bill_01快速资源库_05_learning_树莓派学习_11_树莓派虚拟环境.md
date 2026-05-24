# q
在树莓派上如何使用 `venv` 创建 Python 虚拟环境？
# a
使用命令：
```bash
python3 -m venv myenv
```
`myenv` 是虚拟环境的目录名，可根据需要自定义。

# q
如何激活已创建的 Python 虚拟环境？
# a
在终端中运行：
```bash
source myenv/bin/activate
```
激活后命令行前缀会出现 `(myenv)` 表示已进入该环境。

# q
如何查找树莓派上所有已创建的虚拟环境目录？
# a
使用 `find` 命令在用户主目录下搜索以 `env` 结尾的目录：
```bash
find ~ -type d -name "*env"
```
示例输出：
```
/home/pi/tflite1/tflite1-env
/home/pi/myenv
```

# q
如何为 pip 配置清华大学镜像源？
# a
执行以下命令：
```bash
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
```
该配置会写入 `/home/pi/.config/pip/pip.conf`，使后续 `pip install` 默认从清华源下载包。

