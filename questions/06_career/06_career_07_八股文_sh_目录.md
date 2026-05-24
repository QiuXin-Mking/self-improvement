# q
如何获取当前路径并赋值给变量？
# a
```bash
current_path=$(pwd)
```

# q
如何计算Shell脚本的执行时间？
# a
```bash
start_time=$(date +%s)  # 记录开始时间
end_time=$(date +%s)    # 记录结束时间
elapsed_time=$((end_time - start_time))
echo "脚本运行时间为: $elapsed_time 秒"
```

# q
如何使用cp命令强制覆盖单个文件和递归覆盖目录？
# a
```bash
# 强制覆盖单个文件
cp -f source.txt target.txt

# 强制覆盖目录中的文件（需递归）
cp -rf source_dir/ target_dir/
```

