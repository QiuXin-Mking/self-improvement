# q
在ESP-IDF环境中遇到“pip is not valid (ERROR_INVALID_PIP)”错误如何解决？
# a
运行以下命令修复pip：
```
C:\Espressif\tools\idf-python\3.11.2\python.exe -m ensurepip
```

# q
ESP-IDF编译时ninja报错并出现一堆错误码，最简单的处理方式是什么？
# a
直接删除所有 build 文件（或 build 目录），然后重新编译。

