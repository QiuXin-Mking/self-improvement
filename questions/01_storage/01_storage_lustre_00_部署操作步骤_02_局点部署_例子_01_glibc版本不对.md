# q
如何查询当前系统的 glibc 版本？
# a
使用命令 `ldd --version` 查看 glibc 版本信息。

# q
升级 glibc 的核心操作步骤是什么？
# a
```bash
# 1. 在解压的 glibc 目录下安装 RPM 包
rpm -Uvh *
# 2. 重建 RPM 数据库
rpmdb --rebuilddb
# 3. 在解压的 gcc 目录下安装 GCC 相关的 RPM 包
rpm -Uvh *
```

# q
升级 glibc 时遇到依赖报错 `glibc-common = 2.17-106.el7_2.4 is needed by (installed) glibc-2.17-106.el7_2.4.i686` 应如何解决？
# a
使用 `rpm -e --nodeps` 卸载有冲突的旧版本包，忽略依赖关系：
```bash
rpm -e --nodeps glibc-2.17-106.el7_2.4.i686
```

