# q
lustre 查询网络状态如何查询,命令是啥
# a
lnetctl stats show

# q
lnetctl stats show 命令常用的调用关系结构是怎么样的
# a
1. lnetctl 命令解析
2. jt_show_stats() - 命令处理函数
3. lustre_lnet_show_stats() - 库函数封装，ioctl 系统调用
4. LNet ioctl 处理入口
5. lnet_counters_get() - 收集统计计数器
6. lnet_counters_get_common_locked() - 收集通用计数器

