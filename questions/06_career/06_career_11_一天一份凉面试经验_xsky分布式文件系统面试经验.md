# q
如何实现 C++ 中的 shared_ptr？
# a
shared_ptr 的核心是通过引用计数管理资源生命周期：  
- 内部通常包含两个指针：一个指向被管理的对象，一个指向控制块（包含引用计数和弱引用计数）。  
- 构造函数将计数初始化为 1，拷贝构造函数递增计数，析构函数递减计数，当计数归零时释放对象。  
- 需要重载 `*`、`->`、`operator bool` 等操作，并正确处理移动语义和线程安全。  
示例骨架：  
```cpp
template<typename T>
class shared_ptr {
    T* ptr;
    int* ref_count;
public:
    explicit shared_ptr(T* p = nullptr) : ptr(p), ref_count(new int(1)) {}
    shared_ptr(const shared_ptr& other) : ptr(other.ptr), ref_count(other.ref_count) {
        ++(*ref_count);
    }
    ~shared_ptr() {
        if (--(*ref_count) == 0) {
            delete ptr;
            delete ref_count;
        }
    }
    T& operator*() { return *ptr; }
    T* operator->() { return ptr; }
    // … 省略赋值、移动等
};
```

# q
桶排序的基本思想与实现步骤是什么？
# a
桶排序将数据分到多个有序的桶中，每个桶内单独排序，最后按顺序合并。  
步骤：  
1. 确定桶的数量（如数据范围 / 桶大小）。  
2. 遍历数据，将每个元素分配到对应的桶中（映射函数）。  
3. 对每个桶内的元素进行排序（常用插入排序或递归使用桶排序）。  
4. 按桶的顺序依次输出所有元素。  
适用于数据均匀分布的场景，时间复杂度在理想情况下可达 O(n)。

# q
在 Lustre 文件系统中，一个客户端持有文件句柄不释放，另一个客户端写入该文件，之前打开的客户端能否读到新写入的内容？
# a
能读到，但取决于锁机制与一致性模型。Lustre 采用分布式锁管理缓存的一致性，支持 POSIX 强一致性，通过 LDLM（Lustre Distributed Lock Manager）提供字节粒度锁。  
写入客户端在修改前会获取相应字节范围的写锁，并撤销其他客户端的读锁，导致持有文件句柄的客户端缓存失效。再次读取时，将重新向元数据服务器（MDS）或对象存储服务器（OSS）请求锁与最新数据，因此能看到其他客户端写入后的内容。如果字节粒度锁实现有误或配置为弱一致性，则可能读到旧数据。

# q
什么是弱存储系统？简单介绍其特性。
# a
弱存储系统通常指放宽一致性要求、不保证强一致性的存储架构。典型特征包括：  
- 允许数据在节点间的复制存在延迟，牺牲部分一致性以换取高可用和分区容忍性（CAP 中的 AP 倾向）。  
- 读写不必立刻反映在所有副本上，可能读到过时数据（最终一致性）。  
- 常采用 Quorum 机制、向量时钟、反熵修复等方式管理分歧与同步。  
与传统 POSIX 强一致系统（如 Lustre）不同，它适用于大规模、多数据中心的场景，如 Amazon Dynamo、Cassandra 等。

