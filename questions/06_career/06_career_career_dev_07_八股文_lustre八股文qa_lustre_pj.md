# q
如何使用 lfs 命令为 Lustre 目录递归设置 project ID 并启用继承？
# a
```
lfs project -s -p 3 -r /mnt/lustre/zzq_test
```
- `-r`：递归处理所有子目录和文件  
- `-s`：设置 PROJID_INHERIT 属性，使新文件和子目录继承父目录的 project ID  
- `-p <ID>`：指定 project ID（此处为3）

# q
如何查看 Lustre 文件或目录的 project ID？
# a
方法一：
```
lfs project /mnt/lustre/zzq_test/yourscript.py
```
方法二：
```
getfattr -n trusted.projid /mnt/lustre/zzq_test
```
输出示例：
```
# file: mnt/lustre/zzq_test
trusted.projid="3"
```

# q
如何通过扩展属性设置文件不可变以防止误删除，以及如何查询和移除该属性？
# a
- 设置不可变：
  ```
  chattr +i /mnt/lustre/goofys
  ```
- 查询属性：
  ```
  lsattr /mnt/lustre/goofys
  ```
  输出中带 `i` 表示不可变。
- 移除不可变：
  ```
  chattr -i /mnt/lustre/goofys
  ```

# q
如何查看 Lustre 挂载点的文件系统类型、容量和挂载选项？
# a
方法一：使用 `df -T`
```
df -T /mnt/lustre/
```
示例输出：
```
Filesystem                           Type    1K-blocks      Used Available Use% Mounted on
172.31.0.26@tcp:.../nas              lustre  1123820852 393704572 672264376  37% /mnt/lustre
```
方法二：使用 `mount` 并过滤
```
mount | grep /mnt/lustre
```
示例输出中可看到挂载选项（如 `checksum`, `flock`, `statfs_project` 等）。

