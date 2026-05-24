# q
如何运行NAS文件导入工具？
# a
- 进入工具目录：`cd /var/lib/sdsom/tools/nas-import`
- 使用指定Python解释器执行：`/var/lib/sdsom/venv/bin/python nas_import.py`

# q
导入模式有哪两种？如何通过交互选择？检查导入时需要额外提供什么路径？
# a
- 全新导入：运行时输入 `1`，表示所有文件全新导入。
- 检查导入（增量/检查）：输入除 `1` 以外的任意内容，然后会提示 `Please enter the check path:`，需要输入一个检查路径，且该路径必须是 NAS 挂载目录的子目录（例如 `/home/data`）。

# q
配置文件 `import.conf` 中各个参数的含义及其默认值是什么？
# a
配置文件位于 `/var/lib/sdsom/tools/nas-import/import.conf`，主要参数：
- `nas_path`：NAS 挂载路径（示例 `/nas_pool/nas01`）
- `bucket`：目标 RGW Bucket 名称
- `akey` / `skey`：访问密钥
- `dst_url`：RGW 服务的 Endpoint（示例 `10.10.8.11:8080`）
- `limit`：导入数据时每次提交的行数（示例配置为 `10`，注释说明默认值 `50`）
- `pool_size`：导入文件到 RGW 时的并发协程数（示例配置为 `50`，注释说明默认值 `2`）
- `seconds`：打印剩余文件数的时间间隔（秒）（示例配置为 `60`，注释说明默认值 `30`）

