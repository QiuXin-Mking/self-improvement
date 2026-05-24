# q
Lustre中`lsattr`输出中的`Project_Hierarchy`属性表示什么？
# a
`Project_Hierarchy`表示该文件或目录已被分配了一个Project ID，并被纳入项目的层级管理。它通常通过`lfs project -s -p`命令递归设置，使得目录及其下所有内容都携带该属性，用于配额、数据放置等管理。

