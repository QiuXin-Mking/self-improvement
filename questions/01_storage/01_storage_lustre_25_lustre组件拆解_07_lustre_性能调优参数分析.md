# q
Lustre中checksum_pages参数的作用是什么，如何查看和修改？
# a
`checksum_pages` 是 Lustre 客户端 `llite` 模块的参数，用于控制客户端与内核之间数据传输时是否启用页校验和。可通过以下命令查看：
```bash
lctl get_param llite.*.checksum_pages
```
通过以下命令永久修改为0（关闭）：
```bash
lctl set_param -P llite.<fsname>-<instance>.checksum_pages=0
```

# q
Lustre的max_rpcs_in_flight参数含义是什么，如何批量调整？
# a
`max_rpcs_in_flight` 是 Lustre 中控制客户端到 MDT/OST 同时处于飞行状态的最大 RPC 请求数量的参数，影响元数据与数据 I/O 的并发度。可使用通配符查看：
```bash
lctl get_param *.nas_test*.max_rpcs_in_flight
```
通过 ansible 批量调整为256：
```bash
ansible lustre -m shell -a 'lctl set_param *.nas_test*.max_rpcs_in_flight=256'
```

# q
Lustre中max_dirty_mb参数的作用是什么，如何设置？
# a
`max_dirty_mb` 控制 Lustre 客户端允许缓存的最大脏数据量（单位 MB），超出后客户端会强制将数据刷新到 OST。查看命令：
```bash
lctl get_param *.nas_test*.max_dirty_mb
```
批量设置为200 MB：
```bash
ansible lustre -m shell -a 'lctl set_param *.nas_test*.max_dirty_mb=200'
```

# q
fio输出中slat、clat和lat分别代表什么延迟？
# a
- **slat (submit latency)**：提交延迟，从发起 I/O 操作到提交给操作系统的时间。
- **clat (completion latency)**：完成延迟，I/O 操作从提交到完成的时间。
- **lat (latency)**：总延迟，包含提交和完成延迟。三者常以 min、max、avg、stdev 统计，单位微秒或毫秒。

