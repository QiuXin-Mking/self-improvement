# q
FUSE 文件系统的核心原理是什么？
# a
FUSE (Filesystem in Userspace) 允许在用户空间实现文件系统，核心是通过注册一个 `fuse_operations` 结构体，其中包含一组回调函数（如 `getattr`、`readdir`、`open`、`read`、`write` 等），内核 VFS 层将文件操作请求转发给用户态守护进程，由这些回调函数处理并返回结果。

# q
FUSE 开发需要安装哪些依赖？如何编译一个 FUSE 程序？
# a
需要安装 `fuse3-devel` 开发包（CentOS/RHEL 下执行 `sudo yum install fuse3-devel`）。编译使用 pkg-config 获取头文件和库路径：
```bash
gcc -Wall `pkg-config fuse3 --cflags --libs` hello_fuse.c -o hello_fuse
```
若遇到 "Package fuse3 not found" 错误，需确保 `PKG_CONFIG_PATH` 包含 `fuse3.pc` 所在目录（如 `/usr/lib/pkgconfig` 或 `/usr/lib64/pkgconfig`）。

# q
如何运行并验证一个简单的 FUSE 文件系统？
# a
1. 创建挂载点目录：`mkdir ~/fuse_mount`
2. 运行 FUSE 程序：`./hello_fuse ~/fuse_mount`
3. 测试文件系统：`ls ~/fuse_mount` 可看到文件 `hello`，执行 `cat ~/fuse_mount/hello` 输出 `Hello World!`

