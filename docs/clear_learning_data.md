# 清除复习进度数据

本指南说明如何清除项目的复习进度数据，用于测试或重新开始学习。

## 存储位置

项目中有两种数据存储方式：

| 数据源 | 文件路径 | 说明 |
|--------|----------|------|
| JSON 文件 | `data/learning_data.json` | 旧版本的数据存储方式 |
| SQLite 数据库 | `data/app.db` | 新版本的数据存储方式（含用户认证） |

**注意**: 两种数据源可能同时存在，清理脚本会同时处理两者。

## 使用方法

### 方法一：使用 Go 脚本（推荐）

```bash
cd /Users/qiuxin/code/Self-improvement
go run scripts/clear_learning_data.go
```

**优点**:
- 无需额外依赖
- 逻辑更可靠
- 统一的处理方式

### 方法二：使用 Shell 脚本

```bash
cd /Users/qiuxin/code/Self-improvement
./scripts/clear_learning_data.sh
```

**依赖**:
- `jq` - 用于处理 JSON 文件
- `sqlite3` - 用于处理 SQLite 数据库

**安装依赖**:
```bash
# macOS
brew install jq sqlite3

# Ubuntu/Debian
sudo apt-get install jq sqlite3

# CentOS/RHEL
sudo yum install jq sqlite3
```

## 功能说明

### 执行的操作

脚本会执行以下操作：

1. **备份原始文件**
   - JSON 文件备份到 `data/learning_data.json.bak`
   - SQLite 文件备份到 `data/app.db.bak`

2. **清除复习进度数据**
   - 重置所有问题的学习进度
   - 保留问题内容和创建时间

3. **显示统计结果**
   - 显示清除的 JSON 数据数量
   - 显示清除的 SQLite 数据数量
   - 显示总计清除数量

### 清除的字段

以下字段会被重置：

| 字段 | 重置值 | 说明 |
|------|---------|------|
| `level` | 1 | 重置为初始记忆等级 |
| `next_review` | 当前时间 | 立即可复习 |
| `review_count` | 0 | 清除复习次数统计 |
| `correct_count` | 0 | 清除正确次数统计 |
| `last_reviewed` | 空值 | 清除最后复习时间 |

### 保留的内容

以下内容会被保留：

| 内容 | 说明 |
|------|------|
| 问题文本 | `question_text` |
| 答案文本 | `answer_text` |
| 创建时间 | `created_at` |
| 来源文件 | `source` |
| 用户关联 | `user_id` |

## 示例输出

```
✓ 成功清除 JSON 数据中 1 个问题的复习进度
✓ 成功清除 SQLite 数据库中 10 个问题的复习进度

总计: 成功清除 11 个问题的复习进度数据
```

## 注意事项

1. **自动备份**: 脚本会自动创建备份文件，数据丢失可以恢复
2. **用户隔离**: SQLite 数据库支持多用户，清除操作会影响所有用户
3. **服务停止**: 清除数据前建议停止 web 服务，避免并发问题

## 恢复数据

如果需要恢复数据：

```bash
# 恢复 JSON 数据
cp data/learning_data.json.bak data/learning_data.json

# 恢复 SQLite 数据
cp data/app.db.bak data/app.db
```
