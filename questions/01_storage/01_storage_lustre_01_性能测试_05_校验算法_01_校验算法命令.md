# q
如何查看Lustre客户端每个OSC当前使用的校验和算法？
# a
使用命令：
```bash
lctl get_param osc.*.checksum_type
```
输出会列出每个OST对应的OSC所支持的校验和算法列表，方括号内的算法为当前选中的校验和类型。例如：
```
osc.nas_test-OST0000-osc-ffff930cba86c800.checksum_type=
crc32 adler crc32c t10ip512 [t10ip4K] t10crc512 t10crc4K
```
表示该OSC当前使用的是t10ip4K校验和。

# q
`lctl get_param osc.*.checksum_type`输出中方括号的作用是什么？
# a
方括号内的算法表示该OSC当前激活的校验和算法，未被方括号包裹的为支持但未选用的算法。

# q
从`lctl get_param osc.*.checksum_type`的输出中可以看到Lustre支持哪些校验和算法？
# a
常见支持的校验和算法包括：crc32、adler、crc32c、t10ip512、t10ip4K、t10crc512、t10crc4K。

