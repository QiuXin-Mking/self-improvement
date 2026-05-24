# q
查询lustre 所有mdt线程中，使用率总和是多少？并解释意义。提示 top ,grep, awk，
# a
top -bn1 | grep mdt | awk 'BEGIN{s=0}{s+=$9}END{print s}'

