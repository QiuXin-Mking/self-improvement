# q
goofys是什么？
# a
goofys是一个基于FUSE的S3文件系统客户端，可以将Amazon S3或兼容S3的对象存储桶挂载为本地文件系统。

# q
如何为goofys配置AWS凭证？
# a
可采用两种方式：
1. 创建`~/.aws/credentials`文件，内容为：
```
[default]
aws_access_key_id = YOUR_ACCESS_KEY_ID
aws_secret_access_key = YOUR_SECRET_ACCESS_KEY
```
挂载时通过`--profile`指定配置（如`--profile default`）。
2. 设置环境变量：
```
export AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_ACCESS_KEY
```

# q
goofys挂载S3存储桶的基本命令格式和常用选项是什么？
# a
基本格式：`./goofys [options] <bucket> <mountpoint>`
常用选项：
- `-f`：前台运行
- `-o allow_other`：允许其他用户访问挂载点
- `--endpoint=http://host:port`：指定自定义S3端点
- `--profile <name>`：使用指定的凭证配置
- `--debug_fuse`、`--debug_s3`：启用调试输出

示例：
```
./goofys -f -o allow_other --endpoint=http://127.0.0.1:5085 goofys-test /mnt/goofys-test
```

# q
如何在goofys中使用自定义S3兼容服务（非AWS）？
# a
通过`--endpoint`参数指定自定义S3服务的URL，同时确保凭证已正确配置。例如：
```
./goofys -o allow_other --endpoint=http://ees-southwest-8.edgeray.cn:5085 goofys-test /mnt/goofys-test
```

