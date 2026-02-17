# 基于艾宾浩斯遗忘曲线的知识库学习系统（多租户版）

一个支持多用户的Go Web应用，帮助你根据艾宾浩斯遗忘曲线科学地复习知识。现在支持多用户隔离、身份验证和现代化前端界面。

## ✨ 新增功能

- 👥 **多租户支持**：每个用户拥有独立的学习数据，数据完全隔离
- 🔐 **身份验证**：JWT-based认证，支持注册和登录
- 🌐 **现代化前端**：Vue.js 3 + TypeScript + Vant UI组件库
- 💾 **数据库支持**：使用SQLite存储，支持更复杂的数据结构
- 📱 **响应式设计**：适配移动端和桌面端

## 功能特点

- 📚 **自动解析Markdown文件**：从配置的多个目录下的`.md`文件中提取问题和答案
- 🔧 **灵活的配置**：支持配置多个问题目录，方便组织不同主题的知识
- 🧠 **智能复习调度**：基于艾宾浩斯遗忘曲线，自动安排复习时间
- 📊 **学习统计**：跟踪你的学习进度和正确率
- 🎯 **四种记忆程度**：熟练、一般、忘记、完全忘记，精确反馈记忆状态
- 👥 **多用户支持**：每个用户独立的学习空间，数据完全隔离
- 🔐 **安全认证**：JWT令牌验证，确保数据安全

## 安装

1. 克隆或下载项目到本地

2. 安装Go (>= 1.21) 和 Node.js

3. 安装依赖：
```bash
make deps
```

或者分别安装：

```bash
# Go dependencies
go mod tidy

# Frontend dependencies
cd frontend && npm install
```

## 配置环境变量

复制示例环境文件并配置：

```bash
cp .env.example .env
```

设置JWT密钥：
```
JWT_SECRET=your-very-secure-random-string-for-production
```

## 前端构建

```bash
# 构建前端资源
make frontend-build

# 开发模式运行前端（单独）
cd frontend && npm run dev
```

## 后端运行

### 1. 准备问题文件

在`questions/`目录下创建`.md`文件，格式如下：

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

### 2. 数据库初始化

首次使用需要初始化数据库：

```bash
# 初始化数据库表结构
make db-init

# 运行Web服务（会自动初始化数据库）
make run-web-with-db
```

### 3. 运行Web服务

```bash
# 运行Web服务器
make run-web

# 或者直接运行（需要设置JWT_SECRET）
JWT_SECRET=your-secret-key go run web_server.go
```

访问 http://localhost:5000 开始使用。

### 4. 用户使用流程

1. **注册账户**：访问 `/register` 页面注册新账户
2. **登录系统**：访问 `/login` 页面使用注册的账户登录
3. **初始化知识库**：登录后可在 `/dashboard` 页面初始化知识库
4. **开始学习**：点击开始学习按钮开始复习
5. **查看统计**：在仪表盘查看个人学习统计数据

## Docker 部署

构建并运行Docker容器：

```bash
# 构建并运行
docker-compose up -d

# 构建镜像
docker build -t spaced-repetition-multi-tenant .
```

## API 接口

- `POST /api/register` - 用户注册
- `POST /api/login` - 用户登录
- `GET /api/profile` - 获取用户信息
- `GET /api/stats` - 获取学习统计
- `GET /api/due-questions` - 获取待复习问题
- `POST /api/update-review` - 提交复习反馈
- `POST /api/delete-question` - 删除问题
- `POST /api/init` - 初始化知识库

## 项目结构

```
self-improvement/
├── main.go                     # 主CLI程序
├── web_server.go              # Web服务器（多租户版）
├── internal/
│   ├── models/                # 数据模型（用户、问题）
│   │   ├── user.go
│   │   └── question.go
│   ├── middleware/            # 中间件（认证）
│   │   └── auth.go
│   ├── spacedrepetition/      # 艾宾浩斯算法（多租户版）
│   │   └── spaced_repetition.go
│   └── parser/                # Markdown解析器
│       └── parser.go
├── handlers/                  # API处理器
│   └── auth.go
├── migrations/                # 数据库迁移
│   └── 001_initial_schema.sql
├── frontend/                  # Vue.js前端应用
│   ├── src/
│   │   ├── api/              # API接口定义
│   │   ├── stores/           # Pinia状态管理
│   │   ├── views/            # 页面组件
│   │   ├── components/       # 通用组件
│   │   ├── router/           # 路由配置
│   │   └── assets/           # 静态资源
├── go.mod                    # Go模块定义
├── go.sum                    # Go依赖校验
├── package*.json             # 前端依赖
├── Makefile                  # 构建脚本
├── Dockerfile                # Docker构建文件
├── docker-compose.yml        # Docker编排配置
├── .env.example              # 环境变量示例
├── README.md                 # 说明文档
├── questions/                 # 存放问题文件（.md格式）
└── data/                      # 学习数据（自动生成）
    └── app.db
```

## 编译二进制文件

编译应用程序：

```bash
make build
```

编译后的文件将位于 `bin/` 目录。

## 部署到生产环境

1. 配置环境变量（特别是JWT_SECRET）
2. 设置正确的数据库路径
3. 构建前端资源：`make frontend-build`
4. 使用Docker或直接运行二进制文件

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

- 生产环境中务必设置强密钥的JWT_SECRET
- 确保`.md`文件使用UTF-8编码
- 问题和答案之间用`# q`和`# a`分隔
- 支持多个问题在一个文件中
- 不同用户的数据完全隔离

祝你学习愉快！🎓