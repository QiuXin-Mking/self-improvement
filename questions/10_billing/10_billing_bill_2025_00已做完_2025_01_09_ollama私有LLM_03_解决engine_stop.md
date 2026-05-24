# q
在 Windows 上部署 Ollama 时，如何检查系统的 Hyper‑V 虚拟化服务是否已开启？
# a
打开服务管理控制台（services.msc），确认与 Hyper‑V 相关的服务（如 Hyper‑V主机计算服务、Hyper‑V虚拟机管理等）处于“正在运行”状态。确保驱动已加载并正常运行。

# q
为何需要检查 bcdedit 中 hypervisorlaunchtype 的值？正确的设置是什么？
# a
因为 Ollama 需要 Hyper‑V 或 WSL2 的虚拟化支持，hypervisorlaunchtype 控制 Windows 启动时是否加载虚拟机监控程序。使用命令 `bcdedit` 查看该项，应设置为 `Auto`，否则虚拟化环境无法启动。

