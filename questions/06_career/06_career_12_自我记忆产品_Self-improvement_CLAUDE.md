# q
这个项目是什么系统？
# a
这是一个基于艾宾浩斯遗忘曲线的间隔重复学习系统，通过 CLI 和 Web 界面帮助用户科学地复习知识。系统解析 Markdown 格式的问答文件，跟踪学习进度，并根据记忆保留算法安排复习计划。

# q
CLI 应用程序的核心命令有哪些？
# a
```bash
# 初始化知识库（解析配置目录中的 .md 文件）
python train.py --init

# 启动训练会话（默认命令）
python train.py

# 查看学习统计
python train.py --stats
```

# q
间隔重复算法中不同反馈等级对应的初始复习间隔是多少？
# a
- Level 1（熟练/Proficient）：168 小时（7 天）
- Level 2（一般/Fair）：72 小时（3 天）
- Level 3（忘记/Forgotten）：24 小时（1 天）
- Level 4（完全忘记/Completely Forgotten）：2 小时

# q
回答正确后，每个记忆等级的间隔乘数规则是什么？
# a
正确回答时，间隔乘数分别为：Level 1 为 2.5 倍，Level 2 为 1.8 倍，Level 3 为 1.3 倍，Level 4 为 1.0 倍。此外，当用户总正确率超过 80% 时，还会额外应用 1.2 倍的乘数。

# q
学习数据存储在哪个文件，其核心数据结构是怎样的？
# a
数据保存在 `data/learning_data.json` 中，结构为：
```json
{
  "questions": {
    "q_<hash>": {
      "question": "问题内容",
      "answer": "答案内容",
      "level": 1-4,
      "next_review": "ISO datetime",
      "review_count": 整数,
      "correct_count": 整数,
      "created_at": "ISO datetime",
      "last_reviewed": "ISO datetime"
    }
  }
}
```

