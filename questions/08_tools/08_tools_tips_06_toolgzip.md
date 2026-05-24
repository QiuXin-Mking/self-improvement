# q
如何安装 gzip？
# a
在 RHEL/CentOS 系统上使用 `yum install gzip -y` 命令安装。

# q
如何使用 gzip 压缩单个文件？
# a
使用 `gzip filename.txt` 命令，压缩后原文件会被替换为 `filename.txt.gz`。

# q
如何将 gzip 压缩结果输出到指定路径而不删除原文件？
# a
使用 `gzip -c filename.txt > /tmp/filename.txt.gz`，通过 `-c` 将压缩数据写入标准输出，重定向到目标文件。

# q
如何解压 .gz 文件？
# a
使用 `gzip -d filename.txt.gz` 命令解压。若需将解压内容输出到指定位置而不影响原压缩文件，可使用 `gzip -d -c /tmp/filename.txt.gz > /tmp/d6z/filename.txt`。

