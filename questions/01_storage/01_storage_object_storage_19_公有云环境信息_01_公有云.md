# q
如何使用curl测试公有云S3兼容对象存储的连通性？
# a
使用GET请求调用数据迁移服务的S3检查接口，需传递AccessKey、SecretKey、Endpoint、Region等参数。示例命令：
```bash
curl -g -i -X GET 'http://<host>:5001/api/data_migrate/s3check?access_key=<AccessKey>&endpoint=https%3A%2F%2Fs3-<region>.wcsapi.com&secret_access_key=<SecretKey>&target_or_source=source&region_name=<region>'
```
该接口用于验证S3服务的可达性与凭证有效性。

# q
使用aws CLI连接公有云S3存储时，区域（region）配置错误会导致什么错误？如何正确配置？
# a
若`aws configure`中设置的`Default region name`与Endpoint中实际的region不一致（例如配置为`us-east-1`但端点指向`cn-south-1`），执行操作时会报错：
```
An error occurred (AuthorizationHeaderMalformed) ... the region 'us-east-1' is wrong; expecting 'cn-south-1'.
```
正确做法：在执行`aws configure`时，`Default region name`必须与Endpoint中的region一致。如Endpoint为`http://s3-cn-south-1.wcsapi.com`，region应设置为`cn-south-1`。

# q
公有云对象存储的控制台登录通常通过什么URL访问？
# a
通过特定格式的Web控制台URL访问，典型格式为：
```
https://<storage-host>:9002/auth/login?service=<url-encoded-service-url>
```
例如：
```
https://storage-szys01.wangsucloud.com:9002/auth/login?service=https:%2F%2Fstorage-szys01.wangsucloud.com%2Fauth%2Flogin%2F
```
登录时需使用有效的管理员账号和密码。

