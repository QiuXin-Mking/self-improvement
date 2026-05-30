#!/bin/bash
# KnowLoop 部署验证脚本
# 用法: bash docs/operations/verify-deploy.sh [BASE_URL]
# 默认 BASE_URL=http://115.190.235.149:4430

set -e

BASE="${1:-http://115.190.235.149:4430}"
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
        echo "     got: $(echo "$actual" | head -c 200)"
        FAIL=$((FAIL+1))
    fi
}

echo "========================================="
echo "KnowLoop 部署验证"
echo "目标: $BASE"
echo "时间: $(date '+%Y-%m-%d %H:%M:%S')"
echo "========================================="
echo ""

# 1. Service health
echo "1. 服务健康检查"
HTTP=$(curl -s -o /dev/null -w "%{http_code}" --connect-timeout 5 "$BASE/")
check "首页返回 200" "200" "$HTTP"

# 2. Demo login
echo "2. Demo 账户登录"
LOGIN=$(curl -s --connect-timeout 5 -X POST "$BASE/api/login" \
    -H 'Content-Type: application/json' \
    -d '{"username":"demo","password":"demo123"}')
check "Demo 登录成功" '"success":true' "$LOGIN"

TOKEN=$(echo "$LOGIN" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])" 2>/dev/null || echo "")

if [ -n "$TOKEN" ]; then
    # 3. Stats
    echo "3. 统计数据"
    STATS=$(curl -s --connect-timeout 5 "$BASE/api/stats" -H "Authorization: Bearer $TOKEN")
    # 检查总题目数（首次部署=8，后续可能更多）
    TOTAL=$(echo "$STATS" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['stats']['total_questions'])" 2>/dev/null || echo "0")
    if [ "$TOTAL" -ge 8 ]; then
        echo "  ✅ 总题目数 $TOTAL (≥8)"
        PASS=$((PASS+1))
    else
        echo "  ❌ 总题目数 $TOTAL (expected ≥8)"
        FAIL=$((FAIL+1))
    fi
    # 显示完整统计
    echo "     $(echo "$STATS" | python3 -c "import sys,json; s=json.load(sys.stdin)['data']['stats']; print(f'总题{s[\"total_questions\"]} | 待复习{s[\"due_questions\"]} | 复习{s[\"total_reviews\"]}次 | 正确率{s[\"accuracy\"]}%')" 2>/dev/null)"
    TOTAL_REVIEWS=$(echo "$STATS" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['stats']['total_reviews'])" 2>/dev/null || echo "0")

    # 4. Due questions
    echo "4. 待复习题目"
    DUE=$(curl -s --connect-timeout 5 "$BASE/api/due-questions" -H "Authorization: Bearer $TOKEN")
    DUE_TOTAL=$(echo "$DUE" | python3 -c "import sys,json; print(json.load(sys.stdin)['data'].get('total',0))" 2>/dev/null || echo "0")
    if [ "$DUE_TOTAL" -gt 0 ]; then
        echo "  ✅ 待复习题目: $DUE_TOTAL 道"
        PASS=$((PASS+1))
    else
        echo "  ⚠️  待复习: 0（所有题目已复习完毕或首次部署未完成）"
    fi

    # 5. Feedback
    echo "5. 反馈流程"
    QID=$(echo "$DUE" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['questions'][0]['id'])" 2>/dev/null || echo "")
    if [ -n "$QID" ]; then
        FB=$(curl -s --connect-timeout 5 -X POST "$BASE/api/update-review" \
            -H 'Content-Type: application/json' \
            -H "Authorization: Bearer $TOKEN" \
            -d "{\"question_id\":\"$QID\",\"feedback\":1}")
        check "反馈提交成功" '"success":true' "$FB"

        STATS2=$(curl -s --connect-timeout 5 "$BASE/api/stats" -H "Authorization: Bearer $TOKEN")
        REVIEWS2=$(echo "$STATS2" | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['stats']['total_reviews'])" 2>/dev/null || echo "0")
        if [ "$REVIEWS2" -gt "$TOTAL_REVIEWS" ]; then
            echo "  ✅ 复习次数递增: $TOTAL_REVIEWS → $REVIEWS2"
            PASS=$((PASS+1))
        else
            echo "  ❌ 复习次数未递增 (before=$TOTAL_REVIEWS after=$REVIEWS2)"
            FAIL=$((FAIL+1))
        fi
    else
        echo "  ⚠️  无待复习题目，跳过反馈测试（非首次部署正常现象）"
    fi
else
    echo "  ⚠️  无法获取 Token，跳过后续测试"
fi

# 7. New registration
echo "6. 新用户注册"
TEST_USER="verify_$(date +%s)"
REG=$(curl -s --connect-timeout 5 -X POST "$BASE/api/register" \
    -H 'Content-Type: application/json' \
    -d "{\"username\":\"$TEST_USER\",\"password\":\"test1234\"}")
check "新用户注册成功" '"success":true' "$REG"

echo ""
echo "========================================="
if [ "$FAIL" -eq 0 ]; then
    echo "结果: ✅ 全部 $PASS 项通过"
    echo "========================================="
    exit 0
else
    echo "结果: ✅ $PASS 通过  ❌ $FAIL 失败"
    echo "========================================="
    exit 1
fi
