# q
LoRa和LoRaWAN有什么区别？
# a
LoRa是一种低功耗、远距离的局域网无线调制技术，属于物理层，由Semtech公司研发。LoRaWAN是基于LoRa技术的低功耗广域通信协议，定义了物理层和MAC层，由LoRa联盟维护，用于构建广域物联网。

# q
使用树莓派的GPIO控制LED闪烁的核心代码逻辑是什么？
# a
使用RPi.GPIO库，设置GPIO模式为BCM，将对应引脚（如17）设为输出。然后在循环中交替设置引脚为高电平（点亮）和低电平（熄灭），并通过time.sleep控制间隔时间，实现闪烁效果。示例代码通过捕获KeyboardInterrupt来安全退出并清理GPIO状态。

# q
典型的树莓派LoRaWAN到Home Assistant的通信链路是怎样的？
# a
树莓派通过LoRa无线发送数据 → LoRaWAN网关接收并转发至网络服务器 → 服务器上的Home Assistant处理数据 → 最终通过App查看或控制。

