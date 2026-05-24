# q
如何在 Home Assistant 的 configuration.yaml 中配置本地 MQTT broker 连接？
# a
```yaml
mqtt:
  broker: localhost
  port: 1883
  username: qiuxin
  password: qiuxin15970759167
```

# q
如何在 Home Assistant 中通过 MQTT 定义一个温度传感器？
# a
```yaml
sensor:
  - platform: mqtt
    name: "Room Temperature"
    state_topic: "home/room/temperature"
    unit_of_measurement: "°C"
    value_template: "{{ value | float }}"
```

# q
MQTT 传感器配置中的 value_template 有什么作用？示例中为何使用 `{{ value | float }}`？
# a
`value_template` 用于从 MQTT 消息的 payload 中提取传感器数值。`{{ value | float }}` 将收到的字符串转换为浮点数，确保数值类型正确，避免因类型不匹配引起的错误。

