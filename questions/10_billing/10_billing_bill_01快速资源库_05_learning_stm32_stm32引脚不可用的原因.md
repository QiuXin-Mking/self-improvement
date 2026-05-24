# q
为什么STM32的PB3、PB4、PA13、PA14、PA15引脚不能直接作为普通GPIO使用？
# a
这些引脚默认复用为JTAG/SWD调试接口。要作为GPIO使用，需先开启AFIO时钟，并禁用JTAG（保留SWD）。示例代码：
```c
RCC_APB2PeriphClockCmd(RCC_APB2Periph_AFIO, ENABLE);  // 开启AFIO时钟
GPIO_PinRemapConfig(GPIO_Remap_SWJ_JTAGDisable, ENABLE);  // 禁用JTAG
```

