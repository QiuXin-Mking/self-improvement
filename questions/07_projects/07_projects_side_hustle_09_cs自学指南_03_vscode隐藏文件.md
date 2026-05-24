# q
如何在 VS Code 的用户设置中隐藏特定文件或文件夹？
# a
通过快捷键 `Ctrl/Command+Shift+P` 打开命令面板，输入 `setting`，选择 `Preferences: Open User Settings (JSON)`，然后在 `files.exclude` 字段中添加 glob 匹配模式，值为 `true`。例如：
```json
"files.exclude": {
    "**/*.pyc": true,
    "**/Logs": true,
    "**/temp": true
}
```

# q
在 `files.exclude` 配置中，如何用 glob 模式匹配不区分大小写的文件夹名？
# a
使用中括号包裹首字母的大小写形式，例如 `"**/[Pp]lugins": true` 将同时匹配 `Plugins` 和 `plugins` 文件夹。这是一种简单的字符类写法。

