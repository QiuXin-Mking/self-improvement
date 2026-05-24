# q
如何使用MQTT自动发现功能在Home Assistant中添加ESP32设备及传感器？
# a
向主题 `homeassistant/device/<device_id>/config` 发布一个JSON配置消息，并设置 `retain` 标志（`-r`）。JSON需包含 `dev`（设备信息）、`cmps`（组件定义）和 `state_topic`。每个组件可指定 `p` 为 `sensor`，`device_class`（如 `temperature`、`humidity`）、`unit_of_measurement` 及 `value_template` 从上报的JSON中提取值。命令示例：
```bash
mosquitto_pub -h mqtt服务器地址 -t homeassistant/device/myesp32-dev1/config -m '@file:discovery.json' -r
```

# q
如何为ESP32新增一个MQTT传感器实体，并将其关联到某个设备？
# a
向主题 `homeassistant/sensor/<platform>/<unique_id>/config` 发送配置JSON，包含 `unique_id`、`name`、`state_topic`、`unit_of_measurement`，并在 `device` 字段指定 `identifiers`（设备唯一标识，如 `"ESP32-02"`）以及其他信息（`manufacturer`、`model` 等）。多个实体使用相同的 `identifiers` 即可归属到同一设备。

# q
MQTT传感器实体的状态值如何上报给Home Assistant？
# a
向实体配置中定义的 `state_topic` 直接发布测量值。例如，若 `state_topic` 为 `HA-ESP32-02/01/state`，则发布 `02`，Home Assistant 即更新该传感器状态。

