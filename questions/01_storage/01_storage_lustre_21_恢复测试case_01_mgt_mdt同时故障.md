# q
Lustre MGT和MDT同时故障时，恢复MDT服务的标准修复流程是什么？
# a
1. 在另一个可用节点上使用 `rbd map` 和 `mount -t lustre` 启动 MGT 服务（例如从节点 174 挂载 MGT）。
2. 在待恢复的 MDT 所在节点上，使用 `tunefs.lustre` 修改 MDT 的 `mgsnode` 参数，指向新的 MGT 节点，同时保留 `servicenode`。
   ```bash
   tunefs.lustre --erase-params --servicenode=192.168.6.172@tcp --mgsnode=192.168.6.174@tcp /dev/rbd0
   ```
3. 挂载 MDT 设备。
4. 如需将 MDT 的 MGS 指向恢复回原始节点，再次执行 `tunefs.lustre` 将 `mgsnode` 改回原地址，然后重新挂载。
   ```bash
   tunefs.lustre --erase-params --servicenode=192.168.6.172@tcp --mgsnode=192.168.6.172@tcp /dev/rbd0
   mount -t lustre /dev/rbd0 /data/lustre_mdt
   ```

# q
如何通过命令行修改 Lustre MDT 设备的 MGS 节点地址？
# a
使用 `tunefs.lustre` 带 `--erase-params` 清除原有参数，并指定新的 `--mgsnode` 和 `--servicenode`。
```bash
tunefs.lustre --erase-params --servicenode=<本节点NID> --mgsnode=<新MGS的NID> /dev/<MDT设备>
```
示例：
```bash
tunefs.lustre --erase-params --servicenode=192.168.6.172@tcp --mgsnode=192.168.6.174@tcp /dev/rbd0
```

# q
挂载 Lustre MDT 时提示 “mount.lustre: mount … failed: Operation already in progress” 的常见原因及解决方法是什么？
# a
原因：通常是由于前一次挂载失败或未完全终止，导致 `/sbin/mount.lustre` 进程残留，系统认为目标服务已经在运行。
解决方法：
1. 查找残留的 mount 进程：
   ```bash
   ps aux | grep mount
   ```
2. 找到对应的 `mount.lustre` 进程并将其 kill，例如 `kill <PID>`。
3. 确认进程已经结束，然后重新执行 `mount -t lustre` 命令。

