# q
ESC API 中 CreateObjectUser 接口返回什么内容？如何获取对象用户的 AccessKey？
# a
CreateObjectUser 使用 POST 方法，原生接口只返回 user id。获取 AccessKey 需要单独调用 DescribeAccessKey（GET 方法），通过 user id 查询对应的秘钥。

# q
ESC API 中对象用户及相关 AccessKey 接口各自限定的 HTTP 方法是什么？
# a
- CreateObjectUser: POST
- DescribeObjectUsers: GET
- UpdateObjectUserName: PUT
- DeleteObjectUser: DELETE
- UpdateObjectUserStatus: PUT
- CreateAccessKey: POST
- DeleteAccessKey: DELETE
- DescribeAccessKey: GET
- UpdateAccessKeyStatus: PUT

