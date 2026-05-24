# q
什么是三级域名？请举例说明其结构。
# a
三级域名位于二级域名之下，通常表示子域。例如 `blog.example.com` 中，`blog` 是三级域名，`example` 是二级域名，`.com` 是顶级域名。

# q
什么是四级域名？如何与泛域名结合使用？
# a
四级域名是三级域名之下的进一步细分，例如 `sales.blog.example.com` 中 `sales` 是四级域名。四级泛域名使用通配符匹配所有四级子域名，例如 `*.blog.example.com` 可以匹配 `sales.blog.example.com`、`news.blog.example.com` 等。

# q
泛域名的典型应用场景是什么？
# a
泛域名用于一次性匹配同一层级下的所有子域名，常见于SSL证书签发、简化DNS管理。例如 `*.example.com` 可匹配任意三级子域名，`*.blog.example.com` 可匹配任意四级子域名。

# q
配置四级泛域名时常用的DNS记录类型有哪些？请给出一个CNAME示例。
# a
常用记录：A记录（指向IPv4）、AAAA记录（指向IPv6）、CNAME记录（别名指向）。配置四级泛域名的CNAME示例：
```text
*.blog.example.com.  IN  CNAME  blog.example.com.
```

