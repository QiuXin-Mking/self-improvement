# q
如何使用OpenStack命令创建对象存储用户并生成EC2凭证？
# a
步骤：
1. 创建租户：`openstack project create Iaas_project`
2. 创建用户：`openstack user create --project Iaas_project --password mypassword Iaas_user`
3. 分配角色：`openstack role add --project Iaas_project --user Iaas_user _member_`
4. 创建EC2凭证：`openstack ec2 credentials create --user Iaas_user --project Iaas_project`
完成后可通过 `/api/object_storage/users` 接口管理对象用户。

# q
对象存储用户管理API的代码路径及功能是什么？
# a
代码位于 `ees_manager/api_gateway/resources/object_storage/`，提供对象用户的获取、创建、删除功能。

# q
排查ESC API创建对象用户时，需要查阅哪些内部文档？
# a
- 对象用户API文档：https://wiki.haplat.net/pages/viewpage.action?pageId=282651216
- 爱捷云API文档：https://wiki.haplat.net/pages/viewpage.action?pageId=183905358
同时需要加载 `keystonerc_admin` 环境变量以获取管理员权限。

