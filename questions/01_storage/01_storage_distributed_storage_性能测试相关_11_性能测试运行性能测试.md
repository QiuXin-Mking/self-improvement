# q
如何部署 VDBench 测试环境并启动后台性能测试？
# a
1. 在测试节点挂载 ISO 并安装 Java：`mount /dev/sr0 /mnt`，然后 `yum -y install java-1.8.0*`。
2. 进入 VDBench 目录，运行 `./vdbench -t` 验证环境。
3. 后台启动测试：`nohup ./vdbench -f <脚本文件> -o <输出目录> &`。
4. 查看实时输出：`tail -f nohup.out`。
5. 测试完成后卸载 ISO：`umount /mnt`。

# q
VDBench 配置脚本中，`hd`、`sd`、`wd` 参数分别代表什么？如何配置？
# a
- `hd`（Host Definition）：定义测试主机，默认参数如 `hd=default,vdbench=/root/vdbench50406,user=root,shell=ssh`，具体主机通过 `system` 指定 IP，如 `hd=hd1,system=172.22.101.189`。
- `sd`（Storage Definition）：定义存储设备，用 `lun` 指定块设备路径，如 `sd=sd01,hd=hd1,lun=/dev/sdk`，并通过 `threads` 指定并行线程数。
- `wd`（Workload Definition）：定义工作负载，`sd=sd*` 应用所有 sd，`rdpct=0` 表示纯写，`xfersize=1024k` 表示传输块大小。
- `rd`（Run Definition）：定义运行参数，如 `iorate=max` 最大速度，`elapsed=604800` 运行时长，`interval=5` 采样间隔。

# q
如何通过命令行管理 iSCSI 连接的建立与断开？
# a
- 发现并登录所有目标：`iscsiadm -m discovery -t st -p <target_IP> -l`
- 登出所有已连接的节点：`iscsiadm -m node -U all`
- 登出指定目标：`iscsiadm -m node -T <目标IQN> -u`

# q
性能测试中查看 CPU 和内存信息的常用命令有哪些？
# a
- 查看内存概况：`cat /proc/meminfo`
- 查看 CPU 型号：`cat /proc/cpuinfo | grep name | sort | uniq`
- 统计物理 CPU 个数：`cat /proc/cpuinfo | grep "physical id" | sort | uniq | wc -l`
- 查看每个 CPU 的核心数：`cat /proc/cpuinfo | grep "cpu cores" | uniq`
- 实时监控各 CPU 核心：运行 `top`，然后按 `1` 展开所有核心。

