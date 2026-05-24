# q
OCACHE读提升(READ LIFT)机制的核心概念是什么？
# a
读提升机制指当IO大小小于10MB时，在从HDD读取数据的同时，将数据写入SSD cache，以提升后续读取性能。具体步骤：1. 从HDD读取数据；2. 将数据写入SSD。

# q
ocache_model_need_rl判断读提升的条件是什么？
# a
通过ocache_read_cmd_type_judge检查IO大小，若所有IO大小均小于10MB，则ocache_model_need_rl判定需要执行读提升。

# q
在OCACHE读写路径中，数据是如何流经各层的？
# a
IO请求依次经过sio → osd → ocache，最后从ocache访问hdd zdev或cdisk zdev（SSD cache）。刷盘时，首先向内存申请cache资源，然后数据通过cdisk zdev读到cache，再从cache写入hdd。

