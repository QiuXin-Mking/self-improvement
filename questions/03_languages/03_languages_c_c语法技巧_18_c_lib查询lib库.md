# q
如何在Linux系统中查询C库（如libudev）的安装路径？
# a
使用 `whereis` 命令，例如 `whereis libudev`，它会返回库文件（如 `/usr/lib64/libudev.so`）和头文件（如 `/usr/include/libudev.h`）的路径。

# q
找到C库的头文件后，如何查看其内容？
# a
切换到头文件所在目录（如 `cd /usr/include/`），然后用文本编辑器打开，例如 `vi libudev.h`。

