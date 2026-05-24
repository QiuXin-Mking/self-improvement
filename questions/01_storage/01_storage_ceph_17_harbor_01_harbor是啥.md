# q
Harbor是什么
# a
Harbor 是一个开源的云原生容器镜像仓库，由 VMware 开发并捐献给了 CNCF。它用于存储和分发 Docker 容器镜像，提供企业级功能，如角色权限控制、镜像签名、漏洞扫描、多因子认证、Web UI 和 RESTful API 等。

# q
Harbor提供了哪些主要的安全性功能
# a
Harbor的安全性功能包括：
- 细粒度的角色权限控制（RBAC）
- 通过 Notary 实现镜像签名，确保镜像完整性和来源可信
- 集成 Clair、Trivy 等漏洞扫描工具，自动检测镜像安全漏洞

# q
Harbor支持哪些身份验证方式
# a
支持本地数据库认证、LDAP/AD 集成、OIDC 等多种身份验证方式，并可基于用户和组进行细粒度访问控制。

# q
如何使用 Docker Compose 部署 Harbor
# a
主要步骤：
1. 下载安装包并解压：
   ```sh
   wget https://github.com/goharbor/harbor/releases/download/v2.4.0/harbor-online-installer-v2.4.0.tgz
   tar xvf harbor-online-installer-v2.4.0.tgz
   cd harbor
   ```
2. 编辑 harbor.yml 配置文件
3. 执行安装脚本：
   ```sh
   ./install.sh
   ```

