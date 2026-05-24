# q
如何使用命令查询 Lustre 文件系统中文件或目录的 FID？
# a
使用 `lfs path2fid` 命令，例如：
```bash
lfs path2fid <path>
```
若要显示路径中每一级父目录的 FID，可加 `--parents` 参数：
```bash
lfs path2fid --parents <path>
```
示例输出：
```
chenlou: [0x200000403:0xc:0x0]
lustre_create_sh.py: [0x200000403:0x6:0x0]
```

