# E2E 演示功能验证 SOP（AI 可执行）

## 概述

验证 KnowLoop 端到端演示流程：访客通过「一键体验」登录 → 浏览仪表盘 → 完成一轮学习 → 查看统计更新。

**目标**: 确保 Demo 体验路径完整可用，无断点、无报错。  
**执行者**: AI Agent（通过 curl + HTML 解析）  
**前置**: 服务已部署且运行中，demo 账户可登录  

## 配置

```yaml
base_url: http://115.190.235.149:4430
demo_username: demo
demo_password: demo123
check_timeout_ms: 5000
```

## 验证步骤

### Phase 1: 环境就绪检查

#### Step 1.1 服务可达性

```bash
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" --connect-timeout 5 "{{base_url}}/")
```

| 条件 | 判定 |
|------|------|
| `HTTP_CODE == 200` | ✅ PASS |
| `HTTP_CODE != 200` | ❌ FAIL — 服务不可达，终止验证 |

#### Step 1.2 Demo 账户可登录

```bash
LOGIN_RESP=$(curl -s --connect-timeout 5 -X POST "{{base_url}}/api/login" \
  -H 'Content-Type: application/json' \
  -d '{"username":"{{demo_username}}","password":"{{demo_password}}"}')

TOKEN=$(echo "$LOGIN_RESP" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])" 2>/dev/null)
USER_ID=$(echo "$LOGIN_RESP" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['user_id'])" 2>/dev/null)
```

| 条件 | 判定 |
|------|------|
| `LOGIN_RESP.success == true` | ✅ PASS |
| `TOKEN` 非空 | ✅ PASS |
| `USER_ID` 为整数 | ✅ PASS |
| 任一不满足 | ❌ FAIL — Demo 账户异常 |

### Phase 2: 数据完整性检查

#### Step 2.1 统计数据

```bash
STATS=$(curl -s --connect-timeout 5 "{{base_url}}/api/stats" \
  -H "Authorization: Bearer $TOKEN")
```

提取字段：
- `total_questions`: 总题目数
- `due_questions`: 待复习数
- `total_reviews`: 总复习次数
- `accuracy`: 正确率

| 条件 | 判定 |
|------|------|
| `total_questions >= 8` | ✅ PASS |
| `total_questions < 8` | ❌ FAIL — Demo 题目未正确预置 |

#### Step 2.2 待复习题目

```bash
DUE=$(curl -s --connect-timeout 5 "{{base_url}}/api/due-questions" \
  -H "Authorization: Bearer $TOKEN")
```

提取：`questions[]` 数组，每个元素包含 `id`, `question`, `answer`

| 条件 | 判定 |
|------|------|
| `questions` 是数组（即使为空） | ✅ PASS |
| `questions` 不存在 | ❌ FAIL — API 返回格式异常 |

如果 `questions` 非空，记录 `QUESTION_ID = questions[0].id` 供后续使用。

#### Step 2.3 分类列表

```bash
CATS=$(curl -s --connect-timeout 5 "{{base_url}}/api/categories" \
  -H "Authorization: Bearer $TOKEN")
```

| 条件 | 判定 |
|------|------|
| `categories` 是数组且至少包含 1 个分类 | ✅ PASS |
| `categories` 为空或不存在 | ⚠️ WARN — 分类功能可能未正常工作 |

### Phase 3: 核心学习流程

#### Step 3.1 获取一道待复习题目

若有待复习题目（`due_questions > 0`），执行完整反馈流程。若无（`due_questions == 0`），跳过 Phase 3，直接进入 Phase 4。

```bash
# 取第一道待复习题目
QID=$(echo "$DUE" | python3 -c "import sys,json; qs=json.load(sys.stdin)['data']['questions']; print(qs[0]['id'] if qs else '')" 2>/dev/null)
```

#### Step 3.2 提交反馈（熟练）

```bash
FB=$(curl -s --connect-timeout 5 -X POST "{{base_url}}/api/update-review" \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d "{\"question_id\":\"$QID\",\"feedback\":1}")
```

| 条件 | 判定 |
|------|------|
| `FB.success == true` | ✅ PASS |
| `FB.success == false` | ❌ FAIL — 反馈提交失败 |

#### Step 3.3 验证统计更新

```bash
STATS_AFTER=$(curl -s --connect-timeout 5 "{{base_url}}/api/stats" \
  -H "Authorization: Bearer $TOKEN")
```

提取 `total_reviews_after`。

| 条件 | 判定 |
|------|------|
| `total_reviews_after == total_reviews + 1` | ✅ PASS — 统计正确递增 |
| `total_reviews_after != total_reviews + 1` | ❌ FAIL — 统计未更新 |

#### Step 3.4 再次提交反馈（忘记）

换一道题测试不同反馈级别：

```bash
QID2=$(echo "$DUE" | python3 -c "import sys,json; qs=json.load(sys.stdin)['data']['questions']; print(qs[1]['id'] if len(qs)>1 else '')" 2>/dev/null)

FB2=$(curl -s --connect-timeout 5 -X POST "{{base_url}}/api/update-review" \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d "{\"question_id\":\"$QID2\",\"feedback\":3}")
```

| 条件 | 判定 |
|------|------|
| `FB2.success == true` | ✅ PASS |

### Phase 4: 新用户注册与隔离验证

#### Step 4.1 注册新用户

```bash
TEST_USER="e2e_test_$(date +%s)"
REG=$(curl -s --connect-timeout 5 -X POST "{{base_url}}/api/register" \
  -H 'Content-Type: application/json' \
  -d "{\"username\":\"$TEST_USER\",\"password\":\"test1234\"}")

NEW_TOKEN=$(echo "$REG" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])" 2>/dev/null)
```

| 条件 | 判定 |
|------|------|
| `REG.success == true` | ✅ PASS |
| `NEW_TOKEN` 非空 | ✅ PASS |

#### Step 4.2 验证数据隔离

新用户应该看到独立的数据，不与 demo 用户混淆。

```bash
NEW_STATS=$(curl -s --connect-timeout 5 "{{base_url}}/api/stats" \
  -H "Authorization: Bearer $NEW_TOKEN")
```

| 条件 | 判定 |
|------|------|
| `new_total_questions == 0` | ✅ PASS — 新用户无数据，隔离正常 |
| `new_total_questions > 0` | ❌ FAIL — 数据隔离异常 |

#### Step 4.3 新用户初始化知识库

```bash
INIT=$(curl -s --connect-timeout 10 -X POST "{{base_url}}/api/init" \
  -H "Authorization: Bearer $NEW_TOKEN")
```

| 条件 | 判定 |
|------|------|
| `INIT.success` 有响应（成功或失败均可） | ✅ PASS — 初始化接口正常 |

### Phase 5: 前端资源完整性

#### Step 5.1 检查前端资源文件

```bash
INDEX=$(curl -s --connect-timeout 5 "{{base_url}}/")
```

检查 `index.html` 中引用的资源：

| 条件 | 判定 |
|------|------|
| HTML 包含 `<div id="app">` | ✅ PASS — Vue 挂载点存在 |
| HTML 包含 `<script` 和 `src=` | ✅ PASS — JS 资源引用存在 |
| HTML 响应体 < 100 字节 | ❌ FAIL — 前端资源可能缺失 |

#### Step 5.2 检查关键 JS/CSS 文件可访问

```bash
# 从 index.html 中提取第一个 JS 资源路径并验证
JS_PATH=$(echo "$INDEX" | grep -oP 'src="([^"]+\.js)"' | head -1 | sed 's/src="//;s/"//')
if [ -n "$JS_PATH" ]; then
  JS_CODE=$(curl -s -o /dev/null -w "%{http_code}" --connect-timeout 5 "{{base_url}}$JS_PATH")
fi
```

| 条件 | 判定 |
|------|------|
| `JS_CODE == 200` | ✅ PASS |
| `JS_CODE != 200` | ❌ FAIL — 前端 JS 资源 404 |

### Phase 6: 错误处理验证

#### Step 6.1 无效 Token

```bash
ERR1=$(curl -s --connect-timeout 5 "{{base_url}}/api/stats" \
  -H "Authorization: Bearer invalidtoken123")
```

| 条件 | 判定 |
|------|------|
| 返回 401 或 error 响应 | ✅ PASS — 认证拦截正常 |
| 返回 200 + 数据 | ❌ FAIL — 认证绕过漏洞 |

#### Step 6.2 无效反馈值

```bash
# 构造一个属于 demo 用户的任意 question_id
ANY_QID=$(curl -s --connect-timeout 5 "{{base_url}}/api/due-questions" \
  -H "Authorization: Bearer $TOKEN" | \
  python3 -c "import sys,json; print(json.load(sys.stdin)['data']['questions'][0]['id'])" 2>/dev/null || echo "nonexistent")

ERR2=$(curl -s --connect-timeout 5 -X POST "{{base_url}}/api/update-review" \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d "{\"question_id\":\"$ANY_QID\",\"feedback\":99}")
```

| 条件 | 判定 |
|------|------|
| 返回错误（400 或 `"success":false`） | ✅ PASS — 参数校验正常 |
| 返回 `"success":true` | ❌ FAIL — 缺少参数校验 |

---

## 判定矩阵

| Phase | 步骤 | 权重 | 阻塞？ |
|-------|------|------|--------|
| 1.1 | 服务可达 | MUST | 是 |
| 1.2 | Demo 登录 | MUST | 是 |
| 2.1 | 统计数据 | MUST | 否 |
| 2.2 | 待复习题目 | MUST | 否 |
| 2.3 | 分类列表 | SHOULD | 否 |
| 3.1-3.4 | 学习反馈流程 | MUST* | 否 |
| 4.1 | 新用户注册 | MUST | 否 |
| 4.2 | 数据隔离 | MUST | 否 |
| 4.3 | 初始化知识库 | SHOULD | 否 |
| 5.1 | 前端 HTML | MUST | 否 |
| 5.2 | JS 资源 | MUST | 否 |
| 6.1 | 无效 Token | SHOULD | 否 |
| 6.2 | 无效参数 | SHOULD | 否 |

> \* Phase 3 仅在 `due_questions > 0` 时执行，`due_questions == 0` 时自动跳过。

**最终判定**：
- 🟢 **PASS**: 所有 MUST 项通过
- 🟡 **WARN**: 所有 MUST 通过但有 SHOULD 失败
- 🔴 **FAIL**: 任一 MUST 项失败

---

## AI 执行指令

当 AI Agent 执行此 SOP 时：

1. 按 Phase 顺序执行，Phase 1 阻塞项失败则终止
2. 每步记录 `[PASS/FAIL/WARN/SKIP]` 状态
3. 收集 `total_questions`, `due_questions`, `total_reviews`, `accuracy` 作为基线数据
4. 若 Phase 3 执行了反馈，对比前后 stats 确认递增
5. 最终输出判定矩阵汇总表
6. 所有 API 调用设置 `--connect-timeout 5` 避免长时间阻塞

### 期望输出格式

```
## E2E 验证报告

**时间**: 2026-05-30 12:00:00
**目标**: http://115.190.235.149:4430

| Phase | 步骤 | 结果 | 详情 |
|-------|------|------|------|
| 1.1 | 服务可达 | ✅ | HTTP 200 |
| 1.2 | Demo 登录 | ✅ | Token 获取成功 |
| 2.1 | 统计数据 | ✅ | 总题8 待复习0 复习8次 |
| 2.2 | 待复习 | ✅ | API 响应正常 |
| 2.3 | 分类列表 | ✅ | 1 个分类 |
| 3.x | 学习流程 | ⏭️ | 无待复习题目，跳过 |
| 4.1 | 新用户注册 | ✅ | 注册成功 |
| 4.2 | 数据隔离 | ✅ | 新用户 total_questions=0 |
| 4.3 | 初始化 | ✅ | API 响应正常 |
| 5.1 | 前端 HTML | ✅ | 挂载点存在 |
| 5.2 | JS 资源 | ✅ | 资源可访问 |
| 6.1 | 无效 Token | ✅ | 认证拦截正常 |
| 6.2 | 无效参数 | ✅ | 参数校验正常 |

**结论**: 🟢 PASS (11/11 MUST 通过, 0/2 SHOULD 通过)
```
