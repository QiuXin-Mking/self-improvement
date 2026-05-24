# q
如何在Lustre的每个MDT上注册changelog用户，以便记录特定类型的元数据变更？
# a
在每个MDT上使用 `lctl` 命令注册，指定设备、用户名和记录类型掩码。示例：
```bash
lctl --device nas_test-MDT0000 changelog_register --user root --mask OPEN
```
成功后会输出类似 `nas_test-MDT0000: Registered changelog userid 'cl1-root'`。  
注意：需要对所有MDT（如MDT0000, MDT0001, ...）分别执行此命令。

# q
如何查看指定MDT的changelog记录，例如查看最新的几条或某个范围的记录？
# a
使用 `lfs changelog` 命令。  
- 查看全部记录（显示前几行用 head）：
```bash
lfs changelog nas_test-MDT0000 | head -n 5
```
- 查看指定范围的记录（记录编号 21 到 25）：
```bash
lfs changelog nas_test-MDT0000 21 25
```

# q
如何清除指定MDT上某个changelog用户的部分记录？
# a
使用 `lfs changelog_clear` 命令，指定 MDT 设备、用户标识和要清除到的记录编号。  
示例：清除用户 `cl1-root` 在 `nas_test-MDT0000` 上编号 ≤3 的记录：
```bash
lfs changelog_clear nas_test-MDT0000 cl1-root 3
```

# q
如何查看和动态修改 Lustre MDT 的 changelog 记录类型掩码（mask）？
# a
- 查看当前掩码：
```bash
lctl get_param mdd.nas_test-MDT0000.changelog_mask
```
- 动态设置为仅记录硬链接事件：
```bash
lctl set_param mdd.nas_test-MDT0000.changelog_mask=HLINK
```
- 动态设置为记录所有事件：
```bash
lctl set_param mdd.nas_test-MDT0000.changelog_mask=ALL
```
还可以通过 `/sys/fs/lustre/mdd/<fsname>-MDTxxxx/changelog_mask` 等 sysfs 接口查看或修改。

