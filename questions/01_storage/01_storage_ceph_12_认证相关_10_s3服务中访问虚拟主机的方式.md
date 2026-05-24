# q
什么是S3桶访问的虚拟主机风格（Virtual-Host Style）？
# a
虚拟主机风格是一种S3桶访问方式，将桶名称作为URL子域的一部分，格式如 `http://bucket-name.example.com/object-key`。该方式依赖DNS配置将请求路由到正确的桶，使URL更简洁易读。

# q
什么是S3桶访问的路径风格（Path-Style）？
# a
路径风格将桶名称放在URL路径中，格式如 `http://example.com/bucket-name/object-key`。这是另一种常见的S3桶访问方式。

# q
Ceph RADOS Gateway中`rgw_dns_name`配置项的作用是什么？
# a
`rgw_dns_name`用于设置RADOS Gateway的基础DNS名称（即服务的主域名），从而支持虚拟主机风格的S3桶访问。正确配置后，客户端可以通过 `bucket-name.<rgw_dns_name>` 的形式访问桶。

