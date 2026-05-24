# q
clang-format 是什么，它集成于哪个项目？
# a
clang-format 是一个代码格式化工具，集成于 LLVM 项目，可在 GitHub 上通过 `https://github.com/llvm/llvm-project/releases` 下载。

# q
在 VS Code 中配置 clang-format 时，`settings.json` 里需要设置哪些关键选项？
# a
```json
{
    "clang-format.style" : "file",
    "clang-format.assumeFilename": "./.clang-format",
    "clang-format.executable": "/usr/bin/clang-format",
    "clang-format.fallbackStyle": "Google",
    "editor.defaultFormatter": "xaver.clang-format"
}
```
其中 `assumeFilename` 指定风格文件路径（如 `./.clang-format` 或绝对路径），`executable` 指定 clang-format 可执行文件位置，`fallbackStyle` 为回退风格（如 `Google`）。

# q
如何生成一个 `.clang-format` 配置文件并基于 LLVM 风格进行自定义？
# a
使用 clang-format 的命令行选项：
```
clang-format.exe --style=LLVM --dump-config > ./.clang-format
```
该命令将 LLVM 风格的配置内容导出到当前目录的 `.clang-format` 文件中，之后可手动修改。

