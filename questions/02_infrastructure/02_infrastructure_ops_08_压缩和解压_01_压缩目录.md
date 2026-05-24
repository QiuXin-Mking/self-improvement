# q
tar 命令中的 -c、-z、-v、-f 选项分别起什么作用？
# a
- `-c`：创建新的归档文件  
- `-z`：通过 gzip 进行压缩或解压  
- `-v`：显示处理过程中的文件列表（verbose 模式）  
- `-f`：指定归档文件名，必须紧跟文件名使用（例如 `-czvf archive.tar.gz`）

# q
如何使用 tar 将一个目录压缩为 `.tar.gz` 格式的归档文件？
# a
使用 `tar -czvf <归档文件名> <目录路径>` 命令。例如：
```bash
tar -czvf pak.tar.gz ./inotify_downloads/
```
这会将 `./inotify_downloads/` 目录及其内容压缩打包为 `pak.tar.gz`。

# q
如何通过组合 pip download 和 tar 制作一个离线安装的 Python 包归档，并安装它？
# a
（1）先用 `pip download` 将包及其依赖下载到指定目录：
```bash
pip download inotify -d inotify_downloads --index-url https://mirrors.aliyun.com/pypi/simple/
```
（2）用 `tar` 将下载目录压缩打包：
```bash
tar -czvf pak.tar.gz ./inotify_downloads/
```
（3）最后使用 `pip install` 直接安装打包文件（假设内部结构可被 pip 识别）：
```bash
pip install pak.tar.gz
```

