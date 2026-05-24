# q
如何在 Ansible playbook 中使用 `import_tasks` 引入外部任务文件？
# a
使用 `import_tasks:` 指定文件路径，可以配合 `when` 条件控制是否导入该任务文件。示例：
```yaml
- name: upload rpm
  import_tasks: upload_tasks.yml
  when: upload_flag == "yes"
```

# q
在 `import_tasks` 中，`when` 条件应放在什么位置？
# a
`when` 必须与 `import_tasks` 处在同一缩进层级（同 level），这是控制该导入操作是否执行的方式。

# q
如何在任务执行时将操作委托给控制节点（本地执行）？
# a
使用 `delegate_to: localhost`，可以将任务（如 `command` 或 `shell`）交由 Ansible 控制节点执行，常用于文件拷贝等场景。示例：
```yaml
- name: copy mdbs rpm
  command: scp {{ mdbs_path }} {{ inventory_hostname }}:{{ mdbs_path }}
    warn=false
  delegate_to: localhost
```

# q
如何在 `when` 中表达多个条件必须同时满足（逻辑与）？
# a
在 `when` 下使用列表形式，每个元素作为一个条件，列表中的所有条件均为 true 时任务才会执行。示例：
```yaml
when:
  - inventory_hostname != groups.nodes[0]
  - env_path != ""
```

