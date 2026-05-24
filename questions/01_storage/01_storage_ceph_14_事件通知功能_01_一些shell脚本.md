# q
如何用Shell脚本计算AWS S3 V2签名的signature？
# a
先将签名要素拼接成`stringToSign`，格式为：`HTTPVerb\n\nContentType\nDate\nResource`。然后使用openssl命令生成签名：
```bash
signature=`echo -en ${stringToSign} | openssl sha1 -hmac ${s3Secret} -binary | base64`
```
其中`s3Secret`为密钥，最终`signature`为Base64编码的结果。

# q
S3分段上传包含哪几个核心步骤？
# a
1. 创建分段上传：发送`POST`到资源URL并带`?uploads`参数，从响应中获取`UploadId`。
2. 上传分段：发送`PUT`到资源URL，必须携带`?partNumber=`和`&uploadId=`参数，每个分段上传后记录返回的`ETag`。
3. 完成分段上传：发送`POST`到资源URL并带`?uploadId=`，请求体为XML格式的`CompleteMultipartUpload`，列出所有分段的`ETag`和对应的`PartNumber`。

# q
完成分段上传的XML请求体是什么结构？
# a
结构为`<CompleteMultipartUpload>`根元素，内部包含多个`<Part>`子元素，每个`<Part>`包含`<ETag>`和`<PartNumber>`。示例：
```xml
<CompleteMultipartUpload>
  <Part>
    <ETag>"etag值"</ETag>
    <PartNumber>1</PartNumber>
  </Part>
  ...
</CompleteMultipartUpload>
```

