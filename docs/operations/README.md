# Operations SOP 索引

## 目录

```
docs/operations/
├── README.md               # 本文件 — SOP 索引与摘要
├── deploy-verify.md        # 部署验证 SOP（人 + AI）
├── deploy-verify.sh        # 部署验证一键脚本
├── e2e-demo-verify.md      # E2E 演示功能验证 SOP（AI 可执行）
└── verify-deploy.sh        # 部署验证自动化脚本
```

## SOP 摘要

### 1. 部署验证 SOP

| 字段 | 值 |
|------|-----|
| **文件** | `deploy-verify.md` |
| **执行者** | 人 + AI |
| **触发条件** | 每次部署到 huoshan 服务器后 |
| **预计耗时** | 3 分钟（脚本自动）/ 10 分钟（手动） |
| **自动化** | `bash verify-deploy.sh` |

**验证范围**：
- 服务存活（进程 + 日志）
- API 健康检查
- Demo 账户登录 + 统计
- 待复习题目获取
- 反馈提交流程
- 新用户注册
- 前端页面加载

**关键检查点**：
1. `HTTP 200` — 首页可达
2. `total_questions >= 8` — Demo 题目预置
3. `"success":true` — 反馈流程正常
4. 新用户 `total_questions == 0` — 数据隔离

---

### 2. E2E 演示功能验证 SOP

| 字段 | 值 |
|------|-----|
| **文件** | `e2e-demo-verify.md` |
| **执行者** | AI Agent |
| **触发条件** | 部署后 / 发布前 / 定期巡检 |
| **预计耗时** | 2 分钟（全自动） |
| **自动化** | AI 逐 Phase 执行 curl + 解析 |

**验证范围（6 个 Phase）**：
1. **环境就绪** — 服务可达、Demo 登录
2. **数据完整性** — 统计、待复习、分类
3. **核心学习流程** — 反馈提交、统计更新
4. **新用户隔离** — 注册、数据隔离、初始化
5. **前端资源** — HTML 结构、JS/CSS 可达
6. **错误处理** — 无效 Token、无效参数

**判定规则**：
- 🟢 PASS — 所有 MUST 通过
- 🟡 WARN — MUST 全过但 SHOULD 有失败
- 🔴 FAIL — 任一 MUST 失败

**AI 执行要点**：
- Phase 1 阻塞项失败则终止
- Phase 3 无待复习题目时自动跳过
- 最终输出判定矩阵汇总表
- 所有 curl 需 `--connect-timeout 5`

---

## 使用场景

| 场景 | 执行哪个 SOP | 谁执行 |
|------|-------------|--------|
| 刚部署完新版本 | 部署验证 SOP | 人/脚本 |
| 发布前最终检查 | E2E 验证 SOP | AI |
| 每日健康巡检 | E2E 验证 SOP | AI (cron) |
| 用户反馈 Bug | E2E 验证 SOP | AI |
| Demo 页面打不开 | 部署验证 SOP | 人 |
| 新功能上线后 | E2E 验证 SOP | AI |

## 快速命令

```bash
# 部署 + 验证一条龙
cd /Users/qiuxin/code/Self-improvement
bash scripts/deploy_to_huoshan.sh && bash docs/operations/verify-deploy.sh

# 仅验证（不部署）
bash docs/operations/verify-deploy.sh

# AI E2E 验证（复制 e2e-demo-verify.md 内容给 AI 执行）
```
