# q
获取当前路径并复制到一个变量
# a
current_path=$(pwd)

# q
写个计算脚本执行时间的sh代码？包括开始时间、结束时间、时间差和导出
# a
```bash
start_time=$(date +%s)  # 记录开始时间
end_time=$(date +%s)    # 记录结束时间
elapsed_time=$((end_time - start_time))
echo "脚本运行时间为: $elapsed_time 秒"
```

# q
cp命令如何强制覆盖单个文件，如何强制覆盖目录中的文件（需递归）？
# a
- 强制覆盖单个文件：`cp -f source.txt target.txt`
- 强制覆盖目录中的文件（需递归）：`cp -rf source_dir/ target_dir/`

