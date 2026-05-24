# q
NUMA架构是什么，如何理解它
# a
NUMA（Non-Uniform Memory Access，非一致内存访问）是一种多处理器架构。系统中每个处理器（或处理器组）有自己“本地”的内存，访问本地内存延迟低，访问其他处理器的“远程”内存延迟更高。它解决了SMP架构中所有CPU共享同一内存总线带来的扩展瓶颈问题。

# q
SMP（对称多处理）架构的核心特点是什么
# a
SMP（Symmetric Multi-Processing）架构中，所有CPU共享同一条总线和同一块物理内存，无主从之分，每个CPU对内存的访问路径和延迟一致，因此也被称为UMA（Uniform Memory Access，一致内存访问）。工作负载可均匀分配到所有处理器，但总线带宽会成为扩展瓶颈。

# q
NUMA与SMP（UMA）的关键区别是什么
# a
核心区别在于内存访问延迟：SMP（UMA）中所有处理器访问内存的延迟均等，共享单一内存总线；NUMA中处理器访问本地内存快，访问远程内存慢，内存物理分布在不同节点上，可扩展性更好，但需要软件感知NUMA拓扑才能获得最佳性能。

# q
如何在Linux系统上手动控制NUMA内存分配和CPU亲和性
# a
使用`numactl`工具。示例：
```bash
numactl --cpunodebind=0 --membind=0 myapp
```
该命令将进程绑定到NUMA节点0的CPU上运行，并强制只从节点0分配内存，从而避免跨节点远程访问带来的性能下降。

# q
AMD处理器在实现NUMA架构时使用了哪些关键技术
# a
AMD采用多芯片模块（MCM）封装，并通过Infinity Fabric互联技术将不同的处理器模块和内存控制器连接在一起，这在EPYC系列处理器中尤为典型，实现高效的NUMA拓扑。

