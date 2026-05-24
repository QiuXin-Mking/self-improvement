# q
如何完全重置并重新启动 Home Assistant 的 Docker 容器？
# a
停止并删除旧容器，清空配置目录后重新创建容器：
```
sudo docker stop home-assistant
sudo docker rm home-assistant
rm -rf /mnt/config/.storage
rm -rf /mnt/config
mkdir -p /mnt/config
sudo docker run -d --name home-assistant --restart=unless-stopped -e TZ=Europe/London -v /mnt/config:/config -p 80:8123 ghcr.io/home-assistant/home-assistant:stable
```

# q
启动 Home Assistant 容器时如何挂载配置目录并映射端口？
# a
使用 `-v /mnt/config:/config` 挂载持久化配置，使用 `-p 80:8123` 将主机 80 端口映射到容器 8123 端口。完整命令示例：
```
sudo docker run -d --name home-assistant --restart=unless-stopped -e TZ=Europe/London -v /mnt/config:/config -p 80:8123 ghcr.io/home-assistant/home-assistant:stable
```

# q
该 Home Assistant 实例的访问域名和默认登录凭据是什么？
# a
域名：`http://www.homeassistant-qx.com`  
账号：`qiuxin`  
密码：`qiuxin@MK@159`

