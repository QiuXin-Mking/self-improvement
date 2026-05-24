# q
SELinux是什么？
# a
SELinux（Security-Enhanced Linux）是一个Linux内核模块和安全子系统，用于提供访问控制安全策略。它采用类型强制（Type Enforcement）策略，为进程、文件、目录、端口等分配类型标签，通过比较标签来决定访问权限。SELinux默认启用并强制实施，其默认策略为目标策略（targeted policy），只保护选定的服务（如httpd、sshd等）。它提供三种运行模式：Enforcing（强制）、Permissive（宽容）和Disabled（关闭）。

# q
chrony服务是什么？
# a
chrony是一个开源自由的网络时间协议（NTP）客户端与服务器软件，用于保持计算机系统时钟与NTP服务器同步。它由两个程序组成：chronyd（后台守护进程，负责调整系统时钟，平滑修正时间偏差）和chronyc（用户界面，用于监控和配置）。chrony可以在间歇性网络、拥塞网络或温度变化等条件下良好运行，初始同步后不会停止时钟，避免影响依赖时间单调性的应用。

# q
SELinux有哪三种运行模式？
# a
1. Enforcing（强制模式）：SELinux安全策略被严格执行，违反策略的操作会被拒绝并记录。
2. Permissive（宽容模式）：SELinux会打印警告信息，但不拒绝任何操作，常用于排错。
3. Disabled（关闭模式）：SELinux完全不运行，无任何访问控制。

