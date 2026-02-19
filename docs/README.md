# 基于艾宾浩斯遗忘曲线的知识库学习系统

一个Go命令行应用，帮助你根据艾宾浩斯遗忘曲线科学地复习知识。

## 功能特点

- 📚 **自动解析Markdown文件**：从配置的多个目录下的`.md`文件中提取问题和答案
- 🔧 **灵活的配置**：支持配置多个问题目录，方便组织不同主题的知识
- 🧠 **智能复习调度**：基于艾宾浩斯遗忘曲线，自动安排复习时间
- 📊 **学习统计**：跟踪你的学习进度和正确率
- 🎯 **四种记忆程度**：熟练、一般、忘记、完全忘记，精确反馈记忆状态

## 安装

1. 克隆或下载项目到本地

2. 安装Go (>= 1.21)

3. 安装依赖：
```bash
make deps
```

或者直接运行:
```bash
go mod tidy
```

## 使用说明

### 1. 准备问题文件

默认在`questions/`目录下创建`.md`文件，格式如下：

```markdown
# q
你的问题

# a
你的答案

# q
另一个问题？

# a
另一个答案。
```

每个问题格式：
- `# q` 开头，下面是问题内容
- `# a` 开头，下面是答案内容

#### 配置多个目录（可选）

编辑`question_input`文件可以配置多个目录：

```
questions
my_questions
work_knowledge
path/to/other/questions
```

这样可以按主题组织不同的问题，例如：
- `questions/` - 基础问题
- `work_knowledge/` - 工作相关知识
- `interviews/` - 面试题

### 2. 初始化知识库

首次使用需要初始化知识库：

```bash
make init
# 或者
go run main.go --init
```

这将扫描配置的所有目录下的`.md`文件，提取问题和答案到知识库。

### 3. 开始训练

每天上班运行以下命令开始训练：

```bash
make run-cli
# 或者
go run main.go
```

程序会显示今天需要复习的问题。

### 4. 使用流程

1. **查看问题**：程序显示问题，你在纸上默写答案
2. **查看答案**：输入 `a` 或 `answer` 查看标准答案
3. **反馈记忆程度**：
   - 输入 `1` 或 `熟练` - 记得很清楚
   - 输入 `2` 或 `一般` - 记得但不熟练
   - 输入 `3` 或 `忘记` - 忘记了部分内容
   - 输入 `4` 或 `完全忘记` - 完全不记得
4. **继续下一个问题**：重复上述流程

### 5. 查看统计

查看学习统计信息：

```bash
make stats
# 或者
go run main.go --stats
```

## 艾宾浩斯遗忘曲线算法

系统根据你的反馈自动调整复习间隔：

- **熟练**：7天后复习，正确后间隔递增2.5倍
- **一般**：3天后复习，正确后间隔递增1.8倍
- **忘记**：1天后复习，正确后间隔递增1.3倍
- **完全忘记**：2小时后复习，保持原间隔

连续正确回答后，系统会自动提高你的记忆等级，延长复习间隔。

## 数据存储

学习数据保存在 `data/learning_data.json` 文件中，包含：
- 所有问题和答案
- 复习历史
- 记忆等级
- 下次复习时间
- 学习统计

## 命令参考

```bash
# 初始化知识库
go run main.go --init

# 开始训练（默认）
go run main.go

# 查看统计
go run main.go --stats

# 使用Makefile命令
make init     # 初始化
make run-cli  # 运行CLI
make run-web  # 运行Web服务
make build    # 构建二进制文件
make stats    # 查看统计
```

## 提示

- 建议每天固定时间运行程序，养成习惯
- 诚实反馈记忆程度，系统才能准确安排复习
- 数据文件会自动保存，无需手动操作
- 可以在任何时候输入 `q` 退出训练

## 项目结构

```
self-improvement/
├── main.go                     # 主CLI程序
├── web_server.go              # Web服务器
├── internal/
│   ├── spacedrepetition/      # 艾宾浩斯算法
│   │   └── spaced_repetition.go
│   └── parser/                # Markdown解析器
│       └── parser.go
├── go.mod                    # Go模块定义
├── go.sum                    # Go依赖校验
├── Makefile                  # 构建脚本
├── README.md                 # 说明文档
├── 使用指南.md               # 中文快速指南
├── 部署指南.md               # 服务器部署指南
├── QUICK_START_DEPLOY.md     # 快速部署参考
├── questions/                 # 存放问题文件（.md格式）
└── data/                      # 学习数据（自动生成）
    └── learning_data.json
```

## Web界面

启动Web服务：

```bash
make run-web
# 或者
go run web_server.go
```

访问 http://localhost:5000 来使用Web界面进行学习。

## 编译二进制文件

编译CLI和Web应用程序：

```bash
make build
# 或者分别构建
make build-cli  # 构建CLI应用
make build-web  # 构建Web应用
```

编译后的文件将位于 `bin/` 目录。

## 示例问题文件

创建 `questions/python_basics.md`:

```markdown
# q
什么是Python中的装饰器？

# a
装饰器是一种设计模式，在Python中允许用户在不修改原函数代码的情况下，动态地给函数添加功能。装饰器本质上是一个接受函数作为参数并返回函数的高阶函数。

# q
Python中的GIL是什么？

# a
GIL (Global Interpreter Lock) 是Python解释器中的一个互斥锁，它确保同一时刻只有一个线程执行Python字节码。这意味着即使在多核CPU上，Python的多线程也无法真正实现并行执行CPU密集型任务。
```

## 注意事项

- 确保`.md`文件使用UTF-8编码
- 问题和答案之间用`# q`和`# a`分隔
- 支持多个问题在一个文件中
- 程序会自动跳过已导入的问题（基于问题内容的hash）

祝你学习愉快！🎓