# q
Alembic 是什么，它的主要作用是什么
# a
Alembic 是 SQLAlchemy 的数据库迁移工具，用于版本化管理数据库模式变更（如添加表、列、修改结构），确保不同环境（开发、测试、生产）之间数据库一致性。

# q
如何初始化一个 Alembic 迁移环境
# a
使用命令 `alembic init alembic`，这会在项目根目录创建 `alembic` 目录，包含配置文件和迁移脚本模板。

# q
Alembic 如何自动生成迁移脚本
# a
使用 `alembic revision --autogenerate -m "描述信息"` 命令，Alembic 会自动检测 SQLAlchemy 模型与当前数据库之间的差异，并生成相应的升级 SQL 语句。

# q
如何将数据库升级到最新版本
# a
运行 `alembic upgrade head`，它会执行所有未应用的迁移脚本，将数据库 schema 更新到最新状态。

# q
Alembic 如何回滚到上一个版本
# a
使用 `alembic downgrade -1` 回滚一步，或通过 `alembic downgrade <revision_id>` 回滚到指定版本。

