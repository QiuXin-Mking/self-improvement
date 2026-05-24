# q
为什么 STM32 的 PB3、PB4、PA13、PA14、PA15 引脚默认不能作为普通 GPIO 使用？
# a
因为它们默认被 JTAG/SWD 调试接口占用。PA13、PA14 用于 SWD，PB3、PB4、PA15 用于 JTAG，所以上电后这些引脚处于调试功能模式，无法直接作为普通 I/O 引脚。

# q
如何将 PB3、PB4、PA15 释放为普通 GPIO（同时保留 SWD 调试功能）？
# a
需要开启 AFIO 时钟并禁用 JTAG 接口：
```c
RCC_APB2PeriphClockCmd(RCC_APB2Periph_AFIO, ENABLE); // 开启 AFIO 时钟
GPIO_PinRemapConfig(GPIO_Remap_SWJ_JTAGDisable, ENABLE); // 禁用 JTAG，保留 SWD
```
之后再正常配置这些引脚的 GPIO 模式即可。

