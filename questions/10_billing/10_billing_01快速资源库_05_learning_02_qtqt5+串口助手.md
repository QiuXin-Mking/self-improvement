# q
如何使用清华大学镜像源安装 PyQt5 及串口助手所需的 Python 包？
# a
```bash
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyqt5
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyqt5-tools
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyserial
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ matplotlib
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyinstaller
```
其他可用国内源：阿里云 `https://mirrors.aliyun.com/pypi/simple/`，豆瓣 `http://pypi.douban.com/simple/`。

# q
在 PyCharm 中如何配置 Qt Designer 作为外部工具？
# a
路径：`File -> Settings -> Tools -> External Tools`，添加工具：
- Program: `C:\ProgramData\Anaconda3\Library\bin\designer.exe`
- Arguments: 留空或按需设置
- Working directory: `$FileDir$`

# q
在 PyCharm 中如何配置 pyuic5 工具将 `.ui` 文件转换为 `.py` 文件？
# a
外部工具配置：
- Program: Python 解释器路径（例如 `E:\...\venv\Scripts\python`，可通过 `where python` 查找）
- Arguments: `-m PyQt5.uic.pyuic  $FileName$ -o $FileNameWithoutExtension$.py`
- Working directory: `$FileDir$`

# q
在 Windows 命令行中如何查看当前 Python 解释器的完整路径？
# a
```cmd
where python
```
执行后会列出所有 Python 可执行文件的路径，例如：
```
E:\016 淘宝-春哥-做单\2022\010_Learn_Project\004_qt5_串口\py\pythonProject\venv\Scripts\python
```

