# q
如何查看 RGW bucket 的内部 ID 及统计信息？
# a
使用命令：
```
radosgw-admin bucket stats --bucket-id=23d62b95-a477-443e-859b-2b3aa2c50b00.1877119039.1
```
可以获取指定 bucket 的统计信息和内部标识符。

# q
如何取消一个未完成的 S3 分段上传（multipart upload）？
# a
使用 AWS CLI 的 `abort-multipart-upload` 命令，需指定 bucket、key 和 upload-id：
```
aws s3api abort-multipart-upload --bucket upload --key 20250824153406705.7z --upload-id 18938166915
```

# q
如何从 RADOS 底层直接导出或删除残留的多分段上传对象？
# a
- **导出对象**：
  ```
  rados -p obj_data_3fb_045b35ee get '23d62b95-a477-443e-859b-2b3aa2c50b00.1877119039.1__multipart_18938166915/20250824153406705.7z.2~RavgtcGyM13EjGi3c2QlQ5xk2wvEnQ8.77' backup_20250824153406705.7z.part2
  ```
- **删除对象**：
  ```
  rados -p obj_data_3fb_045b35ee rm '23d62b95-a477-443e-859b-2b3aa2c50b00.1877119039.1__multipart_18938166915/20250824153406705.7z.2~RavgtcGyM13EjGi3c2QlQ5xk2wvEnQ8.77'
  ```
对象名称格式为 `<bucket-id>__multipart_<upload-id>/<key>.<part>~<随机后缀>`。

