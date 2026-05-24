# q
在VSCode中配置ESP-IDF项目后，如何打开配置菜单以设置传感器型号？
# a
在终端中运行以下命令：
```bash
idf.py menuconfig
```

# q
在menuconfig界面中，如何选择BMI270传感器？
# a
进入标有“？”或“BMI”的配置路径，找到传感器选择项（可能位于“Component config”下的传感器设置中），使用方向键选择 **BMI270**，按回车确认，然后按 **ESC** 退出并保存配置。

