# q
如何根据容器ID或名称获取已运行容器的docker run命令？
# a
使用提供的脚本，传入容器ID或名称作为参数：
```bash
./get_docker_run_command.sh <container_id_or_name>
```
脚本利用 `docker inspect` 获取容器详细信息，通过 `jq` 提取镜像、环境变量、端口和卷挂载，并组装成完整的 `docker run -d` 命令。

# q
在脚本中，如何使用jq提取容器的环境变量并格式化为 `-e` 参数？
# a
通过以下jq命令将环境变量数组用 ` -e ` 连接：
```bash
echo "$INSPECT" | jq -r '.[0].Config.Env | join(" -e ")'
```

# q
脚本中如何提取容器暴露的端口映射并转换为 `-p` 参数？
# a
使用下列管道命令：
```bash
echo "$INSPECT" | jq -r '.[0].NetworkSettings.Ports | to_entries[] | "\(.value[0].HostPort):\(.key)"' | sed 's/\/tcp//g' | sed 's/\/udp//g' | xargs
```
该命令解析端口JSON，生成 `HostPort:ContainerPort` 格式，去除协议后缀，并用空格连接以便作为 `-p` 的参数。

# q
脚本中如何提取容器的卷挂载并格式化为 `-v` 参数？
# a
通过以下jq命令：
```bash
echo "$INSPECT" | jq -r '.[0].Mounts | map("\(.Source):\(.Destination)") | join(" -v ")'
```
该命令将每个挂载点转换为 `Source:Destination` 并用 ` -v ` 连接。

