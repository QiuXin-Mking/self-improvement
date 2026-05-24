# q
如何从审计日志定位S3 list-objects操作性能异常？
# a
分析审计日志中 GET /{bucket} 请求的总耗时字段。日志格式中第12列（`"0.316006534s"`）为总响应时间，最后一列为实际处理时间。正常请求耗时在 **0.01~0.02s** 左右，异常请求可达 **3.9s**（如 `"3.904081087s"`），说明存在性能瓶颈或网络延迟。日志示例：

```
2024-12-09T10:59:31.354+0800  0 Audit:  "binbash" "-" "GET /geneway-data 1.1" ... "200" "0" "GET" "10.176.177.2" "aws-cli/..." "0" "3.904081087s" "0.000000" "0.000006" "0.000003" ... "3.899813" "10.176.177.31"
```

所有请求返回 200 说明功能正常，问题集中在响应时间波动。

# q
S3API list-objects命令中`--max-keys 0`参数的作用是什么？
# a
`--max-keys 0` 用于**测试连接延迟而不实际返回对象列表**。该参数将返回的对象数量限制为0，避免传输大量数据干扰性能测试。常用于排查对象存储服务连接性能问题：

```bash
aws s3api list-objects --bucket geneway-data --endpoint-url http://10.176.177.31:5085 --max-keys 0
```

结合 `time` 命令可精确测量端到端耗时：
```bash
time (aws s3api list-objects --bucket geneway-data --endpoint-url http://10.176.177.31:5085 --max-keys 0)
```

# q
弹性计算虚机跑不满问题中S3 list操作的标准排查流程是什么？
# a
1. **验证基础连通性**：使用 `--max-keys 0` 排除数据传输干扰，测试不同 endpoint 的响应时间：
   ```bash
   aws s3api list-objects --bucket geneway-data --endpoint-url http://10.176.177.31:5085 --max-keys 0
   aws s3api list-objects --bucket geneway-data --endpoint-url http://10.176.177.30:5085 --max-keys 0
   ```

2. **记录完整列表耗时**：执行无限制的 list 操作并用 `date` 标记起止时间：
   ```bash
   date ; aws s3api list-objects --bucket geneway-data --endpoint-url http://10.176.177.31:5085 ; date
   ```

3. **分析审计日志**：检查 `GET /{bucket}` 请求的总耗时字段（第12列），对比正常基线（~0.02s）与异常值（>3s），判断是否存在间歇性性能退化。

4. **排查方向**：若端点延迟正常但虚机CPU跑不满，问题可能在应用层并发模型或I/O等待，而非对象存储本身。

