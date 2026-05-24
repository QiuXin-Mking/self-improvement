# q
在 CentOS 系统上安装 Docker Community Edition（CE）需要依次执行哪些命令？
# a
```bash
sudo yum update
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io
```

# q
如何用 Docker 在 host 网络模式下运行 Home Assistant，并挂载配置文件？
# a
```bash
sudo docker run -d --name home-assistant --restart=unless-stopped \
  -e TZ=YOUR_TIME_ZONE \
  -v /PATH_TO_YOUR_CONFIG:/config \
  --network=host \
  ghcr.io/home-assistant/home-assistant:stable
```

# q
将 Home Assistant 从 host 网络模式改为映射 80 端口，应该如何操作？
# a
先停止并删除原容器：
```bash
sudo docker stop home-assistant
sudo docker rm home-assistant
```
创建配置目录并启动新容器：
```bash
mkdir -p /mnt/config
sudo docker run -d --name home-assistant --restart=unless-stopped \
  -e TZ=Europe/London \
  -v /mnt/config:/config \
  -p 80:8123 \
  ghcr.io/home-assistant/home-assistant:stable
```

# q
在 CentOS 上通过 Python 虚拟环境安装 Home Assistant 需要哪些命令？
# a
```bash
sudo yum update
sudo yum install python3 python3-dev python3-venv python3-pip libffi-dev libssl-dev
mkdir homeassistant
cd homeassistant
python3 -m venv venv
source venv/bin/activate
```

