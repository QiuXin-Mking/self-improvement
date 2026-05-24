# q
通过MQTT发现协议向Home Assistant注册一个ESP32设备时，配置消息应发布到哪个Topic？
# a
发布到 `homeassistant/device/<device_id>/config`，例如 `homeassistant/device/myesp32-dev1/config`，建议使用 retain 标志（`-r`）以便Home Assistant重启后仍能保留设备。

# q
在ESP32设备的MQTT发现配置中，如何将多个传感器（如温度和湿度）关联到同一个物理设备？
# a
在配置JSON的`device`字段中设置相同的`identifiers`（如 `"ESP32-02"`），这样所有实体都会归属到同一个设备（例如名为“ESP32-02”的设备）。

# q
使用`mosquitto_pub`命令行工具发布保留的MQTT发现消息，命令格式是什么？
# a
```bash
mosquitto_pub -h <mqtt服务器地址> -t homeassistant/device/myesp32-dev1/config -m '@file:discovery.json' -r
```
其中`-r`表示保留消息，`@file`表示从文件读取JSON内容。

# q
为ESP32新增一个温度传感器实体时，MQTT发现配置中必须包含哪些关键字段？
# a
必须包含：
- `unique_id`：实体的唯一ID，如 `HA-ESP32-02-01`
- `name`：实体显示名称，如“温度传感器”
- `state_topic`：用于接收状态值的主题，如 `HA-ESP32-02/01/state`
- `device`：包含`identifiers`等字段以关联到具体设备
- `unit_of_measurement`（可选但推荐）如 `°C`
- `icon`（可选）如 `mdi:thermometer`

# q
在Home Assistant的MQTT发现中，一个设备下如何添加多个同类型实体（例如两个温度传感器）？
# a
为每个实体发布独立的发现消息到各自的主题，如：
- `homeassistant/sensor/HA/HA-ESP32-02-01/config`
- `homeassistant/sensor/HA/HA-ESP32-02-02/config`
在各自的JSON中使用不同的`unique_id`和`state_topic`，但`device.identifiers`保持一致即可归属到同一设备。

