# q
ESP32烧录成功的日志中通常会依次出现哪些关键步骤和提示信息？
# a
正常烧录日志的典型步骤序列：
1. 连接串口：`Connecting.....`
2. 识别芯片：`Chip is ESP32-D0WD-V3 (revision v3.1)`
3. 上传 stub：`Uploading stub...` → `Running stub...` → `Stub running...`
4. 更改波特率：`Changing baud rate to 921600`
5. 配置 Flash 大小：`Configuring flash size...`
6. 擦除分区：`Flash will be erased from 0x... to 0x...`
7. 写入各分区（bootloader、partitions、ota data、app）并逐百分比显示进度，每个分区写入后会出现 `Hash of data verified.`
8. 完成硬复位：`Hard resetting via RTS pin...`

缺少上述任一步骤或出现错误（如 `Timed out waiting for packet header`），均意味着烧录未成功。

# q
如何从烧录日志判断 ESP32 是否借助自动复位电路进入了下载模式？
# a
如果日志末尾出现 `Hard resetting via RTS pin...`，说明烧录工具通过 RTS/DTR 引脚自动控制了 EN 和 IO0 的电平时序，使芯片直接进入下载模式并完成烧录，无需手动按下 BOOT 按钮。若日志卡在 `Connecting...` 阶段，且未见该提示，通常意味着自动复位电路未起作用，必须采用手动按键或拉低 IO0 的方式才能进入下载模式。

# q
ESP32 烧录日志中四个被写入的分区地址分别存放什么内容？
# a
- `0x00001000`：bootloader（引导程序）
- `0x00008000`：partition table（分区表）
- `0x0000e000`：ota data（OTA 数据分区，日志中标注为“不知道啥玩意”）
- `0x00010000`：application firmware（应用固件）

每个分区在写入完成后都会进行哈希校验并显示 `Hash of data verified.`，以确保传输正确性。

