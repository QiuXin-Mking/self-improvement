# q
如何使用 Docker 启动 EMQX MQTT 服务器（以本地成功执行为例）？
# a
```bash
docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 18083:18083 emqx
```

# q
EMQX Web Dashboard 的默认访问端口和默认登录凭证是什么？
# a
访问端口：18083（Web Dashboard），登录方式 `http://<IP>:18083`。  
默认账号：`admin`，默认密码：`public`。

# q
文档中记录的 EMQX 实例修改后的密码是什么？
# a
修改后的密码为 `qiuxin_15970759167`。

