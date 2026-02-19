# 艾宾浩斯遗忘曲线学习系统 - 部署测试计划

## 项目概述

这是一个基于艾宾浩斯遗忘曲线的间隔重复学习系统，包含：
- **后端**: Go + Gin + GORM + SQLite
- **前端**: Vue 3 + Vite + Vant UI
- **数据存储**: SQLite 数据库

## 环境检查

### 1. 依赖检查
```bash
# 检查 Go 版本 (需要 1.19+)
go version

# 检查 Node.js 版本 (需要 16+)
node --version
npm --version
```

### 2. 现有文件结构
```
Self-improvement/
├── go.mod                 # Go 依赖配置
├── web_server.go          # Web 服务器主文件
├── main.go                # CLI 应用主文件
├── internal/              # Go 内部包
│   ├── models/           # 数据模型 (User, Question)
│   ├── middleware/       # JWT 认证中间件
│   ├── parser/          # Markdown 问题解析器
│   └── spacedrepetition/ # 间隔重复算法
├── templates/            # HTML 模板
├── static/              # 静态资源
├── frontend/            # Vue 前端
│   ├── src/
│   ├── package.json
│   └── vite.config.ts
├── questions/           # 问题 Markdown 文件
├── question_input_linux # 问题目录配置
└── data/                # 数据库目录 (自动创建)
```

## 部署步骤

### 步骤 1: 安装后端依赖
```bash
cd /Users/qiuxin/code/Self-improvement
go mod tidy
```

### 步骤 2: 安装前端依赖
```bash
cd frontend
npm install
```

### 步骤 3: 检查问题文件
确保 `questions/` 目录下有有效的 Markdown 文件：
```bash
ls -la questions/
cat questions/example.md
```

### 步骤 4: 启动后端服务
```bash
cd /Users/qiuxin/code/Self-improvement
go run web_server.go
```
服务将在 `http://localhost:5000` 启动

### 步骤 5: 启动前端开发服务器 (另开一个终端)
```bash
cd /Users/qiuxin/code/Self-improvement/frontend
npm run dev
```
前端将在 `http://localhost:3000` 启动

## 测试计划

### 测试用例 1: 用户注册和登录
1. 访问 `http://localhost:3000`
2. 点击注册，创建新用户
   - 用户名: testuser
   - 密码: testpass123
3. 使用相同凭据登录
4. 验证登录成功，获取 token

### 测试用例 2: 初始化知识库
1. 登录后，点击"初始化知识库"
2. 验证成功导入 questions 目录中的问题
3. 检查统计信息显示正确

### 测试用例 3: 学习流程
1. 查看待复习问题列表
2. 点击"查看答案"
3. 选择记忆程度反馈 (熟练/一般/忘记/完全忘记)
4. 验证问题正确更新，加载下一个问题
5. 完成所有问题后显示完成状态

### 测试用例 4: 问题管理
1. 在学习过程中点击"删除问题"
2. 确认删除后，验证问题从列表中移除
3. 验证统计信息正确更新

### 测试用例 5: 查看统计
1. 访问仪表板/主页
2. 验证以下统计显示正确：
   - 总问题数
   - 待复习数
   - 正确率
   - 总复习次数

### 测试用例 6: 间隔重复算法
1. 初始化后，立即再次加载问题
2. 验证刚刚回答的问题仍然在列表中（因为设置了短期复习间隔）
3. 选择"熟练"回答后，验证下次复习时间被延长

### 测试用例 7: API 直接测试 (可选)
```bash
# 注册
curl -X POST http://localhost:5000/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test123"}'

# 登录
curl -X POST http://localhost:5000/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"test123"}'

# 获取统计 (需要 token)
curl -X GET http://localhost:5000/api/stats \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 预期结果

### 成功标准
- [ ] 后端服务正常启动，端口 5000 可访问
- [ ] 前端服务正常启动，端口 3000 可访问
- [ ] 用户能够成功注册和登录
- [ ] 能够成功初始化知识库并导入问题
- [ ] 学习流程完整，查看答案、提交反馈、切换问题正常
- [ ] 能够删除问题
- [ ] 统计信息正确显示和更新
- [ ] JWT 认证正常工作
- [ ] 数据正确存储在 SQLite 数据库中

### 可能遇到的问题及解决方案

1. **端口被占用**
   - 修改 `web_server.go` 中的 `PORT` 环境变量
   - 修改 `frontend/vite.config.ts` 中的 `server.port`

2. **CORS 错误**
   - 确认 `web_server.go` 中已配置 CORS 允许所有来源
   - 检查前端代理配置正确

3. **数据库错误**
   - 删除 `data/app.db` 重新初始化
   - 检查 `data` 目录权限

4. **问题未导入**
   - 确认 `questions/` 目录下有 `.md` 文件
   - 确认文件格式正确 (`# q` 和 `# a` 标记)

5. **前端编译错误**
   - 运行 `rm -rf frontend/node_modules && npm install` 重新安装依赖
   - 检查 Node.js 版本是否兼容

## 清理测试数据 (可选)
```bash
# 删除数据库
rm data/app.db

# 删除前端构建产物
rm -rf frontend/dist
```

## 生产部署建议

1. **构建前端**
```bash
cd frontend
npm run build
```

2. **配置后端服务静态文件**
   将 `frontend/dist` 内容复制到 Go 项目静态目录

3. **使用环境变量**
```bash
export DATABASE_PATH=/path/to/data.db
export PORT=8080
```

4. **使用反向代理** (如 Nginx)
   - 配置 SSL
   - 配置 gzip 压缩
