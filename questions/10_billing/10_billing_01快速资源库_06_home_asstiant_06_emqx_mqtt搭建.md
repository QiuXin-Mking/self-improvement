# q
如何使用 Docker 启动 EMQX 容器？
# a
```bash
docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx
```
该命令已在本地成功执行，使用了默认的 `emqx` 镜像（社区版）。

# q
EMQX 默认暴露的主要端口及其用途是什么？
# a
- `1883`：MQTT TCP 连接端口
- `8083`：MQTT WebSocket 端口
- `8084`：MQTT WebSocket over TLS 端口
- `8883`：MQTT over TLS 端口
- `18083`：Dashboard 管理后台端口

# q
EMQX Dashboard 的默认访问地址和凭证是什么？
# a
访问地址示例：`http://103.24.176.213:18083/#/dashboard/overview`  
默认账号：`admin`  
默认密码：`public`  
注意：生产环境首次登录后应立即修改密码。

