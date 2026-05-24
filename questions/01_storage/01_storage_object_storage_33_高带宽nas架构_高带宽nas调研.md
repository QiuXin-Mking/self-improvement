# q
NAS（网络附加存储）是什么
# a
NAS（Network Attached Storage）是一种通过网络提供文件级存储访问的专用存储设备，允许不同客户端共享存储空间。它通常支持多种文件共享协议（如NFS、CIFS/SMB），使客户端像访问本地磁盘一样访问远程文件。

# q
CIFS 与 NFS 在 NAS 使用场景中的主要区别是什么
# a
CIFS（Common Internet File System）基于 SMB 协议，主要用于 Windows 环境下的网络文件共享和打印机共享；NFS（Network File System）则用于 Unix/Linux 系统间的文件分享。Windows 客户端建议使用 CIFS 协议，Linux 客户端建议使用 NFS 协议。

# q
Samba 是什么及其主要作用
# a
Samba 是一个开源的软件套件，实现了 SMB/CIFS 协议，允许 Unix/Linux 系统与 Windows 系统之间进行文件和打印机共享，同时也能充当 Windows 域控制器，提供身份验证和授权服务。

# q
Neutron 在 OpenStack 中的核心概念有哪些
# a
Neutron 是 OpenStack 的网络服务项目，核心概念包括：网络（逻辑隔离的 L2 网段）、子网（IP 地址段及网关/DNS 等配置）、端口（网络接口，可附加到虚拟机）、路由器（连接不同网络）、安全组（防火墙规则控制流量）、浮动 IP（动态分配的公共 IP 用于外部访问）。

