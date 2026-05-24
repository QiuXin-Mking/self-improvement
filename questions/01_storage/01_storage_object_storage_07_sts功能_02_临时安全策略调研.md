# q
Ceph Object Gateway中STS API与S3 API的共存方式是什么？
# a
默认情况下，STS和S3 API共存于同一命名空间，两者可以通过Ceph Object Gateway的同一端点访问。

# q
如何理解AssumeRole在临时凭证体系中的作用？
# a
AssumeRole允许一个AWS实体（如用户、角色或服务）临时获取另一IAM角色的权限。成功承担角色后会获得临时凭证（访问密钥ID、秘密访问密钥、会话令牌），凭证具有目标角色的权限并存在有效期。信任策略以JSON文档定义哪些实体可承担该角色，常用于跨账户访问、服务角色和联合身份。它由AWS STS服务的AssumeRole操作实现。

# q
Ceph Object Gateway中STS临时凭证的核心API有哪些？
# a
核心API包括GetSessionToken（获取与原始AWS凭证权限相同的临时凭证）、AssumeRole（通过信任策略获取指定角色的临时权限）等。Ceph的STSLite实现支持这些STS API。

# q
IAM角色与IAM用户的主要区别是什么？
# a
IAM角色是一种虚拟用户，没有永久身份凭证（登录密码或访问密钥），需要被可信实体扮演后才能获得临时安全令牌（STS Token）。IAM用户则拥有永久凭证。角色适用于临时授权、跨账号访问和联合身份场景。

# q
Ceph中如何通过radosgw-admin管理IAM角色和策略？
# a
使用`radosgw-admin role create`创建角色，`radosgw-admin role list`列出角色，`radosgw-admin role-policy put`为角色设置权限策略，`role-policy get`获取策略。角色通过RoleName、RoleId和Arn标识，策略绑定到角色上。

