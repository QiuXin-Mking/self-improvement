# q
Git 中 `core.autocrlf` 配置的作用是什么？它的 `input`、`true`、`false` 三种取值分别有什么效果？
# a
`core.autocrlf` 用于控制 Git 在提交和检出时是否自动转换行尾（LF 与 CRLF）。三种取值的行为：
- **input**：提交时将 CRLF 转为 LF，检出时不转换，文件始终保持 LF。
- **true**：检出时将 LF 转为 CRLF，提交时将 CRLF 转回 LF，确保仓库内始终为 LF，适合 Windows 工作区。
- **false**：不进行任何自动转换，保存文件原本的行尾格式。

# q
`core.autocrlf` 的 `input` 和 `true` 在检出与提交时的行为有何核心区别？
# a
- **input** 只在提交时将 CRLF 转为 LF，检出时不改变行尾，因此工作区文件始终为 LF。
- **true** 在检出时将仓库中的 LF 转为 CRLF（适配 Windows），提交时再将 CRLF 转回 LF，工作区文件通常是 CRLF，仓库内为 LF。

# q
如何用 `sed` 命令递归删除一个文件中所有的 `\r` 字符？
# a
可以使用以下命令直接删除文件中所有回车符（`\r`）：
```bash
sed -i 's/\r$//' make.sh
```
该命令会原地修改文件，将每行末尾的 `\r`（即 CRLF 中的 CR）删除，使行尾变为 LF。

