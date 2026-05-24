# q
创建 YUM 源需要使用的核心命令是什么？
# a
核心命令是 `createrepo`。需要事先安装该工具：
```bash
yum install createrepo
```
然后在包含 RPM 包的目录下执行 `createrepo <目录>`，会在目录中生成 `repodata` 元数据目录及相关元数据文件。

# q
如何使用 `createrepo` 从零开始创建一个本地 YUM 源？
# a
步骤如下：
1. 安装 `createrepo`：`yum install createrepo`
2. 创建存放 RPM 包的目录，并将所有 RPM 文件复制进去：
   ```bash
   mkdir -p ~/repo
   cp /path/to/*.rpm ~/repo/
   ```
3. 运行 `createrepo ~/repo`，命令会在 `~/repo` 下生成 `repodata` 目录及其中的索引和依赖元数据文件（如 `.xml.gz`、`.repomd.xml` 等）。

# q
在客户端配置本地 YUM 源时，`.repo` 文件需要包含哪些关键配置项？
# a
在 `/etc/yum.repos.d/` 下创建 `.repo` 文件，关键配置项包括：
- `[repo-id]`：仓库唯一标识
- `name`：仓库描述名称
- `baseurl`：仓库路径，本地文件使用 `file:///path/to/repo`，网络共享可使用 `http://`、`ftp://` 等
- `enabled=1`：是否启用（1 表示启用）
- `gpgcheck=0`：是否检查 GPG 签名（0 表示不检查）

示例：
```ini
[mylocalrepo]
name=My Local Repository
baseurl=file:///path/to/repo
enabled=1
gpgcheck=0
```

# q
客户端如何从新配置的本地 YUM 源安装软件包？
# a
使用 `yum install` 并指定 `--enablerepo=<repo-id>` 来启用对应仓库。例如：
```bash
yum install package-name --enablerepo=mylocalrepo
```
如果该仓库已默认启用（`enabled=1`），则可以直接安装，无需 `--enablerepo` 选项。同时需要确保客户端能访问 YUM 源目录（本地文件权限或网络共享可达）。

