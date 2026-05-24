# q
如何初始化一个使用AWS SDK v2的Go模块？
# a
使用 `go mod init <模块名>` 创建模块，然后通过 `go get` 安装所需依赖。示例：
```bash
go mod init testproj
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
```

# q
使用AWS SDK v2操作S3服务需要安装哪些核心依赖包？
# a
需要安装以下三个包：
- `github.com/aws/aws-sdk-go-v2`
- `github.com/aws/aws-sdk-go-v2/config`
- `github.com/aws/aws-sdk-go-v2/service/s3`

安装命令：
```bash
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/s3
```

# q
如何永久设置Go模块代理为国内镜像（如 goproxy.cn）？
# a
执行以下命令：
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```
该命令会写入Go环境配置文件，永久生效。

