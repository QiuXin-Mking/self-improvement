# q
解决ESP32 Spark烧录异常时，如何在menuconfig中配置BMI270传感器驱动？
# a
在项目根目录执行 `idf.py menuconfig`，在配置界面按 `/` 搜索“BMI”，选择“BMI270”相关选项，然后保存并退出（Esc → 保存）。

