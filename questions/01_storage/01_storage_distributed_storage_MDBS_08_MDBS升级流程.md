# q
MDBS升级过程中上传压缩包的命令是什么？
# a
使用 `upgrade mgt upload --file /root/mdbs18/upgrade.tar.gz` 将指定路径的升级包上传到目标环境。

# q
MDBS升级过程中有哪些核心命令行操作步骤？
# a
核心步骤包括：
1. 上传压缩包：`upgrade mgt upload --file /root/mdbs18/upgrade.tar.gz`
2. 检查压缩包：`upgrade mgt check`
3. 继续升级：`upgrade mgt continue_upgrade`
4. 查看升级日志：`tail -f /var/log/mdbs/ansible.log`

