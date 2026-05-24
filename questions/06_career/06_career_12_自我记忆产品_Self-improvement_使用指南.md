# q
问题文件的标准格式是什么？
# a
在 `questions/` 目录下创建 `.md` 文件，每个问答对使用 `# q` 和 `# a` 独占一行，后接内容，例如：
```markdown

# q
你的问题内容
# a
你的答案内容
```

# q
如何配置系统扫描多个问题目录？
# a
编辑 `config.json`，在 `questions_dirs` 数组中添加目录路径：
```json
{
  "questions_dirs": [
    "questions",
    "my_questions",
    "path/to/other/questions"
  ]
}
```

# q
初始化知识库的命令是什么？
# a
```bash
python train.py --init
```
该命令会扫描配置的所有目录下的 `.md` 文件，提取问题和答案。

# q
训练中的四种反馈选项及其对应的复习间隔是什么？
# a
- `1` 或 `熟练` → 7天后复习
- `2` 或 `一般` → 3天后复习
- `3` 或 `忘记` → 1天后复习
- `4` 或 `完全忘记` → 2小时后复习

