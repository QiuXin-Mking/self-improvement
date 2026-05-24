# q
磁盘异常数量问题的典型根因是什么？
# a
通常根因是RAID卡通过`arcconf`工具获取的物理设备列表与系统中WWID符号链接（如`/usr/lib/python3.6/site-packages/ees_manager/api_gateway/resources`下的wwn-*到`sda`等设备的映射）不一致，或中间组件（`es_agw`, `nd_agent`, `hard_api.py`）在传递磁盘数据时发生错误，导致统计异常。

# q
如何定位磁盘异常数量问题中磁盘WWID与物理设备不匹配的情况？
# a
1. 检查WWID符号链接目录`/usr/lib/python3.6/site-packages/ees_manager/api_gateway/resources`，确认每个`wwn-*`是否正确指向对应的`sdX`设备。
2. 使用`arcconf`工具获取物理设备列表，对比`get_pysical_devices`函数（参考`01_问题单/25_磁盘异常数量/get_py_de.py`）的输出。
3. 沿调用链`ArcconfCliTool -> hard_api.py -> nd_agent -> es_agw`检查日志，确认各环节获取的磁盘数据是否一致。

