# q
如何使用 Python 的 configparser 模块修改 INI 格式的配置文件？
# a
使用 configparser.ConfigParser() 创建解析器，调用 read() 读取文件；检查目标 section 是否存在，若不存在用 add_section() 添加；使用 set(section, option, value) 设置参数值；最后用 open() 以写入模式打开文件，调用 config.write() 将修改写回。

# q
该 Python 脚本修改了 Ceph 配置文件的哪个参数，设置的值是什么？
# a
修改了 /etc/ceph/ceph.conf 文件中 [osd] 部分的 osd_op_num_shards 参数，将其值设置为 19。

# q
示例代码中如何避免因目标 section 缺失而导致修改失败？
# a
先用 config.has_section(section) 检查，若不存在则通过 config.add_section(section) 自动添加该 section。

