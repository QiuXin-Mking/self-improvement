# q
项目中使用的STM32开发板和音频扩展板型号分别是什么？
# a
开发板型号：STM32 Nucleo F401RE；扩展板型号：CCA02M2（MEMS麦克风扩展板）。

# q
计划用于蓝牙通信的模块型号是什么？
# a
选用 **CC2540 AT-09** 模块（基于 TI CC2540 的 BLE 模块）。

# q
音频信号处理准备采用什么优化技术？
# a
使用 **双缓冲（double buffering）** 技术，交替利用两个缓冲区进行数据接收和处理，以提高实时性和吞吐量。

# q
双缓冲区方案中 buffer A 和 buffer B 的具体分工和大小如何？
# a
```plaintext
buffer B  size 8K  <-- INPUT
buffer B           --> buffer A
buffer A  size 8K  --> process
```
即设置两个 8K 缓冲区，输入数据先填入 buffer B，随后转移到 buffer A 进行处理。

# q
该项目的音频处理核心任务是什么？
# a
通过音频信号处理，检测黄蜂的声音并区分黄蜂与其他昆虫。

