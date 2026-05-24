# q
查询 Lustre 文件或目录的 FID 的命令是什么？
# a
使用 `lfs path2fid` 命令。基本用法：
```bash
lfs path2fid <文件或目录路径>
```
示例：
```bash
lfs path2fid chenlou lustre_create_sh.py
# 输出：chenlou: [0x200000403:0xc:0x0]
#        lustre_create_sh.py: [0x200000403:0x6:0x0]
```

# q
`lfs path2fid` 输出的 FID 格式是什么样的？
# a
FID 的格式为 `[0x十六进制:0x十六进制:0x0]` 的形式，例如 `[0x200000403:0xc:0x0]`，包含三个以冒号分隔的部分，每个部分为十六进制数。

# q
如何使用 `lfs path2fid` 保留父目录信息？
# a
添加 `--parents` 选项，输出会在 FID 后保留原路径的父目录前缀。示例：
```bash
lfs path2fid --parents stg_ssd_4-101-221/Fio.0.0
# 输出：[0x200000407:0x65:0x0]/Fio.0.0
```

