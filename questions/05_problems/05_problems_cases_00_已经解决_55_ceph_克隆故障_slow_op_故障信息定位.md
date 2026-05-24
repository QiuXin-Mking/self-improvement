# q
如何定位与特定 rbd_data 关联的 RBD 镜像？
# a
遍历指定池中的所有 RBD 镜像，使用 `rbd info` 检查是否包含给定的 rbd_data。以下脚本以 vms 池为例：
```sh
arr=($(rbd ls -p vms))
rbd_data=706798b867499d
for img in "${arr[@]}"; do
   echo "$img"
   if rbd info vms/"$img" | grep -q "$rbd_data"; then
       echo "this is not good"
       echo "$img"
   fi
done
```
如果镜像的详细信息中匹配到该 rbd_data，则输出警告和镜像名称。

