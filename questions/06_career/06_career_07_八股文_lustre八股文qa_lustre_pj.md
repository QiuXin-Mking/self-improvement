# q
在 Lustre 中如何为目录设置 project ID 并使其子项继承该 ID？
# a
使用命令：
```bash
lfs project -s -p 3 -r /mnt/lustre/zzq_test
```
- `-r`：递归处理该目录下的所有子目录和文件。
- `-s`：设置 `PROJID_INHERIT` 属性，使新创建的文件和子目录自动继承父目录的 project ID。
- `-p <ID>`：指定要设置的 project ID（此处为 3）。

# q
如何查询文件或目录在 Lustre 中的 project ID？
# a
两种方法：
1. 使用 `lfs project` 命令：
   ```bash
   lfs project /mnt/lustre/zzq_test/yourscript.py
   ```
2. 使用 `getfattr` 查看扩展属性：
   ```bash
   getfattr -n trusted.projid /mnt/lustre/zzq_test
   ```
   输出示例中 `trusted.projid="3"` 表示 project ID 为 3。

# q
如何利用 ext4 的 `chattr` 保护文件或目录防止误删除？
# a
- 设置不可变属性（只读，不可删除或修改）：
  ```bash
  chattr +i /mnt/lustre/goofys
  ```
- 移除不可变属性：
  ```bash
  chattr -i /mnt/lustre/goofys
  ```
- 查看属性：
  ```bash
  lsattr /mnt/lustre/goofys
  ```
  输出中 `----i-----------------` 表示设置了 `i` 属性。

# q
如何查看某个挂载点的文件系统类型和挂载详情？
# a
两种常用方法：
1. 使用 `df -T` 查看文件系统类型及使用情况：
   ```bash
   df -T /mnt/lustre/
   ```
   输出示例显示文件系统类型为 `lustre`。
2. 使用 `mount` 过滤查看挂载参数：
   ```bash
   mount | grep /mnt/lustre
   ```
   输出中会包含设备路径、挂载点、文件系统类型（`lustre`）及挂载选项（如 `checksum`, `flock`, `encrypt` 等）。

