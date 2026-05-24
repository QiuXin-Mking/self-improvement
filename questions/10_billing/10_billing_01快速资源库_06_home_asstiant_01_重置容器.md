# q
如何停止并删除现有的 Home Assistant 容器？
# a
```bash
sudo docker stop home-assistant
sudo docker rm home-assistant
```

# q
重置 Home Assistant 配置需要删除哪些目录？
# a
删除 `/mnt/config/.storage` 和整个 `/mnt/config`，然后重新创建：
```bash
rm -rf /mnt/config/.storage
rm -rf /mnt/config
mkdir -p /mnt/config
```

# q
如何使用 docker run 重新创建 Home Assistant 容器？
# a
```bash
sudo docker run -d --name home-assistant --restart=unless-stopped -e TZ=Europe/London -v /mnt/config:/config -p 80:8123 ghcr.io/home-assistant/home-assistant:stable
```

# q
访问 Home Assistant 的域名和登录凭据是什么？
# a
域名：`http://www.homeassistant-qx.com`  
账号：`qiuxin`  
密码：`qiuxin@MK@159`

