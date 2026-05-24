# q
如何完全重置并重新部署 Home Assistant Docker 容器？
# a
依次执行以下命令：
```bash
sudo docker stop home-assistant
sudo docker rm home-assistant
rm -rf /mnt/config/.storage
rm -rf /mnt/config
mkdir -p /mnt/config
sudo docker run -d --name home-assistant --restart=unless-stopped -e TZ=Europe/London -v /mnt/config:/config -p 80:8123 ghcr.io/home-assistant/home-assistant:stable
```

# q
Home Assistant 容器重置后，使用的访问域名、账号和密码是什么？
# a
- 域名：`http://www.homeassistant-qx.com`
- 账号：`qiuxin`
- 密码：`qiuxin@MK@159`

# q
在重置 Home Assistant 容器时，需要清除哪些持久化数据目录？
# a
需要删除 `/mnt/config/.storage` 和整个 `/mnt/config` 目录，然后重新创建 `mkdir -p /mnt/config`。

