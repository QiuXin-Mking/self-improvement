# q
什么是S3预签名URL？它的核心作用是什么？
# a
预签名URL（Presigned URL）允许你生成一个临时URL，无需暴露AWS凭证，即可向他人授权上传文件到S3桶（或下载文件）。它包含了签名信息，在有效期内可被任何持有该URL的人使用。

# q
如何使用AWS CLI命令生成S3上传预签名URL？
# a
使用 `aws s3 presign` 命令，指定S3对象路径及有效期：
```sh
aws s3 presign s3://your-bucket-name/your-object-key --expires-in 3600
```
- `your-bucket-name` 为桶名
- `your-object-key` 为目标对象键（文件名）
- `--expires-in` 定义URL有效秒数（例中为1小时）

# q
生成预签名URL后，如何通过curl上传文件？
# a
使用 `PUT` 请求和 `--upload-file` 参数：
```sh
curl --request PUT --upload-file /path/to/your/file "https://your-bucket-name.s3.amazonaws.com/your-object-key?AWSAccessKeyId=...&Expires=...&Signature=..."
```
将本地文件路径替换为实际文件，URL替换为生成的完整预签名地址。

# q
`--expires-in` 参数在预签名URL中起什么作用？
# a
指定预签名URL的有效时长，单位为秒。过期后URL自动失效，无法继续使用。例如 `--expires-in 3600` 表示有效期为3600秒（1小时）。

# q
使用预签名URL上传文件时需要注意哪些关键点？
# a
- 预签名URL有效期由 `--expires-in` 决定，过期后需重新生成。
- 生成者必须具有对目标S3桶和对象的上传权限（PutObject）。
- URL中已嵌入临时签名，无需额外认证，因此需妥善保管，避免泄露。

