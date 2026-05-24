# q
如何在 Home Assistant 的 configuration.yaml 中配置 MQTT 代理的连接信息？
# a
在 `configuration.yaml` 中添加 `mqtt` 字段，指定 `broker`（如 localhost）、`port`（如 1883）、`username` 和 `password`。示例：
```yaml
mqtt:
  broker: localhost
  port: 1883
  username: qiuxin
  password: qiuxin15970759167
```

# q
如何通过 MQTT 传感器获取房间温度数据，并正确设置单位与值模板？
# a
在 `sensor` 下定义 `platform: mqtt` 的传感器，设置以下关键参数：
- `name`: "Room Temperature"
- `state_topic`: "home/room/temperature"
- `unit_of_measurement`: "°C"
- `value_template`: `"{{ value | float }}"` （将接收到的值转换为浮点数）

# q
如何通过 MQTT 传感器获取房间湿度数据？
# a
与温度传感器类似，配置不同的名称、主题和单位：
```yaml
sensor:
  - platform: mqtt
    name: "Room Humidity"
    state_topic: "home/room/humidity"
    unit_of_measurement: "%"
    value_template: "{{ value | float }}"
```

