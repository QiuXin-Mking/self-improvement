# q
EES环境中通常涉及哪几个主要网络网段，各自的典型地址范围是什么？
# a
- external（外网）：例如 `10.3.0.41.0/24`
- internal（内网）：例如 `10.2.0.0/24`
- stor（存储网）：例如 `10.2.1.0/24`，通常比 internal 第三位 +1
- pxe（安装部署网）：以 `172.29.` 开头，例如 `172.29.0.0/16`
- ipmi（带外管理网）：以 `172.28.` 开头

# q
bond0 上的 monitor 和 lo 接口通常配置什么类型的地址？
# a
bond0 上的 monitor 和 lo 配置的是伪公网段地址（模拟公网使用的私有地址段）。

# q
EES 环境中核心功能 IP 的一般分配规律是什么？
# a
通常按顺序分配：网关 `.1`，monitor `.2`，evip `.3`，endpoint `.4`。  
如果该伪公网段有 IaaS 使用，则 IaaS 使用 `.1` 至 `.9`，monitor `.10`，evip `.11`，endpoint `.12`。  
OSPF 使用的就是 endpoint 的 IP。

# q
endpoint 的内网 IP 和外网 IP 分别对应哪个网段？
# a
- 内网 IP 使用 internal 网段地址
- 外网 IP 使用 external 网段地址  
以往有读写 IP 分离的场景，若无特殊要求则两者填写相同的 IP。

