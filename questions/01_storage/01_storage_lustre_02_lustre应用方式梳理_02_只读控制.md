# q
如何查看Lustre MDT的只读状态？
# a
使用以下命令查看指定MDT设备的只读参数值（0 表示可写，1 表示只读）：
```bash
lctl get_param mdt.*-MDT0000.readonly
```

# q
如何动态设置Lustre MDT为只读或可写模式？
# a
通过 `lctl set_param` 修改 MDT 的 `readonly` 参数：
- 设置为只读：
  ```bash
  lctl set_param mdt.<文件系统名>-MDT0000.readonly=1
  ```
- 设置为可写：
  ```bash
  lctl set_param mdt.<文件系统名>-MDT0000.readonly=0
  ```
例如：
```bash
lctl set_param mdt.nas_test-MDT0000.readonly=1
```

# q
如何查询某个文件所属的 MDT 索引？
# a
使用 `lfs getstripe` 命令的 `--mdt-index` 选项，指定文件路径：
```bash
lfs getstripe --mdt-index /mnt/lustre/test_175
```
该命令会返回该文件元数据所在的 MDT 索引编号。

