# 🎓 基于艾宾浩斯遗忘曲线的知识库学习系统

一个基于艾宾浩斯遗忘曲线的科学学习系统，支持命令行（CLI）和 Web 界面两种使用方式。系统通过智能算法自动安排复习时间，帮助你高效记忆和掌握知识。

## ✨ 核心功能

### 📚 知识管理
- **Markdown 格式**：使用简单的 Markdown 格式管理问题和答案
- **多目录支持**：可配置多个问题目录，按主题组织知识
- **自动解析**：自动扫描并解析 `.md` 文件中的问答对
- **去重机制**：基于问题内容哈希自动去重，避免重复

### 🧠 智能复习算法
基于 **艾宾浩斯遗忘曲线**，根据你的反馈自动调整复习间隔：
- **熟练**（Level 1）：7 天后复习，间隔递增 2.5 倍
- **一般**（Level 2）：3 天后复习，间隔递增 1.8 倍
- **忘记**（Level 3）：1 天后复习，间隔递增 1.3 倍
- **完全忘记**（Level 4）：2 小时后复习，保持原间隔
- **准确率奖励**：整体准确率 > 80% 时，所有间隔额外乘以 1.2 倍

### 📊 学习统计
- 总问题数
- 今日待复习数
- 总复习次数
- 正确次数
- 正确率统计

### 🌐 双界面支持

#### CLI 命令行界面
- 适合本地开发环境
- 支持键盘快捷操作
- 支持删除低质量问题

#### Web 界面
- **Vue 3 + Vite** 构建的现代化前端
- **响应式设计**，支持移动端访问
- **用户认证**：注册、登录、JWT 令牌验证
- **实时统计**：查看学习进度和统计数据
- **直观交互**：卡片式问答界面，反馈按钮

## 🛠️ 技术栈

### 后端
- **Go** - 高性能后端语言
- **Gin** - Web 框架
- **GORM** - ORM 框架
- **SQLite** - 轻量级数据库
- **JWT** - 用户认证

### 前端
- **Vue 3** - 渐进式 JavaScript 框架
- **Vite** - 快速构建工具
- **Pinia** - 状态管理
- **Vue Router** - 路由管理
- **Vant** - 移动端 UI 组件库
- **SCSS** - 样式预处理器

## 📦 安装部署

### 本地开发

```bash
# 克隆项目
git clone https://github.com/your-username/self-improvement.git
cd self-improvement

# 安装 Go 依赖
go mod tidy

# 初始化知识库
go run main.go --init

# 运行 CLI 应用
go run main.go

# 运行 Web 服务（默认端口 5000）
go run web_server.go
```

### 使用 Makefile

```bash
make deps      # 安装依赖
make init      # 初始化知识库
make run-cli   # 运行 CLI 应用
make run-web   # 运行 Web 服务
make build     # 构建二进制文件
```

### Docker 部署

```bash
# 使用 Docker Compose
docker-compose up -d

# 或者单独构建运行
docker build -t spaced-repetition .
docker run -p 5000:5000 spaced-repetition
```

## 📝 使用方法

### 准备问题文件

创建 `.md` 文件，使用以下格式：

```markdown
# q
什么是闭包？

# a
闭包是指有权访问另一个函数作用域中变量的函数。在 JavaScript 中，闭包是在函数内部创建的函数，可以访问外部函数的变量。

# q
什么是原型链？

# a
原型链是 JavaScript 中实现继承的机制。每个对象都有一个内部链接到另一个对象的原型，当访问属性时，会沿着原型链向上查找。
```

### 配置问题目录

创建 `question_input` 文件（Linux）或 `question_input_windows` 文件（Windows），每行一个目录路径：

```
questions/javascript
questions/python
work/interviews
```

### CLI 使用流程

1. 查看问题 → 在纸上默写答案
2. 输入 `a` 查看标准答案
3. 输入反馈（1-4）：
   - `1` 熟练
   - `2` 一般
   - `3` 忘记
   - `4` 完全忘记
4. 继续下一个问题

### Web 界面使用

1. 访问 http://localhost:5000
2. 注册/登录账号
3. 点击"开始复习"
4. 查看问题并回忆答案
5. 点击"显示答案"
6. 选择反馈按钮
7. 自动跳转到下一个问题

## 📂 项目结构

```
self-improvement/
├── main.go                     # CLI 主程序
├── web_server.go              # Web 服务器
├── internal/
│   ├── spacedrepetition/      # 艾宾浩斯算法
│   ├── parser/               # Markdown 解析器
│   ├── models/               # 数据模型
│   └── middleware/           # JWT 认证中间件
├── frontend/                 # Vue 3 前端项目
│   ├── src/
│   │   ├── api/            # API 调用
│   │   ├── stores/         # Pinia 状态管理
│   │   ├── views/          # 页面组件
│   │   └── components/     # 可复用组件
│   └── public/            # 静态资源
├── scripts/
│   ├── clear_learning_data.go   # 清除复习数据脚本
│   └── clear_learning_data.sh  # Shell 版本清理脚本
├── docs/                    # 文档目录
├── migrations/               # 数据库迁移文件
├── questions/                # 问题文件存放目录
└── data/                    # 运行时数据（Git 忽略）
    ├── learning_data.json
    └── app.db
```

## 🔒 数据安全

- **用户隔离**：SQLite 数据库支持多用户，每个用户的数据相互隔离
- **密码加密**：使用 bcrypt 对用户密码进行哈希存储
- **JWT 认证**：Web 界面使用 JWT 令牌进行身份验证
- **数据备份**：提供数据清除脚本时自动备份原文件

## 📈 算法原理

### 复习间隔计算

```go
baseInterval = [168, 72, 24, 2][feedback - 1]  // 小时
multiplier = [2.5, 1.8, 1.3, 1.0][feedback - 1]
accuracyBonus = accuracy > 0.8 ? 1.2 : 1.0
nextInterval = baseInterval * multiplier * accuracyBonus
```

### 记忆等级调整

- 连续 3+ 次正确回答：等级提升（间隔延长）
- 回答错误：等级降低（间隔缩短）

## 🧪 测试数据清理

开发测试时可以清除复习进度数据：

```bash
# 使用 Go 脚本（推荐）
go run scripts/clear_learning_data.go

# 使用 Shell 脚本
./scripts/clear_learning_data.sh
```

详细说明请参考 [docs/clear_learning_data.md](docs/clear_learning_data.md)

## 📄 文档

- [使用指南](使用指南.md) - 中文快速上手指南
- [部署指南](部署指南.md) - 服务器部署详细说明
- [API 文档](API_DOCUMENTATION.md) - REST API 接口说明
- [部署测试计划](DEPLOYMENT_TEST_PLAN.md) - 部署验证清单
- [清除数据说明](docs/clear_learning_data.md) - 清理复习数据指南

## 🚀 路线图

- [ ] 支持图片题目
- [ ] 导入 Anki 卡片
- [ ] 多语言支持
- [ ] 移动端原生应用
- [ ] 学习数据可视化图表
- [ ] 学习提醒功能

## 📝 许可证

本项目采用 MIT 许可证开源。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## ⚠️ 注意事项

- 确保 `.md` 文件使用 UTF-8 编码
- 诚实反馈记忆程度，系统才能准确安排复习
- 定期备份数据文件
- 数据文件已配置 Git 忽略，不会提交到仓库

---

**祝学习愉快！🎓**
