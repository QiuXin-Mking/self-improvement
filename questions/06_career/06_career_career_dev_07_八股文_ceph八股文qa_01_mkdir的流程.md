# q
Lustre中mkdir操作在VFS层的入口是什么？
# a
sys_mkdir

# q
Lustre客户端处理mkdir调用的函数是哪个，位于哪个文件？
# a
ll_mkdir（位于 lustre/llite/namei.c:1891）

# q
在Lustre的客户端命名空间（LMV）层，负责选择目标MDT创建目录的函数是什么？
# a
lmv_create（位于 lustre/lmv/lmv_obd.c:1908）

# q
Lustre服务端接收到mkdir请求后，第一个处理函数是什么？
# a
mdt_create（位于 lustre/mdt/mdt_reint.c:498）

# q
Lustre mdd层完成目录创建操作的函数是哪个？
# a
mdd_create（位于 lustre/mdd/mdd_dir.c:2615）

