# q
在rcache的coredump中，write_cmd_split_blist失败的直接原因是什么？
# a
从gdb输出可以看出，child_ctx的`bl`字段为`0x0`（NULL），而`split_size=65536`（64KB）。`write_cmd_split_blist`需要基于有效的`bl`（bio list）进行分裂操作，由于`bl`为空导致分裂失败，进而触发coredump。

# q
笔记中“命中算法有问题，主线是一个ocache的cache对应一个block”揭示了什么根本原因？
# a
该注释指出`rw_ctx_split_hit_cache`相关的命中逻辑存在缺陷。主流程中一个ocache cache条目对应一个block，但读提升（read promotion）过程中，child_ctx的块大小（如1MB）与cache block大小（64KB）不对齐，导致命中后无法正确生成或分裂blist，最终出现`bl`为NULL的错误。

# q
在该coredump分析中，child_ctx的`size`和`split_size`分别是多少，各有什么含义？
# a
- `size = 1048576`（1MB），表示child ctx所覆盖的IO请求大小。
- `split_size = 65536`（64KB），是写命令分割时每个子请求的单位大小，对应一个cache block的大小。
两者不匹配时（1MB vs 64KB），需要将大请求按64KB单位分割成多个子blist，但由于`bl`为空，分割失败。

