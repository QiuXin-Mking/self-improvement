# q
在Shell中如何用date命令获取上一个完整月的第一天（格式YYYYMMDD）？
# a
```sh
date -d '1 month ago' +%Y%m01
```

# q
如何用date命令获取当前时间的纳秒级时间戳，格式为年月日时分秒纳秒？
# a
```sh
date "+%Y%m%d%H%M%S%N"
```

# q
在Shell脚本中，如何获取当前小时数，并判断是否等于13？
# a
```sh
date_time_h=`date "+%H"`
if (( $date_time_h == 13 ))
then
  echo "现在是13点"
fi
```

# q
在Shell中如何根据`uname -s`判断操作系统类型？
# a
```sh
SYSTEM=`uname -s`
if [ $SYSTEM = "Linux" ] ; then
   echo "Linux"
elif [ $SYSTEM = "FreeBSD" ] ; then
   echo "FreeBSD"
else
   echo "其他系统"
fi
```

