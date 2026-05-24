# q
frp 客户端配置文件中如何定义一个 TCP 代理，将本地 3389 端口映射到公网服务器的 5555 端口？
# a
在 `frpc.toml` 中添加以下配置：
```toml
serverAddr = "115.190.235.149"
serverPort = 9000

[[proxies]]
name = "test-tcp"
type = "tcp"
localIP = "127.0.0.1"
localPort = 3389
remotePort = 5555
```
其中 `serverAddr` 为 FRP 服务器 IP，`serverPort` 为服务器监听端口，`[[proxies]]` 数组定义具体代理规则。

# q
FRP 服务端最简配置需要设置什么参数？
# a
只需设置 `bindPort`，例如：
```toml
bindPort = 9000
```
该端口用于接收客户端连接。

# q
在容器中安装 SSH 服务并快速启用 root 密码登录需要执行哪些关键命令？
# a
```shell
apt-get update
apt-get install -y openssh-server
mkdir -p /var/run/sshd
echo "root:your_password" | chpasswd
sed -i 's/#PermitRootLogin.*/PermitRootLogin yes/' /etc/ssh/sshd_config
sed -i 's/#PasswordAuthentication.*/PasswordAuthentication yes/' /etc/ssh/sshd_config
```
如果需要允许 root 登录且使用密码认证，必须修改以上两个 SSH 配置项并重设 root 密码。

