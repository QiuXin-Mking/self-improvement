# q
在 Home Assistant Docker 环境中，如何将 HACS 安装到自定义组件目录？
# a
1. 确认容器 ID：`docker ps`，示例中容器 ID 为 `19ff9bb1626d`。
2. 查看配置目录挂载路径：`docker inspect 19ff9bb1626d`，从 `"Binds"` 字段找到对应关系，例如宿主机 `/mnt/config` 挂载到容器内 `/config`。
3. 进入 custom_components 目录：
   ```bash
   cd /mnt/config/custom_components
   ```
4. 下载 HACS zip 包：
   ```bash
   wget https://github.com/hacs/integration/releases/download/2.0.5/hacs.zip
   ```
5. 解压到 hacs 子目录：
   ```bash
   unzip hacs.zip -d hacs
   ```
6. 重启容器使生效：
   ```bash
   docker restart 19
   ```

# q
Home Assistant 的 Docker 容器通常将配置文件目录挂载在宿主机什么位置？如何确认？
# a
具体位置由 Docker 运行时的挂载参数决定。查看方法：
```bash
docker inspect <容器ID>
```
在输出中的 `"Binds"` 字段会显示宿主机目录和容器内 `/config` 的映射关系，例如：
```
"/mnt/config:/config"
```
这样宿主机上的 `/mnt/config` 就是配置目录。

# q
在 Home Assistant 配置目录中，自定义集成需要放在哪个子目录里？
# a
需要放在配置目录下的 `custom_components` 子目录中。如果该目录不存在，需要手动创建：
```bash
mkdir /mnt/config/custom_components
```

