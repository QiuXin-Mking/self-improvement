# q
在该 ESP32 项目中，motion 参数的值 1、2、3 分别对应什么动作？
# a
- motion 1：前进
- motion 2：停止
- motion 3：后退

# q
指令通过 MQTT 下发给 ESP32 时，使用的主题格式是什么？
# a
主题格式为：`wq/v1.0/{user_id}`，例如 `wq/v1.0/9167`

# q
从日志 `Received action: {'motion': 1, 'speed': 50, 'user_id': '9167', 'user_name': '15970759167'}` 可以看出，动作指令的数据结构通常包含哪些字段？
# a
包含 `motion`（动作类型）、`speed`（速度值）、`user_id`（用户 ID）和 `user_name`（用户名）

