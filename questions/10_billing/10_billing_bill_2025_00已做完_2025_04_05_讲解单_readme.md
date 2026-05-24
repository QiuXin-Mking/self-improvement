# q
如何使用 STM32 的低级库（LL）初始化 USART1，设置波特率 115200、8 数据位、无校验、1 停止位？
# a
```c
USART_InitStruct.BaudRate = 115200;
USART_InitStruct.DataWidth = LL_USART_DATAWIDTH_8B;
USART_InitStruct.StopBits = LL_USART_STOPBITS_1;
USART_InitStruct.Parity = LL_USART_PARITY_NONE;
USART_InitStruct.TransferDirection = LL_USART_DIRECTION_TX_RX;
USART_InitStruct.HardwareFlowControl = LL_USART_HWCONTROL_NONE;
USART_InitStruct.OverSampling = LL_USART_OVERSAMPLING_16;
LL_USART_Init(USART1, &USART_InitStruct);
LL_USART_ConfigAsyncMode(USART1);
LL_USART_Enable(USART1);
```

# q
IIC 总线写单个寄存器的典型时序是什么？请描述流程并给出示例代码。
# a
流程：主设备发送起始条件 → 发送从机地址（写位 0）并等待 ACK → 发送寄存器地址并等待 ACK → 发送数据字节并等待 ACK → 发送停止条件。
示例函数：
```c
void OLED_I2C_WriteByte(uint8_t addr, uint8_t data) {
    OLED_IIC_Start();
    OLED_IIC_Send_Byte(OLED_ADDRESS);
    OLED_IIC_Wait_Ack();
    OLED_IIC_Send_Byte(addr);
    OLED_IIC_Wait_Ack();
    OLED_IIC_Send_Byte(data);
    OLED_IIC_Wait_Ack();
    OLED_IIC_Stop();
}
```

# q
ULN2003A 达林顿管在弱电控制强电时的逻辑真值表是怎样的？为什么有电压差才会产生电流？
# a
真值表：输入端为 1（高电平）时，输出端为 0（低电平，导通）；输入端为 0 时，输出端为 1（截止）。电压差是驱动电流的必要条件，只有存在电位差，电子才会定向移动形成电流。该系统可以将 0～3.3V 的弱电信号转换为 0～5V（或更高）的驱动输出。

