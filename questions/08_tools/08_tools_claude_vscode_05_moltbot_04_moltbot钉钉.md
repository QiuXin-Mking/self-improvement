# q
如何通过命令行配置 clawdbot 的钉钉频道？
# a
使用 `clawbot config set` 命令，指定键为 `channels.dingtalk`，值为完整的 JSON 配置对象，包含 `enabled`、`clientId` 和 `clientSecret` 字段。示例：
```
clawdbot config set 'channels.dingtalk' '{
  "enabled": true,
  "clientId": "dingmecaiuzug9jg5ekc",
  "clientSecret": "vwtBvm_ZN0mw3z38PHlIXd11ZdnsL_S7D4qKnc0hz0h9WJ31qT8z42LerebhQKgv"
}'
```

# q
钉钉频道配置 JSON 中必须包含哪些字段？
# a
必须包含三个字段：`enabled`（布尔值，是否启用）、`clientId`（字符串，钉钉应用 Client ID）、`clientSecret`（字符串，钉钉应用 Client Secret）。

