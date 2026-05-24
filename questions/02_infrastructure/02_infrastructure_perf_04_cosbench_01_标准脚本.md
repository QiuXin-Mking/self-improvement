# q
COSBench 工作负载 XML 配置中包含哪些典型的 `workstage` 阶段？它们各自的作用是什么？
# a
- **init**：初始化阶段，创建测试所需的存储桶（bucket）。例中 `containers=r(1,2)` 表示创建两个桶，前缀为 `s3testqwer`。
- **prepare**：准备阶段，在已创建的桶中上传初始对象。例中上传 10 个 64KB 的对象到每个桶。
- **main**：核心测试阶段，执行实际的读写混合操作，通过 `runtime` 指定运行时长，`workers` 指定并发数，`operation` 及其 `ratio` 定义读写比例。
- **cleanup**：清理阶段，删除由 `prepare` 和 `main` 阶段创建的对象。
- **dispose**：清除阶段，删除由 `init` 创建的存储桶，恢复环境。

# q
在 COSBench 的 `main` 阶段中，如何定义多种操作及其比例？
# a
在 `<work name="main">` 内部使用多个 `<operation>` 子元素，通过 `type` 指定操作类型（如 `read`、`write`），`ratio` 属性指定该操作占总请求的百分比。示例中 `read` 占 80%，`write` 占 20%。配置中还可以用 `containers` 和 `objects` 指定操作的目标桶和对象范围，以及 `sizes` 指定写入对象的大小。

# q
COSBench 的 `storage` 元素如何配置 S3 兼容存储的连接和认证信息？
# a
通过 `type="s3"` 指定存储类型，并在 `config` 属性中以分号分隔的参数形式提供：`accesskey=<accesskey>;secretkey=<secretkey>;endpoint=<endpoint>;proxyhost=<proxyhost>;proxyport=<proxyport>`。实际使用时需替换为有效的密钥、端点及可选的代理设置。

# q
COSBench XML 中 `prepare` 阶段配置的 `sizes=c(64)KB` 表示什么？
# a
表示在准备阶段创建的对象大小固定为 64KB。`c(64)` 是 COSBench 的配置语法，`c` 代表常量（constant），括号内为具体数值和单位（支持 B、KB、MB、GB 等）。也可用 `u(min, max)` 表示随机大小。

