# q
如何使用perf采集多个mdtio进程的调用栈性能数据？
# a
使用以下命令采集指定进程PID（例如56520-56525）的调用栈，持续30秒，保存到文件：
```bash
perf record -g -o mdtio_perf.data -p 56520 -p 56521 -p 56522 -p 56523 -p 56524 -p 56525 -- sleep 30
```
然后使用 `perf report -i mdtio_perf.data --comm=mdtio` 或 `perf report -i mdtio_perf.data` 查看火焰图或调用栈分析结果。

# q
如何一键统计所有mdt进程的CPU使用率之和？
# a
执行以下命令可计算mdt相关进程的CPU总占用百分比：
```bash
top -bn1 | grep mdt | awk 'BEGIN{s=0}{s+=$9}END{print s}'
```
该命令会输出所有mdt进程的%CPU列之和。

# q
在排查MDT高负载时，如何定位高频访问目录并量化其元数据压力？
# a
1. 确定高频目录（例如 `/mnt/users/8/5/1/0/85104905/studio_9219037/SHC1294_CulturalCenter/atc-server06/project/In_Progress`）  
2. 统计目录下文件数量：  
   ```bash
   find <目录> -type f | wc -l
   ```  
   例如该目录有38248个文件  
3. 统计目录总大小：  
   ```bash
   du -sh <目录>
   ```  
通过这些数据评估目录的元数据规模，辅助判断是否因大量小文件导致MDT压力过高。

