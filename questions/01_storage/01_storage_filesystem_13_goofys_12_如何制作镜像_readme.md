# q
如何在镜像中安装 goofys 并复制到系统路径？
# a
使用 wget 下载二进制文件并复制到 `/usr/local/bin/`：
```
wget -v https://github.com/kahing/goofys/releases/download/v0.24.0/goofys
./goofys /usr/local/bin/goofys
```

# q
如何配置 SSH 服务并允许 root 用户通过密码登录？
# a
安装 openssh-server 并修改配置文件：
```bash
apt-get install -y openssh-server && mkdir -p /run/sshd
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config
```

# q
如何安装 Python 3.11 并将其设置为默认 Python3 解释器？
# a
通过 deadsnakes PPA 安装，并使用 update-alternatives 设置默认版本：
```bash
apt-get install -y python3 python3-pip software-properties-common
pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/
add-apt-repository ppa:deadsnakes/ppa
apt-get update
apt-get install -y python3.11 python3.11-distutils
update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.11 1
update-alternatives --set python3 /usr/bin/python3.11
ln -s /usr/bin/python3 /usr/bin/python
curl -sS https://bootstrap.pypa.io/get-pip.py | python
sed -i 's@20.0.2@25.0.1@g' /usr/bin/pip*
```

