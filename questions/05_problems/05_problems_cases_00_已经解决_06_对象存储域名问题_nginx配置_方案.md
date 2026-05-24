# q
如何通过 Nginx 反向代理将子域名请求转发到另一个域名的特定路径前缀（如 `/okok`）？
# a
可以使用 Nginx 的 `rewrite` 结合 `proxy_pass`。实现从 `http://okok.ees-southwest-10.edgeray.cn:5085` 到 `http://ees-southwest-10.edgeray.cn:5085/okok` 的转发配置如下：
```nginx
server {
    listen 5085;
    server_name okok.ees-southwest-10.edgeray.cn;

    location / {
        rewrite ^(.*)$ /okok$1 break;
        proxy_pass http://ees-southwest-10.edgeray.cn:5085;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```
该配置中，`rewrite … break` 将请求的原始 URI 前加上 `/okok`，并停止后续重写，立即转发修改后的 URL 到后端。

# q
在该 Nginx 转发方案中，通过哪些 `proxy_set_header` 指令向后端传递原始请求信息？
# a
配置使用以下头部传递客户端信息：
- `proxy_set_header Host $host;` — 传递客户端请求的 `Host` 头
- `proxy_set_header X-Real-IP $remote_addr;` — 传递客户端真实 IP
- `proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;` — 追加代理链 IP
- `proxy_set_header X-Forwarded-Proto $scheme;` — 传递原始请求协议（http/https）

