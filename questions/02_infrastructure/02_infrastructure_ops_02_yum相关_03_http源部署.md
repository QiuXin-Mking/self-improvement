# q
createrepo 命令的作用是什么？
# a
createrepo 用于根据指定目录中的 RPM 包创建 YUM 仓库的元数据（repodata），常用参数：
- `-pdo /目标目录 /源RPM目录`：将仓库元数据输出到目标目录
- `-update /仓库目录`：在有 RPM 包变化时更新仓库元数据

示例：
```bash
createrepo -pdo /home/yumrepo /home/yumrepo
createrepo -update /home/yumrepo
```

# q
如何使用 SimpleHTTPServer 快速搭建 HTTP YUM 仓库？
# a
在包含 RPM 包及 repodata 的目录（如 /home/yumrepo）下执行：
```bash
nohup python -m SimpleHTTPServer 8080 &>/dev/null &
```
该命令会在后台启动一个监听 8080 端口的 HTTP 服务，将当前目录作为 YUM 源的根路径提供访问。

# q
客户端如何配置基于 HTTP 的自建 YUM 源？
# a
在 `/etc/yum.repos.d/` 下创建 `.repo` 文件，配置示例如下：
```ini
[Myceph]
name=Myceph
baseurl=http://10.3.194.155:80
enable=1
gpgcheck=0
gpgkey=
```
配置完成后执行 `yum makecache` 更新缓存。若需仅从该源测试查询软件包并排除其他源，可使用：
```bash
yum --enablerepo=Myceph --disablerepo=base,extras,updates,epel search ceph
```

