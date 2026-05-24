# q
在 Home Assistant 的 configuration.yaml 中，如何配置本地 MQTT broker 的连接信息？
# a
在 `configuration.yaml` 中添加 `mqtt:` 配置块，指定 `broker: localhost`、`port: 1883`、`username` 和 `password`。例如：
```yaml
mqtt:
  broker: localhost
  port: 1883
  username: qiuxin
  password: qiuxin15970759167
```

# q
如何通过 MQTT 在 Home Assistant 中定义一个显示摄氏温度的房间温度传感器？
# a
在 `configuration.yaml` 的 `sensor` 列表中添加一个使用 `platform: mqtt` 的 sensor，指定 `name`、`state_topic`、`unit_of_measurement: "°C"`，并用 `value_template` 将消息值转为浮点数。示例：
```yaml
sensor:
  - platform: mqtt
    name: "Room Temperature"
    state_topic: "home/room/temperature"
    unit_of_measurement: "°C"
    value_template: "{{ value | float }}"
```

# q
在 MQTT 传感器的 `value_template` 中使用 `{{ value | float }}` 的目的是什么？
# a
它使用 Jinja2 模板过滤器将接收到的原始 MQTT 消息（通常为字符串）转换为浮点数，确保传感器状态被正确解析为数值类型，以便 Home Assistant 进行数值比较、图表绘制等操作。

