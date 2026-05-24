# q
如何使用 SSH 在本地创建一个 SOCKS5 代理端口，将流量通过远程服务器转发出去？
# a
使用 `ssh -D <本地端口> <user>@<远程服务器IP>` 命令。例如：
```bash
ssh -D 1080 ubuntu@152.32.189.229
```
这会在本地开启一个 SOCKS5 代理（端口 1080），本机流量将被加密转发到远程服务器，再由远程服务器访问公网。

# q
如何让当前终端会话的 HTTP/HTTPS 请求通过本地的 SOCKS5 代理？
# a
设置环境变量，使用 `socks5h://` 协议（`h` 表示代理服务器会解析域名）：
```bash
export http_proxy="socks5h://127.0.0.1:1080"
export https_proxy="socks5h://127.0.0.1:1080"
```

# q
如何让 Git 通过本地的 SOCKS5 代理访问远程仓库？
# a
配置 Git 全局代理，使用 `socks5://` 协议：
```bash
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```

