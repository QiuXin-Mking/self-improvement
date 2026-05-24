# q
SSH 免密鉴权如何批量配置集群节点间的双向互信？
# a
使用 `ssh-keygen` 生成密钥对，`ssh-copy-id` 拷贝公钥到目标节点，目标节点会自动写入 `authorized_keys`。可以从文件读取 IP 列表循环执行，实现集群节点间免密登录。

# q
iptables 防火墙规则顺序的核心注意事项是什么？
# a
iptables 按规则顺序匹配，必须先添加 ACCEPT 白名单规则，再添加 DROP 拒绝所有规则；若 DROP 在前将拦截所有流量。规则持久化路径因发行版而异（RHEL: `/etc/sysconfig/iptables`, Debian: `/etc/iptables/rules.v4`）。

# q
使用 YUM 在离线环境搭建本地 HTTP 源的完整流程是什么？
# a
```bash
yumdownloader --resolve <package>   # 下载包及所有依赖
createrepo .                        # 创建元数据仓库
python -m SimpleHTTPServer 80       # 快速启动 HTTP 源
```
每次 RPM 变动后需执行 `createrepo --update` 更新元数据。

# q
rsync 路径末尾的斜杠对同步行为有何影响？
# a
- `rsync -r /src/dir /dst` → 拷贝整个 dir 目录到 /dst 下，结果为 /dst/dir
- `rsync -r /src/dir/ /dst` → 拷贝 dir 目录下的全部内容到 /dst 目录（不含 dir 本身）

# q
TCP 内核参数 `tcp_rmem` 和 `rmem_max` 在高带宽环境下的调优目标是什么？
# a
将 `net.core.rmem_max` 和 `net.core.wmem_max` 调大（如 1.6GB 级别），`net.ipv4.tcp_rmem` 和 `net.ipv4.tcp_wmem` 的最大值调至 100MB+，以提升高带宽网络的 TCP 吞吐能力。使用 `sysctl -w` 即时生效，写入 `/etc/sysctl.conf` 持久化。

