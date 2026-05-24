# q
如何获取指定 epoch 的 OSD map 并保存为二进制文件？
# a
使用 `ceph osd getmap` 命令，指定 `--epoch` 和输出文件即可：
```bash
ceph osd getmap --epoch 100 -o osdmap-100.bin
```

# q
如何将 OSD map 的二进制文件转换为人类可读的文本内容？
# a
使用 `osdmaptool` 工具，配合 `--print` 选项：
```bash
osdmaptool --print osdmap.bin
```

