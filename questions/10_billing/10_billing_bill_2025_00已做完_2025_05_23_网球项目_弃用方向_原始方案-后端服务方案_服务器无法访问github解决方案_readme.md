# q
如何通过 SSH 建立本地 SOCKS5 代理隧道，将本地的网络流量转发到能访问公网的远程服务器？
# a
使用 `ssh -D` 命令在本地监听一个 SOCKS5 端口（如 1080），将流量通过远程服务器转发：
```bash
ssh -D 1080 ubuntu@152.32.189.229
```
执行后本地的 1080 端口即成为一个 SOCKS5 代理，所有发往该端口的流量都会经远程服务器出公网。

# q
如何让当前终端的所有 HTTP/HTTPS 请求走方才建立的 SOCKS5 代理？
# a
设置环境变量，并使用 `socks5h` 让 DNS 解析也通过代理：
```bash
export http_proxy="socks5h://127.0.0.1:1080"
export https_proxy="socks5h://127.0.0.1:1080"
```
需要注意 `socks5h` 与 `socks5` 的区别：`socks5h` 会将域名解析也交给代理服务器完成。

# q
如何配置 Git 使其克隆、拉取等操作通过指定的 SOCKS5 代理？
# a
使用以下全局配置命令：
```bash
git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'
```
配置后 Git 的所有 HTTP/HTTPS 通信都会经过本地 1080 端口的 SOCKS5 代理。

