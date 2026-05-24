# q
如何彻底重置 Home Assistant Docker 容器（包括清除旧配置）？
# a
```bash
sudo docker stop home-assistant
sudo docker rm home-assistant
rm -rf /mnt/config/.storage
rm -rf /mnt/config
mkdir -p /mnt/config
```
然后使用正确的 `docker run` 命令重新创建容器。

# q
创建 Home Assistant 容器的完整 `docker run` 命令是什么？
# a
```bash
sudo docker run -d \
  --name home-assistant \
  --restart=unless-stopped \
  -e TZ=Europe/London \
  -v /mnt/config:/config \
  -p 80:8123 \
  ghcr.io/home-assistant/home-assistant:stable
```

# q
文档中给出的 Home Assistant 本地访问地址、账号和密码是什么？
# a
- 域名：`http://www.homeassistant-qx.com`
- 账号：`qiuxin`
- 密码：`qiuxin@MK@159`

