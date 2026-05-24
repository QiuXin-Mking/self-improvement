# q
在 systemd 中，服务的 masked 状态是什么？
# a
masked 状态是 systemd 中服务的一种状态，通常由管理员使用 `systemctl mask` 命令手动设置。被 mask 的服务无法被启动（包括手动启动或被其他服务依赖启动），以确保它不会意外运行。

# q
如何将一个服务设置为 masked 状态？
# a
使用命令 `sudo systemctl mask <service_name>`，例如 `sudo systemctl mask httpd`。

# q
如何检查服务的当前状态（包括是否被 mask）？
# a
使用 `systemctl status <service_name>` 命令查看服务的状态，如果服务被 mask，输出中会显示 "masked"。

# q
如何解除服务的 masked 状态？
# a
使用命令 `sudo systemctl unmask <service_name>`，例如 `sudo systemctl unmask httpd`。

