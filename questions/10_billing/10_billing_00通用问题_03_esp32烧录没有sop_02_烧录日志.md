# q
如何从串口烧录日志确认ESP32固件烧录成功？
# a
成功烧录的典型日志特征：
- 每个分区写入后出现 `Wrote ... bytes ... Hash of data verified.`，表示数据已通过校验。
- 烧录结束出现 `Hard resetting via RTS pin...`，表示芯片被正常重置。
- 整个过程中没有 `Error` 或超时等异常信息。
例如：
```
Writing at 0x00001000... (100 %)
Wrote 24896 bytes (16262 compressed) at 0x00001000 in 0.3 seconds (effective 584.1 kbit/s)...
Hash of data verified.
...
Hard resetting via RTS pin...
```

# q
ESP32烧录日志中，地址 `0x0000e000` 处显示“不知道啥玩意，”，这个分区是什么？是否正常？
# a
该分区地址 `0x0000e000` 是ESP32的系统数据分区，通常为 `boot_app0.bin`（或 `otadata`），属于Arduino标准烧录布局的一部分，用于OTA更新或系统初始化数据。日志里的“不知道啥玩意，”仅为用户的注释，不代表错误，此次烧录是正常的。

# q
ESP32标准烧录过程中，各分区写入的顺序和地址范围是什么？
# a
根据日志，烧录顺序与典型地址如下：
1. Bootloader：`0x00001000` (`bootloader.bin`)
2. 分区表：`0x00008000` (`partitions.bin`)
3. 系统数据分区：`0x0000e000`（即注释中的“不知道啥玩意”，实际为 `boot_app0.bin` 或 `otadata`）
4. 应用程序固件：`0x00010000`（主程序 `.bin`，如 `liuxin_test.ino.bin`）
日志中的擦除与写入地址与此完全对应。

