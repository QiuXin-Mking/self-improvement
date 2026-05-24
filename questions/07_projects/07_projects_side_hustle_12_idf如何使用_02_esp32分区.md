# q
ESP32的分区表使用什么格式定义？
# a
使用CSV表格定义，包含列：Name, Type, SubType, Offset, Size, Flags。

# q
示例分区表中，nvs分区的类型、子类型、偏移和大小是多少？
# a
类型为data，子类型为nvs，偏移地址0x9000，大小0x4000字节。

# q
示例分区表中，ota_0和ota_1分区的子类型、偏移和大小是多少？
# a
子类型分别为ota_0和ota_1，ota_0偏移0x110000，ota_1偏移0x210000，大小均为1M。

# q
sdkconfig.defaults文件在ESP-IDF项目中的作用是什么？
# a
提供默认配置选项，在首次生成sdkconfig文件时应用；sdkconfig是项目的实际配置文件，包含所有配置选项。

