# q
如何将多个文件打包并压缩为.tar.gz格式？
# a
使用`tar -czf`命令，后面指定输出文件名和要打包的文件。例如：
```bash
tar -czf upgrade.tar.gz mdbs-1.1.18-Linux.rpm mdbs-env-1.0.23-Linux.rpm
```

# q
如何解压.tar.gz文件？
# a
使用`tar -zxvf`命令，后面指定要解压的.tar.gz文件路径。例如：
```bash
tar -zxvf /usr/local/test.tar.gz
```
对于`.tgz`文件同样适用：
```bash
tar -xzvf file.tgz
```

# q
tar如何自动识别压缩格式进行解压？
# a
从1.15版本开始，tar可以自动识别压缩格式，因此无需指定压缩类型参数（如`-z`），直接使用`tar -xvf`即可正确解压。例如：
```bash
tar -xvf filename.tar.gz
```

