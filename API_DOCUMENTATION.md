# 后端 API 接口文档

## 基本信息

- **Base URL**: `http://localhost:8000/api` (开发环境)
- **认证方式**: Bearer Token (JWT)
- **响应格式**: JSON

## 通用响应格式

所有接口返回以下格式：

```json
{
  "success": boolean,
  "error": string | null,
  "data": any | null,
  "message": string | null
}
```

---

## 认证接口

### 1. 用户注册

**接口**: `POST /register`

**说明**: 创建新用户账号，成功后自动登录并返回 token

**请求参数**:
```json
{
  "username": "string",  // 用户名，3-32位
  "password": "string"   // 密码，6-128位
}
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "token": "string",      // JWT token
    "user_id": number,      // 用户ID
    "username": "string"    // 用户名
  },
  "message": "Registration successful"
}
```

**错误响应**:
- `400`: Invalid request format
- `409`: Username already exists
- `500`: Server error

---

### 2. 用户登录

**接口**: `POST /login`

**说明**: 用户登录，成功后返回 JWT token

**请求参数**:
```json
{
  "username": "string",  // 用户名，3-32位
  "password": "string"   // 密码，6-128位
}
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "token": "string",      // JWT token
    "user_id": number,      // 用户ID
    "username": "string"    // 用户名
  },
  "message": "Login successful"
}
```

**错误响应**:
- `400`: Invalid request format
- `401`: Invalid username or password
- `500`: Server error

---

### 3. 获取用户资料

**接口**: `GET /profile`

**说明**: 获取当前登录用户的信息

**请求头**:
```
Authorization: Bearer <token>
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "user_id": number,    // 用户ID
    "username": "string"  // 用户名
  }
}
```

**错误响应**:
- `401`: Authorization header is required / Token is invalid
- `500`: Server error

---

## 学习统计接口

### 4. 获取学习统计

**接口**: `GET /stats`

**说明**: 获取当前用户的学习统计数据

**请求头**:
```
Authorization: Bearer <token>
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_questions": number,   // 总问题数
      "due_questions": number,     // 今日待复习问题数
      "total_reviews": number,     // 总复习次数
      "total_correct": number,     // 总正确次数
      "accuracy": number          // 正确率 (百分比)
    }
  }
}
```

**错误响应**:
- `401`: Unauthorized
- `500`: Server error

---

### 5. 获取待复习问题

**接口**: `GET /due-questions`

**说明**: 获取当前需要复习的问题列表

**请求头**:
```
Authorization: Bearer <token>
```

**成功响应** (有需要复习的问题):
```json
{
  "success": true,
  "data": {
    "questions": [
      {
        "id": number,              // 问题ID
        "question": "string",      // 问题内容
        "answer": "string",        // 答案内容
        "review_count": number,    // 已复习次数
        "correct_count": number,    // 正确次数
        "source": "string"        // 来源文件
      }
    ],
    "total": number              // 问题总数
  }
}
```

**成功响应** (没有需要复习的问题):
```json
{
  "success": true,
  "data": {
    "questions": [],
    "message": "太棒了！今天没有需要复习的问题！"
  }
}
```

**成功响应** (知识库为空):
```json
{
  "success": false,
  "error": "知识库为空！请先初始化知识库。",
  "data": {
    "needs_init": true
  }
}
```

**错误响应**:
- `401`: Unauthorized
- `500`: Server error

---

### 6. 提交复习反馈

**接口**: `POST /update-review`

**说明**: 提交对某个问题的复习反馈，更新学习进度

**请求头**:
```
Authorization: Bearer <token>
```

**请求参数**:
```json
{
  "question_id": "string",  // 问题ID (如 "q_abc123")
  "feedback": number       // 反馈级别：1=熟练, 2=一般, 3=忘记, 4=完全忘记
}
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_questions": number,
      "due_questions": number,
      "total_reviews": number,
      "total_correct": number,
      "accuracy": number
    }
  }
}
```

**错误响应**:
- `400`: Invalid request format
- `401`: Unauthorized
- `404`: Question not found or update failed
- `500`: Server error

---

### 7. 删除问题

**接口**: `POST /delete-question`

**说明**: 从知识库中删除某个问题

**请求头**:
```
Authorization: Bearer <token>
```

**请求参数**:
```json
{
  "question_id": "string"  // 问题ID (如 "q_abc123")
}
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_questions": number,
      "due_questions": number,
      "total_reviews": number,
      "total_correct": number,
      "accuracy": number
    }
  }
}
```

**错误响应**:
- `400`: Invalid request format
- `401`: Unauthorized
- `404`: Question not found
- `500`: Server error

---

### 8. 初始化知识库

**接口**: `POST /init`

**说明**: 重新初始化知识库，从配置的目录中读取 Markdown 问题文件并导入到数据库

**功能说明**:
1. 读取 `question_input` (Windows) 或 `question_input_linux` (Linux) 配置文件
2. 从配置的目录中递归扫描所有 `.md` 文件
3. 解析 Markdown 文件中的问题（以 `# q` 开头）和答案（以 `# a` 开头）
4. 去除重复的问题
5. 跳过已存在的问题
6. 将新问题添加到当前用户的知识库

**请求头**:
```
Authorization: Bearer <token>
```

**成功响应**:
```json
{
  "success": true,
  "data": {
    "message": "成功导入 X 个新问题到知识库！",
    "imported": number,     // 新导入的问题数
    "skipped": number,      // 跳过的已存在问题数
    "duplicates": number,   // 发现的重复问题数
    "stats": {
      "total_questions": number,
      "due_questions": number,
      "total_reviews": number,
      "total_correct": number,
      "accuracy": number
    }
  }
}
```

**错误响应**:
- `400`: 没有找到任何问题！请确保配置的目录下有 .md 文件。
- `401`: Unauthorized
- `500`: Server error

---

## 问题文件格式

Markdown 问题文件需要遵循以下格式：

```markdown
# q
这是第一个问题？

# a
这是第一个问题的答案。

# q
这是第二个问题？

# a
这是第二个问题的答案。
```

## 配置文件说明

### question_input (Windows) / question_input_linux (Linux)

配置文件内容格式：每行一个目录路径

```
/path/to/questions1
/path/to/questions2
```

如果配置文件不存在或为空，默认使用 `questions/` 目录。

## 间隔重复算法说明

### 记忆级别
- **Level 1 (熟练)**: 168 小时 (7 天)
- **Level 2 (一般)**: 72 小时 (3 天)
- **Level 3 (忘记)**: 24 小时 (1 天)
- **Level 4 (完全忘记)**: 2 小时

### 复习间隔计算
根据反馈级别和正确率计算下次复习时间：
- 反馈级别对应的间隔 × 正确率乘数
- 正确率 > 80% 时，额外乘以 1.2 倍
- 答对 3 次以上后，记忆等级提升
- 答错后，记忆等级下降

### 问题 ID
问题 ID 使用问题文本的哈希值生成，确保相同问题始终有相同 ID，防止重复。

---

## 错误码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 / Token 无效 |
| 404 | 资源未找到 |
| 409 | 资源冲突（如用户名已存在） |
| 500 | 服务器内部错误 |

---

## 使用示例

### 1. 注册新用户
```bash
curl -X POST http://localhost:8000/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
```

### 2. 登录
```bash
curl -X POST http://localhost:8000/api/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
```

### 3. 获取统计（需要 token）
```bash
curl -X GET http://localhost:8000/api/stats \
  -H "Authorization: Bearer <token>"
```

### 4. 初始化知识库
```bash
curl -X POST http://localhost:8000/api/init \
  -H "Authorization: Bearer <token>"
```

---

## 重新初始化知识库功能说明

### 什么时候使用"重新初始化知识库"？

1. **首次使用**：刚注册账号后，需要导入问题文件到知识库
2. **更新问题**：当你在 Markdown 文件中添加或修改问题后，需要重新导入
3. **添加新问题文件**：在配置的目录中添加了新的问题文件
4. **清空重建**：想要重新开始学习计划

### 该功能会做什么？

1. 扫描配置文件中指定的目录
2. 解析所有 `.md` 文件中的问题
3. 去除重复问题
4. 只添加当前用户还没有的问题
5. 返回导入统计信息

### 注意事项

- 不会删除已有的问题，只会添加新问题
- 已存在的问题会被跳过，不会重复添加
- 相同文本的问题会被去重
- 需要确保配置的目录下有符合格式的 Markdown 文件
