#!/bin/bash

# 清除复习进度数据脚本

DATA_FILE="data/learning_data.json"
BACKUP_FILE="data/learning_data.json.bak"
DB_FILE="data/app.db"
DB_BACKUP_FILE="data/app.db.bak"
TOTAL_COUNT=0

# 清除 JSON 数据
clear_json_data() {
    if [ ! -f "$DATA_FILE" ]; then
        return
    fi

    # 备份原文件
    cp "$DATA_FILE" "$BACKUP_FILE"
    echo "✓ 已备份 JSON 文件到: $BACKUP_FILE"

    # 重置所有问题的学习进度
    jq '{
        questions: (.questions | with_entries(.value |= {
            id: .id,
            question: .question,
            answer: .answer,
            source: .source,
            level: 1,
            next_review: now | todateiso8601,
            review_count: 0,
            correct_count: 0,
            created_at: .created_at,
            last_reviewed: ""
        })),
        last_updated: now | todateiso8601
    }' "$DATA_FILE" > "${DATA_FILE}.tmp"

    # 检查 jq 是否成功
    if [ $? -eq 0 ]; then
        mv "${DATA_FILE}.tmp" "$DATA_FILE"
        QUESTION_COUNT=$(jq '.questions | length' "$DATA_FILE")
        echo "✓ 成功清除 JSON 数据中 $QUESTION_COUNT 个问题的复习进度"
        TOTAL_COUNT=$((TOTAL_COUNT + QUESTION_COUNT))
    else
        rm -f "${DATA_FILE}.tmp"
        echo "✗ 清除 JSON 数据失败，请检查 jq 是否已安装"
    fi
}

# 清除 SQLite 数据
clear_sqlite_data() {
    if [ ! -f "$DB_FILE" ]; then
        return
    fi

    # 备份原文件
    cp "$DB_FILE" "$DB_BACKUP_FILE"
    echo "✓ 已备份 SQLite 文件到: $DB_BACKUP_FILE"

    # 重置所有问题的学习进度
    sqlite3 "$DB_FILE" << EOF
UPDATE questions
SET level = 1,
    next_review = datetime('now'),
    review_count = 0,
    correct_count = 0,
    last_reviewed = NULL,
    updated_at = datetime('now');
EOF

    # 检查 sqlite3 是否成功
    if [ $? -eq 0 ]; then
        # 获取更新的行数
        QUESTION_COUNT=$(sqlite3 "$DB_FILE" "SELECT COUNT(*) FROM questions;")
        echo "✓ 成功清除 SQLite 数据库中 $QUESTION_COUNT 个问题的复习进度"
        TOTAL_COUNT=$((TOTAL_COUNT + QUESTION_COUNT))
    else
        echo "✗ 清除 SQLite 数据失败，请检查 sqlite3 是否已安装"
    fi
}

# 执行清除
clear_json_data
clear_sqlite_data

# 显示总计
if [ $TOTAL_COUNT -eq 0 ]; then
    echo "没有找到需要清除的数据"
else
    echo ""
    echo "总计: 成功清除 $TOTAL_COUNT 个问题的复习进度数据"
fi
