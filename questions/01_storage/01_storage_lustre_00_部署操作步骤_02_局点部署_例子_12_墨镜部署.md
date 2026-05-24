# q
如何在GRUB中永久添加内核启动参数tsc=reliable和tsc=nowatchdog？
# a
编辑 `/etc/default/grub`，在 `GRUB_CMDLINE_LINUX` 中添加 `tsc=reliable tsc=nowatchdog`，然后执行 `grub2-mkconfig -o /boot/grub2/grub.cfg`（若为UEFI引导还需执行对应efi路径的grub.cfg生成命令），最后重启系统使参数生效。

# q
如何使用rbd命令为Lustre创建并映射块设备？
# a
创建：`rbd create --size <大小> rbd/<镜像名> --image-feature layering`（如 `rbd create --size 10G rbd/mj_mgt00 --image-feature layering`）。
映射：`sudo rbd map rbd/<镜像名> --id admin`。

# q
Lustre文件系统格式化时`--mgs`、`--mdt`、`--ost`参数各自的作用是什么？
# a
- `--mgs`：将设备格式化为管理服务（MGS）节点。
- `--mdt`：将设备格式化为元数据目标（MDT），需指定 `--index`。
- `--ost`：将设备格式化为对象存储目标（OST），需指定 `--index`。
三者均需通过 `--mgsnode` 指定MGS节点的NID列表（如 `192.168.6.188@tcp`）。

# q
在客户端如何挂载Lustre文件系统？
# a
```bash
mount -t lustre -o noatime <MGS_NID>@tcp:/<文件系统名> <挂载点>
```
例如：`mount -t lustre -o noatime 192.168.6.188@tcp:192.168.6.189@tcp:/mj_nas /mnt/lustre`。

