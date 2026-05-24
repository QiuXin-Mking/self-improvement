# q
编译 Lustre 服务器端 RPM 包时，`./configure` 需要哪些关键参数及重要注意事项？
# a
编译服务器端时使用以下命令：
```shell
./configure --enable-server --enable-quota \
  --with-linux=/home/chenlou/er-linux --with-zfs=no
```
然后执行 `make rpms`。  
**特别注意**：不要在 `./configure` 中添加 `--disable-client`，否则会意外移除 `fld.ko` 模块，导致服务器功能异常。

# q
编译 Lustre 客户端 RPM 包时，`./configure` 需要哪些关键参数？
# a
编译客户端时需要使用与服务器相反的选项：
```shell
./configure --disable-server --enable-client --enable-quota \
  --with-linux=/home/chenlou/er-linux --with-zfs=no
```
然后执行 `make rpms`。客户端与服务器使用同一份源码，通过 `--disable-server` 和 `--enable-client` 进行区分。

# q
安装内核 headers RPM 包时若遇到 “obsoleted” 依赖冲突，应如何强制安装？
# a
使用 `rpm` 的 `--force` 和 `--nodeps` 选项忽略依赖与冲突：
```shell
rpm -ivh --force --nodeps kernel-headers-4.18.0-3.2.1.er1.x86_64.rpm
```
或将整个目录下的所有相关 RPM 包一同强制安装：
```shell
rpm -ivh --force --nodeps ./*.rpm
```

# q
直接通过 `make install` 安装 Lustre 后，需要执行哪些系统命令以确保内核模块和共享库正确加载？
# a
安装后必须执行以下命令：
```shell
depmod -a   # 更新内核模块依赖关系
ldconfig -v # 刷新共享库缓存，使程序能访问新的 Lustre 库
```
`depmod -a` 会重建 `/lib/modules/$(uname -r)/modules.dep` 等文件，`ldconfig -v` 则更新 `/etc/ld.so.cache`，否则动态加载模块或运行客户端工具时可能出现符号缺失错误。

