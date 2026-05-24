# q
如何通过命令行修改RGW bucket的分片（shard）数量？
# a
使用 `radosgw-admin reshard bucket` 命令，指定目标 bucket 名称和新的分片数量：
```bash
radosgw-admin reshard bucket --bucket <bucket_name> --num-shards <shard_num>
```

