# q
在Home Assistant的configuration.yaml中，如何配置本地MQTT Broker的连接参数？
# a
```yaml
mqtt:
  broker: localhost
  port: 1883
  username: qiuxin
  password: qiuxin15970759167
```

# q
如何通过MQTT平台在Home Assistant中定义一个温度传感器？
# a
```yaml
sensor:
  - platform: mqtt
    name: "Room Temperature"
    state_topic: "home/room/temperature"
    unit_of_measurement: "°C"
    value_template: "{{ value | float }}"
```

