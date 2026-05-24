# q
如何通过 /proc 文件系统遍历所有进程并统计每个进程打开的文件描述符数量，按数量降序排列？
# a
使用以下命令，遍历所有 PID，统计 `/proc/$pid/fd` 目录下的文件数量，并按 FD 数量降序排列：
```bash
for pid in $(ps -e -o pid | tail -n +2); do
    echo -n "PID: $pid, FD Count: ";
    ls /proc/$pid/fd | wc -l;
done | sort -k5 -n -r
```

# q
如何查看当前 shell 的文件描述符软限制？如何永久提高该限制？
# a
- 查看当前软限制：`ulimit -n`
- 永久修改：编辑 `/etc/security/limits.conf`，添加如下行，然后重新登录或重启系统：
```
* soft nofile <新限制数>
* hard nofile <新限制数>
```

# q
如何使用 lsof 统计每个进程当前打开的文件描述符数量，并按数量降序显示？
# a
使用管道命令：
```bash
lsof | awk '{print $2}' | sort | uniq -c | sort -nr
```
解释：
- `lsof` 列出打开的文件
- `awk '{print $2}'` 提取 PID
- `sort` 排序 PID
- `uniq -c` 统计每个 PID 出现次数（即 FD 数量）
- `sort -nr` 按次数逆序排列

# q
如何查看 systemd 管理服务的默认文件描述符限制（DefaultLimitNOFILE）？
# a
运行：
```bash
systemctl show --property=DefaultLimitNOFILE
```

