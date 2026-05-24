# q
How to install Docker CE on CentOS for Home Assistant?
# a
```
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install -y docker-ce docker-ce-cli containerd.io
```

# q
What Docker command starts Home Assistant with host networking and a configuration directory?
# a
```
sudo docker run -d --name home-assistant --restart=unless-stopped -e TZ=YOUR_TIME_ZONE -v /PATH_TO_YOUR_CONFIG:/config --network=host ghcr.io/home-assistant/home-assistant:stable
```

# q
How to change the Home Assistant container’s host port from 8123 to 80 using Docker?
# a
Stop and remove the existing container, then run a new one with the desired port mapping and a persistent config directory:
```
sudo docker stop home-assistant
sudo docker rm home-assistant
mkdir -p /mnt/config
sudo docker run -d --name home-assistant --restart=unless-stopped -e TZ=Europe/London -v /mnt/config:/config -p 80:8123 ghcr.io/home-assistant/home-assistant:stable
```

