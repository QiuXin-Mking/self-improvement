# q
编译 Lustre 3.3.1 版本需要准备哪些内核相关的 RPM 包？
# a
需要准备以下四个包：
- kernel-4.18.0-3.3.1.er1.x86_64.rpm
- kernel-devel-4.18.0-3.3.1.er1.x86_64.rpm
- kernel-headers-4.18.0-3.2.1.er1.x86_64.rpm
- kernel-debuginfo-4.18.0-3.3.1.er1.x86_64.rpm

# q
如何强制安装并切换到指定内核版本？
# a
1. 使用 `rpm -ivh --nodeps --force ./kernel-*.rpm` 无视依赖强制安装。
2. 根据启动模式更新 grub 配置：
   - BIOS 模式：`grub2-mkconfig -o /boot/grub2/grub.cfg`
   - UEFI 模式：`grub2-mkconfig -o /boot/efi/EFI/centos/grub.cfg`
3. 重启系统后执行 `uname -r` 确认内核版本已切换。

# q
编译 Lustre 服务端时为什么严禁使用 `--disable-client` 选项？
# a
在服务端编译的 `./configure` 命令中若加入 `--disable-client`，会导致 `fld.ko` 模块丢失，这是一个严重隐患，因此应避免使用该选项。

# q
编译 Lustre 时 `--with-linux` 参数应如何指定？
# a
`--with-linux` 应指向内核源码或 kernel-devel 安装后的路径，通常为 `/usr/src/kernels/<kernel-version>`。如果使用完整的内核源码包，则直接指向源码目录即可，例如 `/root/er-linux/`。

# q
编译 Lustre 客户端和服务端的典型配置步骤是什么？
# a
```bash
sh autogen.sh

# 编译客户端（禁止 server 功能，启用 client）
./configure --disable-server --enable-client --enable-quota \
            --with-linux=/usr/src/kernels/4.18.0-3.3.1.er1.x86_64 --with-zfs=no
make rpms

# 编译服务端（不要加 --disable-client）
./configure --enable-server --enable-quota \
            --with-linux=/usr/src/kernels/4.18.0-3.3.1.er1.x86_64 --with-zfs=no
make rpms
```

