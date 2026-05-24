# q
安装ESP-IDF后出现“pip无效”错误，如何修复？
# a
运行以下命令重新安装pip：
```
C:\Espressif\tools\idf-python\3.11.2\python.exe -m ensurepip
```

# q
编译ESP32项目时ninja报出大量错误码，应如何快速恢复？
# a
直接删除项目目录下的所有`build`文件夹（清理构建输出），然后重新编译。

