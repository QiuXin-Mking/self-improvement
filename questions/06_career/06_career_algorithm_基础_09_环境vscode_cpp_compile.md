# q
推荐的 VS Code 实用插件有哪些？各有什么作用？
# a
- **Chinese**：中文界面插件，安装后即显示中文。
- **Partial Diff**：比较代码、文件差异的插件。
- **Rainbow Brackets**：多重括号高亮，不同颜色区分括号，便于阅读。
- **One Dark Pro**：皮肤插件，美化编辑器界面。
- **Window Colors**：多开窗口颜色区分，每个 VS Code 窗口显示不同颜色，方便辨识。
- **Tabnine AI**：AI 代码补全插件，通过 Tab 键补全预测的代码。

# q
如何在 Windows 上安装 CMake 并验证是否安装成功？
# a
1. 从 https://cmake.org/download/ 下载 `.msi` 64 位安装包并安装。
2. 系统变量 `PATH` 中添加 `C:\Program Files\CMake\bin`。
3. 打开命令行输入 `cmake --version` 查看版本信息，能正常输出即安装完成。

# q
如何安装 g++、gcc 和 make，并检查它们是否就绪？
# a
通过安装 Mingw-w64 来获得这些工具：
- 下载并安装 Mingw-w64 到 `C:\Program Files\mingw64\bin`。
- 该目录下包含 `g++.exe`、`gcc.exe`、`make.exe` 等。
- 验证安装：在命令行中分别执行 `g++ --version`、`gcc --version`、`make --version`，有版本输出即说明已就绪。

# q
VS Code 中 CMake Tool 插件的作用是什么？
# a
用于生成 CMake 配置文件 `CMakeLists.txt`，帮助自动化构建配置，方便在 VS Code 中进行 C++ 项目的编译和调试。

