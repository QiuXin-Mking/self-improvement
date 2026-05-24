# q
Optimus是什么，它主要解决什么问题？
# a
Optimus是一个PB级对象存储数据迁移平台，用于将用户原有对象存储（如Ceph、Minio、OpenStack Swift等）中的数据快速、稳定地迁移到新的分布式存储平台，数据规模通常在几百TB以上。它基于全网络、存储平台无关，无需客户端安装，全部通过自动化网络实现迁移，并提供权限认证、Restful API、目录结构一致性保障等功能。

# q
Optimus主动迁移系统的分布式架构是如何实现的？
# a
Optimus主动迁移系统基于Mesos资源调度框架实现。Optimus本身是一个Mesos Framework（Scheduler），它接收Mesos Master的资源Offer，根据资源情况将迁移Job拆分为多个Task，并将Task描述返回给Mesos Master，由Master分配到各个Mesos Agent上执行。Agent上的Executor负责根据URL下载文件并上传到目标S3存储，迁移结果和状态记录在MySQL数据库中。

# q
Optimus如何实现高可用和弹性收缩？
# a
- **弹性收缩**：通过往Mesos集群中添加或移除Agent节点即可弹性伸缩，Agent运行在sandbox中，任务完成后自动清理。没有迁移任务时，Mesos集群可供给其他Framework（如Marathon）使用。
- **高可用**：Mesos Master本身支持高可用。Optimus Framework是无状态程序，核心状态存储在Mesos，Job状态存储在MySQL（使用主备架构），因此可将Optimus放在Marathon、Kubernetes或Pacemaker上实现自动故障恢复。

# q
Optimus自动回源系统方案一的核心工作流程是什么？
# a
方案一使用Openresty作为七层代理，Lua代码检查S3请求的HTTP返回码。若为404，则查询数据库获得bucket与源站的映射，重写URL后从源站下载文件。同时启动一个coroutines，将下载的文件上传回S3存储作为镜像，并将临时文件返回给用户，从而透明实现一次请求完成回源和镜像，后续相同文件不再回源。

# q
为什么Optimus选择Golang开发Mesos Executor和Framework？
# a
主要原因包括：
- Golang静态编译，可打包为单一二进制文件，部署和升级极为方便，无需像Python那样安装依赖和运行预装脚本。
- 可跨CentOS 6、CentOS 7和Ubuntu等不同系统稳定运行，依赖简单。
- 相比Docker镜像，Golang二进制更小。
- 项目为保持一致性，Framework也使用Golang，实践证明Mesos的Golang库在生产环境运行稳定。

