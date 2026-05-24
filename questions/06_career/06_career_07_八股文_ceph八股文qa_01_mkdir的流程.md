# q
介绍下 Lustre mkdir 的流程
# a
```
VFS层: sys_mkdir
  ↓
ll_mkdir (lustre/llite/namei.c:1891)
  ↓
ll_new_node (lustre/llite/namei.c:1537-1686)
  ↓
md_create (lustre/lmv/lmv_obd.c:1908) 
  ↓
lmv_create → 选择目标MDT
  ↓
mdt_create (lustre/mdt/mdt_reint.c:498) - 服务端处理
  ↓
mdd_create (lustre/mdd/mdd_dir.c:2615) - MDD层
  ↓
OSD层 (ldiskfs/zfs)
```

