# q
CephFS 挂载后执行 `ls` 提示 `Permission denied` 的典型根因是什么？
# a
通常是客户端认证凭据不正确或挂载时缺少正确的 keyring 文件，导致挂载用户没有访问 CephFS 的权限。常见场景包括：`mount` 命令中 `name` 参数指定的客户端 ID 未授权，或使用的 `secret` 与 Ceph 集群中的密钥不匹配。

# q
如何从命令行快速检查 CephFS 挂载状态以定位权限问题？
# a
执行 `mount | grep <挂载点>` 查看挂载详情，确认挂载源和挂载选项中的客户端名称（`name=manila`）及是否包含 ACL。例如：
```
192.168.5.128:6789,192.168.5.129:6789,192.168.5.165:6789:/volumes/manila/cc75f301-82d7-4926-913d-5738cec4d968 on /mnt type ceph (rw,relatime,name=manila,secret=<hidden>,acl)
```
若输出中 `name` 与预期客户端不符，或缺少 ACL 标记，通常会导致访问被拒绝。

# q
解决 CephFS 挂载后 `Permission denied` 的标准排查流程是什么？
# a
1. 确认挂载已生效且选项正常：`mount | grep /mnt`
2. 检查客户端认证权限：`ceph auth get client.<name>`，核对客户端密钥和访问能力
3. 使用正确的 keyring 重新挂载，例如：
   ```bash
   mount -t ceph 192.168.5.128:6789,192.168.5.129:6789,192.168.5.165:6789:/volumes/manila/<volume-id> /mnt -o name=manila,secret=<key>
   ```
   或使用 keyring 文件：
   ```bash
   ceph-fuse -m mon1,mon2,mon3 /mnt/cephfs --id <client_id> --keyring /etc/ceph/ceph.client.<client_id>.keyring
   ```
4. 重新挂载后验证目录访问权限（如 `ls /mnt`）

