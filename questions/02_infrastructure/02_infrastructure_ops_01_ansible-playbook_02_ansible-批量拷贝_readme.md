# q
如何使用 ansible copy 模块批量拷贝文件并设置权限和所有者？
# a
通过 `ansible` 命令的 `-m copy` 模块，在 `-a` 参数中指定源路径、目标路径、所有者、组、权限及目录权限。示例：
```bash
ansible test -m copy -a "src=/root/rpm dest=/root/ owner=root group=root mode=0755 directory_mode=0755"
```

# q
如何使用 ansible 批量在远程主机上强制安装 RPM 包？
# a
使用 `shell` 模块执行 `rpm` 命令，添加 `--force --nodeps` 选项忽略依赖和强制安装：
```bash
ansible test -m shell -a "rpm -Uvh --force --nodeps /root/rpm/*.rpm"
```

# q
挂载 NFS 时出现 `bad option; ... helper program` 错误的原因和解决方法是什么？
# a
原因：系统缺少 NFS 挂载辅助工具（如 `/sbin/mount.nfs`），通常因为未安装 `nfs-utils` 包。
解决方法：
```bash
yum install -y nfs-utils
```

# q
ansible 的 `copy` 模块与 `shell` 模块的主要用途区别是什么？
# a
`copy` 模块专门用于从控制节点向远程主机拷贝文件，并可在拷贝时设置文件属性；`shell` 模块用于在远程主机上执行任意 shell 命令，如安装软件、检查版本等。

