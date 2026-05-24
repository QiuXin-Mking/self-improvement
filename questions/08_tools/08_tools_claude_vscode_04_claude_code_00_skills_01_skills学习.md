# q
Claude Code 的 Skills 可以存放在哪两种路径下？
# a
- 项目级别：`<项目名称>/.claude/skills/`
- 全局级别：`<全局路径>/.claude/skills/`

# q
Skills 目录应包含哪些基本组成部分？
# a
```
1. 编写 SKILL.md （必须大写）
2. scripts\
3. templates\
```

# q
SKILL.md 文件必须包含哪两个元数据字段？它们各自的作用是什么？
# a
- `name`：技能的名称
- `description`：描述信息，用于匹配用户的指令，触发该技能

# q
在 SKILL.md 的指令部分，可以包含哪些典型的约束或自动化步骤？
# a
- 约束示例：禁止向用户询问意图、智能适配
- 自动化步骤：定义执行步骤、指定所需资源、引用代码等

