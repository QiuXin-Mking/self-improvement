# q
这个毕设的核心目标是什么？
# a
搭建一个虚拟的小型边缘计算集群，不同容器负责不同功能；部署 K3s 或 MicroK8s 进行容器编排；使用性能测试工具评估并比较集群性能，找出稳定性最好的方案；最后应用具体案例验证性能测试的有效性。

# q
项目中使用的技术栈包含哪些关键组件？
# a
- 容器技术：Docker, containerd  
- 编排系统：K3s, MicroK8s  
- 虚拟平台：Proxmox VE  
- 监控工具：Prometheus, Grafana  
- 实验工具：iperf3, sysbench, stress-ng  
- 开发语言：Bash, Python, YAML（K8s 配置）

# q
根据“虚拟化与容器化在边缘计算应用部署中的比较研究”，容器化相比虚拟机的主要优势是什么？
# a
容器化在能源消耗方面更具生态优势，更轻量灵活，部署更快，资源利用率更高，尤其适合资源受限的边缘环境；但虚拟机在隔离性和安全性方面更强。

# q
在 CPU 和 I/O 磁盘性能测试中，Kubernetes 与 Docker Compose 的表现有何差异？
# a
- CPU 性能：长时间运行中 Kubernetes 表现更好，短时间内 Docker Compose 与 Podman 更优。  
- I/O 磁盘性能：1 分钟测试中 Docker Compose 更好，10 分钟测试中 Kubernetes 显著提升。  
- 内存性能：差异不显著，Kubernetes 因自身复杂性消耗更多资源。

