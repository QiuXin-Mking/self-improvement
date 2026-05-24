# q
RGW_OBJ_NS_MULTIPART 命名空间的用途是什么？
# a
用于多段上传（Multipart Upload）对象。元数据对象和每个分段（part）都挂在这个命名空间下，以避免被普通对象列举扫到，并使用 upload id 作为索引哈希，避免与正式对象冲突。例如在 `rgw_multi.cc::abort_multipart_upload()` 中通过 `obj.init_ns(..., RGW_OBJ_NS_MULTIPART)` 删除相关条目。

# q
RGW_OBJ_NS_SHADOW 命名空间的用途是什么？
# a
用于 RGW 内部的“影子对象”或“条带对象”。例如大对象拆条存储、Copy 操作中的临时对象、manifest 引用的 tail 对象等。这些对象对用户透明，与正式对象隔离，便于 GC 和版本管理。例如 `rgw_rados.cc::copy_obj()` 会生成 `shadow_obj` 存放临时拷贝数据。

# q
RGW_OBJ_NS_MULTIPART 和 RGW_OBJ_NS_SHADOW 共同的设计目标是什么？
# a
两者都是 Ceph RGW 在桶索引中使用的内部命名空间，通过将特殊用途对象（多段上传部件、影子/条带对象）与用户正式上传的对象隔离开，使得这些后台逻辑产生的条目不会出现在用户可见的对象列表中，同时便于生命周期管理、垃圾回收等模块按类型区分处理。

