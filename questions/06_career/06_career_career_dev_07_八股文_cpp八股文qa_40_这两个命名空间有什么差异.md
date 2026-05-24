# q
Ceph RGW 中 `RGW_OBJ_NS_MULTIPART` 和 `RGW_OBJ_NS_SHADOW` 这两个命名空间主要差异是什么？
# a
这两个常量定义的是 RGW 在桶索引里使用的**对象命名空间**，用来隔离特殊用途对象和对用户透明的数据。  
- `RGW_OBJ_NS_MULTIPART`（值为 `"multipart"`）专门用于**多段上传**：`.meta` 元数据对象、每个分段对象都在此命名空间下；使用 upload id 参与索引哈希，避免与正式对象冲突；便于 GC/生命周期管理未完成的多段上传。  
- `RGW_OBJ_NS_SHADOW`（值为 `"shadow"`）用于**影子对象/条带对象**：大文件拆条存储、Copy 操作的临时对象、manifest 中的 tail 对象等，同样对用户列表不可见。  
两者都是 RGW 内部使用的桶索引空间，通过命名空间区分不同类型的条目，对用户透明。

# q
RGW 多段上传的相关对象使用哪个命名空间？为什么？
# a
使用 `RGW_OBJ_NS_MULTIPART`（值为 `"multipart"`）。这样做的目的包括：  
1. 这些条目不会被普通的 `list_objects` 列举出来；  
2. 索引哈希使用 upload id 而非原对象名，防止与正式对象冲突；  
3. 方便生命周期、GC 等模块识别并清理未完成的 multipart 元数据或分段。

# q
RGW 中的影子对象或临时拷贝对象存放在哪个命名空间？请举例说明。
# a
存放在 `RGW_OBJ_NS_SHADOW`（值为 `"shadow"`）命名空间。  
典型用途：  
- 大对象拆条存储时的条带对象；  
- Copy 操作中写到的临时对象（如 `rgw_rados.cc::copy_obj()` 会生成 `shadow_obj` 临时存放拷贝数据）；  
- manifest 中引用的 tail 对象。  
这些中间结果不出现在用户可见的对象列表里，并与正式对象分开，便于 GC 和版本管理。

