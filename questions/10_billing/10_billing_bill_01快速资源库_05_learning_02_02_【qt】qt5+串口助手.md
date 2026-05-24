# q
如何通过清华镜像源安装 PyQt5 及其相关工具包？
# a
使用 pip 命令并通过 -i 参数指定清华镜像源：
```bash
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyqt5
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyqt5-tools
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyserial
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ matplotlib
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ pyinstaller
pip install -i https://pypi.tuna.tsinghua.edu.cn/simple/ py2exe
```

# q
在 PyCharm 中如何配置外部工具将 .ui 文件转换为 .py 文件？
# a
在 PyCharm 的 File -> Settings -> Tools -> External Tools 中添加一个工具，配置如下：
- Program: `E:\016 淘宝-春哥-做单\2022\010_Learn_Project\004_qt5_串口\py\pythonProject\venv\Scripts\python`
- Arguments: `-m PyQt5.uic.pyuic  $FileName$ -o $FileNameWithoutExtension$.py`
- Working directory: `$FileDir$`
同时可能需要设置环境变量 QT_QPA_PLATFORM_PLUGIN_PATH 为 Qt 的 platforms 目录，例如 `C:\ProgramData\Anaconda3\pkgs\qt-5.9.7-vc14h73c81de_0\Library\plugins\platforms`。

# q
如何激活 Windows 上特定项目的 Python 虚拟环境？
# a
在命令行中进入虚拟环境的 Scripts 目录并执行 `activate`，例如：
```cmd
E:\016 淘宝-春哥-做单\2022\010_Learn_Project\004_qt5_串口\py\pythonProject\venv\Scripts>activate
```

# q
在 Windows 中如何查找 Python 命令的实际路径？
# a
在命令提示符中使用 `where python` 命令，它会返回所有可执行 python 的路径，例如：
```
where python
```
可能输出：
```
E:\016 淘宝-春哥-做单\2022\010_Learn_Project\004_qt5_串口\py\pythonProject\venv\Scripts\python
```

