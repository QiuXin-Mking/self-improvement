# q
TCP/IP是什么？
# a
TCP/IP（传输控制协议/网间协议）是一个为广域网设计的工业标准协议集，由运输层、网络层、链路层组成。UDP（用户数据报协议）是TCP/IP协议族中的一种，与TCP相对应。

# q
Socket与TCP/IP协议族的关系是什么？
# a
Socket是应用层与TCP/IP协议族之间的中间软件抽象层，它隐藏了复杂的协议细节，对外提供一组简单接口。Socket相当于门面模式，让用户进程只需调用接口，由Socket内部组织数据以符合TCP/IP协议要求。

# q
在Socket通信中，TCP服务器端的核心流程是什么？
# a
TCP服务器端流程：创建套接字（socket_fd） → 绑定地址和端口（bind） → 启动监听（listen） → 接受客户端连接（accept），此时获得新的客户端套接字（client_fd） → 与客户端进行读写通信（read/write） → 通信结束后关闭套接字（close）。

# q
在Socket通信中，TCP客户端的核心流程是什么？
# a
TCP客户端流程：创建套接字（socket_fd） → 向服务器发起连接（connect） → 连接成功后进行读写通信（read/write） → 通信结束后关闭套接字（close）。客户端不需要绑定固定地址和端口，可以在任意能发起连接的地方使用。

