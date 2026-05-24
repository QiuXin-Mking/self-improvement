# q
LDAP是什么？
# a
LDAP（轻型目录访问协议）是一种基于TCP的客户端/服务器协议，用于访问和管理以树状结构组织的目录信息（如用户、组、密码等）。它本质上是一个统一身份认证协议，具有优异的读性能但写性能较差，常见实现包括OpenLDAP、Microsoft Active Directory等。

# q
LDAP的核心作用是什么？
# a
LDAP的核心作用是实现统一身份认证，通过集中管理多个应用系统的用户账户和密码，让用户只需一套凭证即可登录不同系统，从而简化记忆和维护过程。

# q
Keystone是什么？
# a
Keystone是OpenStack项目中的组件，负责提供身份验证、授权和账户管理服务。它支持多种后端（如SQL和LDAP），可以集成LDAP将用户与组信息存放在外部目录服务器上。

# q
Keystone与LDAP集成时如何工作？
# a
集成时，用户与组的信息实际存储在LDAP服务器上，Keystone通过LDAP后端读取这些信息来创建令牌，从而允许用户访问OpenStack或其他云中的资源，Keystone本身不存储用户数据。

