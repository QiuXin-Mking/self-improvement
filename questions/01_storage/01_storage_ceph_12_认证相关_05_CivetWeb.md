# q
CivetWeb 是什么？
# a
CivetWeb 是一个轻量级、嵌入式的 C 语言 HTTP 服务器库，设计用于嵌入式系统、IoT 设备等资源受限环境，可作为库嵌入应用程序提供 HTTP/HTTPS 服务，无需外部依赖。

# q
CivetWeb 包含哪些关键特性？
# a
- 轻量级、跨平台（Linux, Windows, macOS, QNX 等）
- 内置 HTTPS 支持（OpenSSL/WolfSSL）
- 多线程与异步 I/O，支持高并发
- 原生 WebSocket 支持
- 灵活的 API（自定义请求处理、文件服务、认证等）

# q
CivetWeb 在 Ceph RADOS Gateway 中的作用是什么？
# a
在 Ceph RADOS Gateway 中，CivetWeb 被用作内嵌的 HTTP 服务器，直接处理客户端的 HTTP/HTTPS 请求，提供对象存储网关的 Web 服务功能。

