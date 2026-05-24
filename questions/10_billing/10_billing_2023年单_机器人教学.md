# q
如何配置STM32的RTC预分频器使其输出1Hz信号？
# a
调用 `RTC_SetPrescaler(32767);`，基于32.768kHz晶振（32768Hz），计算公式为：`f_out = f_in / (预分频值 + 1) = 32768 / (32767 + 1) = 1Hz`，即每秒产生一个方波脉冲。

# q
12位ADC的输入电压范围是0-3.3V，数字输出值如何对应电压？
# a
12位ADC的最大数字值为 `1111 1111 1111` (0xFFF)，对应3.3V；最小值为 `0000 0000 0000`，对应0V。每个LSB代表的电压 = 3.3V / 4096 ≈ 0.806mV。

# q
在STM32 HAL库中，如何用代码控制连在PC8引脚上的继电器（低电平触发）？
# a
定义引脚和端口：
```c
#define RELAY_Pin GPIO_PIN_8
#define RELAY_GPIO_Port GPIOC
```
控制继电器动作（例如拉低使继电器吸合）：
```c
HAL_GPIO_WritePin(RELAY_GPIO_Port, RELAY_Pin, GPIO_PIN_RESET);
```

# q
UART异步串行通信最少需要连接哪几根信号线？
# a
最少需要三根线：TXD（发送）、RXD（接收）和 GND（地）。通信双方必须共地。

# q
在STM32休眠模式（__WFI）中，如何避免SysTick中断意外唤醒，确保只由RTC闹钟唤醒？
# a
调用 `HAL_SuspendTick()` 关闭SysTick中断，使其不再产生唤醒事件。需要时通过RTC秒信号或闹钟中断唤醒，同时由于SysTick被暂停，依赖SysTick的 `HAL_Delay` 将失效，可用RTC作为替代时间基准。

