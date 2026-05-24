# q
在 Dockerfile 中直接书写 ` /usr/sbin/sshd` 为什么会报错 "unknown instruction"？
# a
Dockerfile 中每一行必须是一个有效的 Docker 指令（如 `FROM`、`RUN`、`CMD` 等）。`/usr/sbin/sshd` 是一条普通的 shell 命令，不是 Docker 指令，因此构建时解析器无法识别，抛出 "unknown instruction" 错误。正确的做法是将命令放在 `RUN` 或 `CMD` 指令中，例如将启动命令写入脚本并通过 `CMD` 调用。

# q
在较新的 Ubuntu 镜像中使用 `pip3 install` 时出现 "externally-managed-environment" 错误，原因是什么？如何解决？
# a
原因：该 Ubuntu 版本遵循 PEP 668，标记系统 Python 环境为“外部管理”，禁止直接用 `pip` 修改系统级 Python 包，以免破坏系统依赖。  
解决方案有两种：
- 方案1：创建虚拟环境安装
  ```bash
  RUN python3 -m venv /opt/venv && \
      /opt/venv/bin/pip install --upgrade pip && \
      /opt/venv/bin/pip install jupyter
  ```
- 方案2：使用 `--break-system-packages` 标志（有风险）
  ```bash
  RUN pip3 install --break-system-packages jupyter
  ```

# q
在 Docker 中向容器传递 AWS 凭证等环境变量有哪些方式？
# a
主要有两种方式：
1. 在 `docker run` 时使用 `-e` 参数注入：
   ```bash
   docker run -e AWS_ACCESS_KEY_ID=xxx \
              -e AWS_SECRET_ACCESS_KEY=xxx \
              -e AWS_ENDPOINT=xxx \
              <image>
   ```
2. 在 Dockerfile 中使用 `ENV` 指令预设默认值：
   ```dockerfile
   ENV AWS_ACCESS_KEY_ID=your-access-key-id
   ENV AWS_SECRET_ACCESS_KEY=your-secret-access-key
   ENV AWS_ENDPOINT=https://your-endpoint-url
   ```
   运行时可通过 `-e` 覆盖这些值。

