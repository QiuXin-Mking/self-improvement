# Fix deploy_to_huoshan.sh Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fix 4 bugs in deploy_to_huoshan.sh that prevent successful remote deployment to huoshan server.

**Architecture:** The deploy flow is: local frontend build → package Go source + frontend dist → scp to server → server-side Go compile → deploy binary + start service. The script has issues in both the packaging and server-side compilation stages.

**Tech Stack:** Bash, Go 1.19, SSH/SCP

---

## Bug Analysis

1. **Server compile fails**: `cp *.go` copies both `cli_server.go` and `web_server.go` (both `package main`, both have `main()`). Server-side `go build -o web_server *.go` tries to compile them together → duplicate symbol error.

2. **Server start.sh missing env vars**: `.env` file is created but never sourced. web_server needs PORT, JWT_SECRET, DATABASE_PATH from environment.

3. **Local Go build is wasted**: `bash compile.sh` builds a macOS binary that isn't uploaded (server recompiles from source). Only `npm run build` is needed from compile.sh.

4. **handlers/ is dead code but included**: `handlers/auth.go` has broken imports (`internal/middleware` missing module prefix) and is never imported by web_server.go. Including it causes server compile failure.

### Additional Issue

5. **Server build command wrong**: `go build -o web_server *.go` builds in file mode, not package mode. Should be `go build -o web_server .`

---

### Task 1: Fix deploy_to_huoshan.sh — source file selection and build command

**Files:**
- Modify: `deploy_to_huoshan.sh`

- [ ] **Step 1: Replace `*.go` glob with explicit file list for Go source copy**

Edit `deploy_to_huoshan.sh`, replace lines 57-67:

```bash
# Old (lines 58-67):
mkdir -p "$DEPLOY_DIR/handlers"
mkdir -p "$DEPLOY_DIR/internal"
mkdir -p "$DEPLOY_DIR/migrations"

# 上传 Go 源代码（仅编译时使用，编译后删除）
cp *.go "$DEPLOY_DIR/" 2>/dev/null || true
cp go.mod go.sum "$DEPLOY_DIR/" 2>/dev/null || true
cp -r handlers/* "$DEPLOY_DIR/handlers/" 2>/dev/null || true
cp -r internal/* "$DEPLOY_DIR/internal/" 2>/dev/null || true
cp migrations/*.sql "$DEPLOY_DIR/migrations/" 2>/dev/null || true
```

With:

```bash
mkdir -p "$DEPLOY_DIR/internal"

# Only upload web_server.go (not cli_server.go — both are package main, can't coexist)
cp web_server.go "$DEPLOY_DIR/"
cp go.mod go.sum "$DEPLOY_DIR/"
cp -r internal/* "$DEPLOY_DIR/internal/"
```

- [ ] **Step 2: Replace server-side compile command to use package mode**

In the build_and_deploy.sh heredoc (around line 119), replace:

```bash
$GO build -o web_server *.go || {
```

With:

```bash
$GO build -o web_server . || {
```

- [ ] **Step 3: Remove handlers/ and migrations/ references from server-side script**

In the build_and_deploy.sh heredoc, remove line 133:

```bash
# Delete this line:
cp migrations/*.sql "$DEPLOY_PATH/migrations/" 2>/dev/null || true
```

- [ ] **Step 4: Replace compile.sh call with frontend-only build**

Replace lines 42-46:

```bash
if [ ! -f "compile.sh" ]; then
    log_error "compile.sh 不存在！"
    exit 1
fi
bash compile.sh
```

With:

```bash
log_info "编译前端..."
cd frontend && npm run build && cd ..
```

- [ ] **Step 5: Commit**

```bash
git add deploy_to_huoshan.sh
git commit -m "fix: deploy_to_huoshan.sh - fix source file selection and build command"
```

---

### Task 2: Fix server-side start.sh — source .env for environment variables

**Files:**
- Modify: `deploy_to_huoshan.sh`

- [ ] **Step 1: Add .env sourcing in server start.sh heredoc**

In the build_and_deploy.sh heredoc, replace the start.sh creation block (lines 136-163):

```bash
# Old start.sh:
cat > "$DEPLOY_PATH/start.sh" << 'EOF'
#!/bin/bash
APP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PID_FILE="$APP_DIR/.app.pid"
LOG_FILE="$APP_DIR/app.log"

if [ -f "$PID_FILE" ]; then
    OLD_PID=$(cat "$PID_FILE")
    if ps -p "$OLD_PID" > /dev/null 2>&1; then
        kill "$OLD_PID" 2>/dev/null || true
        sleep 2
        kill -9 "$OLD_PID" 2>/dev/null || true
    fi
    rm -f "$PID_FILE"
fi

cd "$APP_DIR"
nohup ./bin/web_server > "$LOG_FILE" 2>&1 &
echo $! > "$PID_FILE"
sleep 2

if ps -p $(cat "$PID_FILE") > /dev/null; then
    echo "服务启动成功 (PID: $(cat $PID_FILE))"
else
    echo "服务启动失败，请查看日志: $LOG_FILE"
    exit 1
fi
EOF
```

With:

```bash
cat > "$DEPLOY_PATH/start.sh" << 'EOF'
#!/bin/bash
APP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PID_FILE="$APP_DIR/.app.pid"
LOG_FILE="$APP_DIR/app.log"

# Source .env if it exists
if [ -f "$APP_DIR/.env" ]; then
    set -a
    source "$APP_DIR/.env"
    set +a
fi

if [ -f "$PID_FILE" ]; then
    OLD_PID=$(cat "$PID_FILE")
    if ps -p "$OLD_PID" > /dev/null 2>&1; then
        kill "$OLD_PID" 2>/dev/null || true
        sleep 2
        kill -9 "$OLD_PID" 2>/dev/null || true
    fi
    rm -f "$PID_FILE"
fi

cd "$APP_DIR"
nohup ./bin/web_server > "$LOG_FILE" 2>&1 &
echo $! > "$PID_FILE"
sleep 2

if ps -p $(cat "$PID_FILE") > /dev/null; then
    echo "服务启动成功 (PID: $(cat $PID_FILE))"
else
    echo "服务启动失败，请查看日志: $LOG_FILE"
    exit 1
fi
EOF
```

- [ ] **Step 2: Commit**

```bash
git add deploy_to_huoshan.sh
git commit -m "fix: source .env in server start.sh so web_server gets env vars"
```

---

### Task 3: Local dry-run validation

**Files:**
- None (verification only)

- [ ] **Step 1: Verify deploy_to_huoshan.sh is syntactically valid**

```bash
bash -n deploy_to_huoshan.sh
```

Expected: no output (no syntax errors).

- [ ] **Step 2: Verify SSH connectivity to huoshan**

```bash
ssh -o ConnectTimeout=5 huoshan 'echo "SSH OK: $(hostname)"'
```

Expected: prints "SSH OK: <hostname>".

- [ ] **Step 3: Verify server has required tools**

```bash
ssh huoshan 'go version 2>/dev/null || echo "Go not installed"; node --version 2>/dev/null || echo "Node not installed"'
```

- [ ] **Step 4: Simulate packaging locally (without SCP)**

```bash
rm -rf deploy_temp deploy-package.tar.gz
bash -c '
source deploy_to_huoshan.sh dummy 2>&1 || true
' 2>/dev/null
```

Check that `deploy_temp/` contains only `web_server.go` (not `cli_server.go`), `internal/`, `go.mod`, `go.sum`, `frontend/dist/`, `build_and_deploy.sh`.

```bash
ls deploy_temp/
ls deploy_temp/internal/
ls deploy_temp/frontend/dist/ | head -5
```

- [ ] **Step 5: Verify cli_server.go is NOT in deploy_temp**

```bash
test -f deploy_temp/cli_server.go && echo "FAIL: cli_server.go found" || echo "OK: cli_server.go not in deploy"
```

Expected: "OK: cli_server.go not in deploy"

---

### Task 4: Manual remote test (run by user)

**Files:**
- None

- [ ] **Step 1: Run the deploy script**

```bash
bash deploy_to_huoshan.sh
```

- [ ] **Step 2: Verify service is running**

```bash
ssh huoshan 'cd /root/spaced-repetition && cat .app.pid && ps -p $(cat .app.pid)'
```

- [ ] **Step 3: Verify API responds**

```bash
ssh huoshan 'curl -s http://localhost:5000/api/register -X POST -H "Content-Type: application/json" -d '"'"'{"username":"test","password":"test123"}'"'"''
```

- [ ] **Step 4: Check logs for errors**

```bash
ssh huoshan 'cd /root/spaced-repetition && tail -20 app.log'
```
