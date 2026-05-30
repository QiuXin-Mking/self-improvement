# 部署验证 SOP

## 概述

KnowLoop 部署到火山云服务器的标准操作流程。涵盖编译、部署、验证和回滚。

**目标服务器**：huoshan (115.190.235.149:22)  
**部署路径**：`/root/spaced-repetition`  
**服务端口**：4430  

## 前置条件

- [ ] 本地可 SSH 连接到 huoshan（`ssh huoshan` 成功）
- [ ] 本地已安装 Node.js（前端编译需要）
- [ ] 本地已安装 Go 1.21+
- [ ] 代码已通过本地编译验证（`go build ./internal/...`）

## 1. 部署流程

### 1.1 编译并部署

```bash
cd /Users/qiuxin/code/Self-improvement
bash scripts/deploy_to_huoshan.sh
```

脚本自动完成：
1. 本地编译前端（`npm run build`）
2. 打包 `web_server.go` + `internal/` + `go.mod` + `frontend/dist`
3. 上传到服务器 `/tmp/spaced-repetition-build`
4. 服务器端编译 Go 二进制
5. 停止旧服务 → 备份数据 → 部署新版本 → 启动服务
6. 用户数据（`data/`、`.env`、`questions/`）自动保留

### 1.2 仅编译不部署

```bash
cd /Users/qiuxin/code/Self-improvement
BUILD_TARGET=linux bash scripts/compile.sh
```

产物在 `bin/web_server` 和 `frontend/dist/`。

## 2. 验证清单

部署完成后按顺序执行以下验证：

### 2.1 服务存活检查

```bash
# 检查进程是否运行
ssh huoshan 'cat /root/spaced-repetition/.app.pid && ps -p $(cat /root/spaced-repetition/.app.pid)'

# 查看启动日志（确认无 panic/error）
ssh huoshan 'tail -20 /root/spaced-repetition/app.log'
```

✅ **通过标准**：进程 PID 存在，日志无 `panic`、`FATAL` 关键字。

### 2.2 健康检查

```bash
# 首页可访问
curl -s -o /dev/null -w "%{http_code}" http://115.190.235.149:4430/

# API 响应正常（注册接口应返回 400，说明路由正常）
curl -s -X POST http://115.190.235.149:4430/api/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"nobody","password":"xxxxxx"}'
```

✅ **通过标准**：首页返回 200，API 有正常 JSON 响应（非 connection refused）。

### 2.3 Demo 账户验证

```bash
# 1. Demo 登录
TOKEN=$(curl -s -X POST http://115.190.235.149:4430/api/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"demo","password":"demo123"}' | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])")

echo "Token: ${TOKEN:0:20}..."

# 2. 统计数据
curl -s http://115.190.235.149:4430/api/stats \
  -H "Authorization: Bearer $TOKEN"

# 3. 待复习题目
curl -s http://115.190.235.149:4430/api/due-questions \
  -H "Authorization: Bearer $TOKEN"

# 4. 完整学习流程：反馈一道题
QID=$(curl -s http://115.190.235.149:4430/api/due-questions \
  -H "Authorization: Bearer $TOKEN" | \
  python3 -c "import sys,json; print(json.load(sys.stdin)['data']['questions'][0]['id'])")

curl -s -X POST http://115.190.235.149:4430/api/update-review \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer $TOKEN" \
  -d "{\"question_id\":\"$QID\",\"feedback\":1}"

# 5. 验证统计已更新
curl -s http://115.190.235.149:4430/api/stats \
  -H "Authorization: Bearer $TOKEN"
```

✅ **通过标准**：

| 检查项 | 首次部署预期 | 非首次部署预期 |
|--------|-------------|----------------|
| Demo 登录 | `"success":true` | `"success":true` |
| 总题目数 | `total_questions: 8` | `total_questions ≥ 8` |
| 待复习数 | `due_questions: 8` | `due_questions ≥ 0` |
| 反馈提交 | `"success":true` | `"success":true` |
| 反馈后统计 | `total_reviews ≥ 1` | `total_reviews 递增 1` |

> 💡 **首次部署** vs **非首次部署**：首次部署时 Demo 题目从未被复习，`due_questions=8, total_reviews=0`。
> 非首次部署时由于之前的测试或使用，部分题目可能已被复习，`due_questions` 可能为 0，属正常现象。

### 2.4 前端页面检查

在浏览器中依次访问：

| 页面 | URL | 检查项 |
|------|-----|--------|
| 登录页 | `http://115.190.235.149:4430/` | 「🎯 一键体验」按钮可见，样式正常 |
| 一键体验 | 点击体验按钮 | 自动登录，跳转 dashboard |
| 仪表盘 | `/dashboard` | 统计卡片（8/8/0%）正常显示，按钮可用 |
| 复习页 | `/learn` | 问题卡片、答案卡片、反馈按钮正常 |
| 注册页 | `/register` | 表单正常渲染 |

✅ **通过标准**：页面无白屏、无 JS 报错、组件渲染完整。

### 2.5 新用户注册验证

```bash
TEST_USER="soptest_$(date +%s)"

# 1. 注册
curl -s -X POST http://115.190.235.149:4430/api/register \
  -H 'Content-Type: application/json' \
  -d "{\"username\":\"$TEST_USER\",\"password\":\"test1234\"}"

# 2. 登录
curl -s -X POST http://115.190.235.149:4430/api/login \
  -H 'Content-Type: application/json' \
  -d "{\"username\":\"$TEST_USER\",\"password\":\"test1234\"}"
```

✅ **通过标准**：注册成功返回 token，登录成功返回 token。

## 3. 一键验证脚本

将以上 API 验证整合为自动化脚本：

```bash
#!/bin/bash
# 文件: docs/operations/verify-deploy.sh
# 用法: bash docs/operations/verify-deploy.sh

set -e
BASE="http://115.190.235.149:4430"
PASS=0
FAIL=0

check() {
    local desc="$1"
    local expected="$2"
    local actual="$3"
    if echo "$actual" | grep -q "$expected"; then
        echo "  ✅ $desc"
        PASS=$((PASS+1))
    else
        echo "  ❌ $desc (expected: $expected)"
        echo "     got: $actual"
        FAIL=$((FAIL+1))
    fi
}

echo "========================================="
echo "KnowLoop 部署验证"
echo "========================================="
echo ""

# 1. Service health
echo "1. 服务健康检查"
HTTP=$(curl -s -o /dev/null -w "%{http_code}" "$BASE/")
check "首页返回 200" "200" "$HTTP"

# 2. Demo login
echo "2. Demo 账户"
LOGIN=$(curl -s -X POST "$BASE/api/login" \
    -H 'Content-Type: application/json' \
    -d '{"username":"demo","password":"demo123"}')
check "Demo 登录成功" '"success":true' "$LOGIN"

TOKEN=$(echo "$LOGIN" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])" 2>/dev/null || echo "")

if [ -n "$TOKEN" ]; then
    # 3. Stats
    echo "3. 统计数据"
    STATS=$(curl -s "$BASE/api/stats" -H "Authorization: Bearer $TOKEN")
    check "总题目数 8" '"total_questions":8' "$STATS"

    # 4. Due questions
    echo "4. 待复习题目"
    DUE=$(curl -s "$BASE/api/due-questions" -H "Authorization: Bearer $TOKEN")
    check "有待复习题目" '"total":' "$DUE"

    # 5. Feedback
    echo "5. 反馈流程"
    QID=$(echo "$DUE" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['questions'][0]['id'])" 2>/dev/null || echo "")
    if [ -n "$QID" ]; then
        FB=$(curl -s -X POST "$BASE/api/update-review" \
            -H 'Content-Type: application/json' \
            -H "Authorization: Bearer $TOKEN" \
            -d "{\"question_id\":\"$QID\",\"feedback\":1}")
        check "反馈提交成功" '"success":true' "$FB"
    fi

    # 6. Verify stats updated
    STATS2=$(curl -s "$BASE/api/stats" -H "Authorization: Bearer $TOKEN")
    check "复习次数已更新" '"total_reviews":1' "$STATS2"
fi

# 7. New registration
echo "6. 新用户注册"
TEST_USER="verify_$(date +%s)"
REG=$(curl -s -X POST "$BASE/api/register" \
    -H 'Content-Type: application/json' \
    -d "{\"username\":\"$TEST_USER\",\"password\":\"test1234\"}")
check "新用户注册成功" '"success":true' "$REG"

echo ""
echo "========================================="
echo "结果: ✅ $PASS 通过  ❌ $FAIL 失败"
echo "========================================="
```

## 4. 日志查看

```bash
# 实时日志
ssh huoshan 'tail -f /root/spaced-repetition/app.log'

# 最近 50 行
ssh huoshan 'tail -50 /root/spaced-repetition/app.log'

# 搜索错误
ssh huoshan 'grep -i "error\|panic\|fatal" /root/spaced-repetition/app.log'
```

## 5. 回滚流程

如验证不通过，回滚到上一个备份：

```bash
# 查看备份列表
ssh huoshan 'ls -d /root/spaced-repetition.backup.*'

# 回滚（替换为实际备份目录名）
ssh huoshan '
  cd /root/spaced-repetition && bash stop.sh
  rm -rf /root/spaced-repetition
  cp -r /root/spaced-repetition.backup.20260530_XXXXXX /root/spaced-repetition
  cd /root/spaced-repetition && bash start.sh
'
```

## 6. 常见问题

| 问题 | 原因 | 解决 |
|------|------|------|
| Demo 登录失败 | 数据库未包含 demo 用户 | 重启服务，`seedDemoUser` 会自动创建 |
| 前端白屏 | 前端资源未正确部署 | 检查 `dist/` 目录是否完整上传 |
| 502/connection refused | 服务未启动 | `ssh huoshan 'cat /root/spaced-repetition/app.log'` 查看错误 |
| Go 编译失败 | 服务器缺少 Go 环境 | 部署脚本会自动安装 golang-1.22 |
| 数据丢失 | 部署时数据库被覆盖 | 部署脚本已做数据备份，检查 `/tmp` 备份 |
