# q
NTP（网络时间协议）的时间来源是什么？它是如何进行传播的？
# a
时间来源：国际标准时间UTC，由原子钟、天文台、卫星等提供。  
传播方式：采用分层传播机制，GPS获取标准时间后，传递给第1层NTP服务器，再由第1层传至第2层，逐层传送直至客户端。

# q
NTP服务的基本特点是什么？如何安装和管理chrony时间同步服务？
# a
NTP（网络时间协议）用于计算机时间同步，提供高精度校正，采用C/S模式，默认端口为UDP 123。  
在Linux中，常用chrony实现NTP客户端/服务器功能。安装：  
```bash
yum -y install chrony
```
启动：  
```bash
systemctl restart chronyd
```
设置开机自启：  
```bash
systemctl enable chronyd
```
查看服务状态：  
```bash
systemctl status chronyd
```
编辑配置文件：  
```bash
vi /etc/chrony.conf
```

# q
如何验证NTP服务是否正常工作？（可用哪些命令）
# a
- 查看服务端口监听：  
```bash
netstat -anptu | grep 123
```  
（如果监听不到，需启动服务）
- 查看chronyd服务状态和时间源：  
```bash
systemctl status chronyd
```
- 使用ntpdate测试与NTP服务器的通信：  
```bash
ntpdate -d <ntp_server_ip>
```  
若有来有回的响应包含 **ntp** 字样，则表示正常。

