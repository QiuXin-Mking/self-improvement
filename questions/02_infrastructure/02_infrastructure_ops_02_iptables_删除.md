# q
如何使用 iptables 列出 INPUT 链的规则并显示行号、详细信息和数据包计数？
# a
使用命令 `sudo iptables -L INPUT --line-numbers -v -n`。选项 `-L INPUT` 指定查看 INPUT 链，`--line-numbers` 显示行号，`-v` 显示详细信息（包括数据包和字节计数），`-n` 以数字形式显示地址和端口。

# q
如何根据行号删除 iptables 规则？
# a
使用 `sudo iptables -D INPUT <行号>` 命令，例如 `sudo iptables -D INPUT 1` 将删除 INPUT 链中行号为 1 的规则。

# q
如何添加一条允许指定源 IP 网段访问的 iptables INPUT 规则？
# a
使用 `sudo iptables -A INPUT -s <源网段> -j ACCEPT`，例如 `sudo iptables -A INPUT -s xxx.xx.xx.x/16 -j ACCEPT` 将允许该 /16 网段的所有流量进入 INPUT 链。

