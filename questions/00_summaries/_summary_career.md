# q
LSM Tree的核心设计思想是什么？
# a
将随机写转化为顺序写，通过WAL保证持久性，MemTable作为内存缓存，达到阈值后批量顺序写入磁盘形成SSTable，并通过多级合并（Compaction）整理数据，减少写放大。适用于写多读少的场景，用于LevelDB、RocksDB、TiKV等系统。

# q
2Q缓存替换算法的结构和主要改进是什么？
# a
2Q使用两级缓存FIFO+LRU，改进为三段式：A1in（接收新缓存项的FIFO）、A1out（记录被淘汰键的FIFO）、Am（主LRU缓存）。数据首次进入A1in，若在A1out中存在则升级到Am，避免低频数据污染缓存。

# q
Claude Code的Skill定义文件SKILL.md应包含哪些关键部分？
# a
SKILL.md包含name（技能名称）、description（用于匹配用户指令的描述）、具体指令（定义技能行为）和资源区（如引用文件）。存放位置为项目级`./claude/skills/`或全局`~/.claude/skills/`。

# q
在Docker中运行Vue前端开发环境时，如何解决Vite服务监听localhost导致宿主机无法访问的问题？
# a
启动Vite开发服务器时添加`--host 0.0.0.0`参数，将其绑定到所有网络接口。同时建议配置淘宝npm镜像源加速依赖安装：`npm config set registry https://registry.npm.taobao.org`。

# q
跳表（Skip List）在哪些存储系统的哪个部分被使用？
# a
跳表用于Redis的Sorted Set有序集合和LevelDB/RocksDB的MemTable内存表，提供平均O(log n)的插入、删除、查找性能，实现简单且支持范围查询。

