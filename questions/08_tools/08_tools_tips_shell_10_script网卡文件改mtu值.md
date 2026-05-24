# q
在 Bash 脚本中，如何检查当前是否以 root 用户身份运行？
# a
使用环境变量 `$EUID`，当值为 0 时表示 root 身份。示例：
```bash
if [[ $EUID -ne 0 ]]; then
    echo "请使用root权限运行此脚本。"
    exit 1
fi
```

# q
如何使用 `getopts` 解析 `-i` 和 `-m` 命令行选项？
# a
```bash
while getopts "i:m:" opt; do
    case $opt in
        i) interface=$OPTARG ;;
        m) mtu_value=$OPTARG ;;
        \?) echo "无效的选项： -$OPTARG" ;;
    esac
done
```
冒号表示该选项需要参数，`$OPTARG` 存储参数值。

# q
该脚本如何在配置文件中修改或新增 MTU 参数？
# a
先用 `grep -q "^MTU=" "$config_file"` 检测是否存在；若存在则用 `sed -i` 替换：
```bash
sed -i "s/^MTU=.*/MTU=$mtu_value/" "$config_file"
```
若不存在则追加：
```bash
echo "MTU=$mtu_value" >> "$config_file"
```

# q
修改完 MTU 配置后，如何让设置立即生效？
# a
执行 `systemctl restart network` 重启网络服务。

# q
脚本如何检查必要参数是否已提供？
# a
通过 `[[ -z $interface || -z $mtu_value ]]` 判断变量是否为空，若为空则打印用法并退出：
```bash
if [[ -z $interface || -z $mtu_value ]]; then
    echo "缺少必要的参数。"
    echo "用法: $0 -i <interface> -m <mtu_value>"
    exit 1
fi
```

