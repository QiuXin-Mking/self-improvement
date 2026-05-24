# q
使用 `iscsiadm` 发现 iSCSI 目标时应指定什么发现模式（`-m`）和发现类型（`-t`）？
# a
应使用 `-m discovery` 指定发现模式，并使用 `-t st`（SendTargets）作为发现类型。
完整示例：
```sh
iscsiadm -m discovery -t st -p 192.168.0.1
```

# q
在 `iscsiadm` 发现命令中如何开启详细的调试输出？
# a
通过添加 `-d8` 选项，将调试级别设为最高（8），以输出最详细的交互和错误信息。
示例：
```sh
iscsiadm -m discovery -t st -d8 -p 192.168.0.1
```

# q
在 `iscsiadm` 发现目标时，如何指定 iSCSI 目标服务器的 IP 地址？
# a
使用 `-p` 参数后跟目标 IP 地址，例如 `-p 192.168.0.1`。该参数指定要连接的 iSCSI 门户（portal）地址。

