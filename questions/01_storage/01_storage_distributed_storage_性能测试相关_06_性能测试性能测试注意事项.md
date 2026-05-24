# q
如何清理iSCSI缓存并重置initiator配置？
# a
依次执行以下命令：先登出所有节点，删除节点记录，清除send_targets和nodes目录下的缓存文件，最后确认initiator名称。
```sh
iscsiadm -m node -u
sleep 2
iscsiadm -m node -o delete

rm -rf /var/lib/iscsi/send_targets/*
rm -rf /var/lib/iscsi/nodes/*

cat /etc/iscsi/initiatorname.iscsi
```

# q
前端如何发现并登录后端iSCSI target？
# a
首先使用discovery命令扫描可用的target：
```sh
iscsiadm -m discovery -t st -p 172.250.80.80:3260
iscsiadm -m discovery -t st -p 172.250.80.81:3260
iscsiadm -m discovery -t st -p 172.250.80.82:3260
```
然后使用发现到的IQN和IP登录：
```sh
iscsiadm -m node -T iqn.2010-05.com.macrosan.target:rout2:692bb5d528284af982f6e72c33af5369:1478 -p 172.250.80.81:3260 -l
```
通用格式为：
```sh
iscsiadm -m node -T <iqn> -p <ip>:3260 -l
```

# q
性能测试中如何修改刷盘水位线，其配置路径是什么？
# a
刷盘水位线配置文件位于 `/engine-fs/ocache/<卷编号>/wf_low`，内容示例为 `flush_low_level: [5]`。修改方式为：
```sh
echo 69 > /engine-fs/ocache/25/wf_low
echo 69 > /engine-fs/ocache/26/wf_low
# ... 对需要调整的卷依次执行
```
批量查看：
```sh
cat /engine-fs/ocache/*/wf_low
```

# q
如何使用nmon进行性能监控？命令中参数含义是什么？
# a
启动nmon监控的命令：
```sh
./nmon -s10 -c370 -f -m ~/nmon
```
参数含义：
- `-s10`：每10秒采集一次数据
- `-c370`：总共采集370次（总时长约3700秒）
- `-f`：输出为文件格式
- `-m ~/nmon`：结果文件保存到~/nmon目录

