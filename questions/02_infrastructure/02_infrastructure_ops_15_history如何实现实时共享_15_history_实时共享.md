# q
如何实现多个shell会话之间实时共享命令历史？
# a
在 `~/.bashrc` 中设置 `PROMPT_COMMAND` 环境变量，结合 `history -a` 和 `history -r` 实现。配置如下：
```bash
export PROMPT_COMMAND='history -a; history -r'
```
`history -a` 将当前会话的新命令追加到历史文件，`history -r` 立即从历史文件读取命令到当前会话，使多个会话能实时看到彼此的命令历史。同时建议设置足够大的 `HISTSIZE` 和 `HISTFILESIZE`，并使用 `shopt -s histappend` 确保多次退出时历史记录是追加而非覆盖。

# q
`PROMPT_COMMAND` 中的 `history -a` 和 `history -r` 分别有什么作用？
# a
- `history -a`：把当前 shell 会话中新产生的历史命令追加到历史文件（通常是 `~/.bash_history`）中，而不是等到退出时才写入。
- `history -r`：从历史文件中读取所有命令并加载到当前 shell 会话的内存历史列表中，使得其他会话刚写入的命令能立即在本会话中可见。

# q
`shopt -s histappend` 的作用是什么？
# a
启用 `histappend` 后，当多个 shell 会话退出时，它们各自的历史记录会以追加方式写入历史文件，而不是互相覆盖。这配合 `PROMPT_COMMAND` 的实时同步，可以更完整地保留所有会话的命令历史。

# q
如何使用 Ansible Playbook 批量配置 history 实时共享？
# a
可以编写一个包含 `lineinfile` 模块的 playbook，确保 `~/.bashrc` 中存在相关配置行。例如：
```yaml
---
- name: Update .bashrc for history management
  hosts: lustre_hosts
  become: yes
  tasks:
    - name: Ensure HISTSIZE is set
      lineinfile:
        path: ~/.bashrc
        regexp: '^export HISTSIZE='
        line: 'export HISTSIZE=10000'
        state: present
        create: yes

    - name: Ensure HISTFILESIZE is set
      lineinfile:
        path: ~/.bashrc
        regexp: '^export HISTFILESIZE='
        line: 'export HISTFILESIZE=20000'
        state: present
        create: yes

    - name: Ensure PROMPT_COMMAND is set
      lineinfile:
        path: ~/.bashrc
        regexp: '^export PROMPT_COMMAND='
        line: "export PROMPT_COMMAND='history -a; history -r'"
        state: present
        create: yes

    - name: Ensure histappend is set
      lineinfile:
        path: ~/.bashrc
        regexp: '^shopt -s histappend'
        line: 'shopt -s histappend'
        state: present
        create: yes
```
再配合 `ansible-playbook -i inventory update_bashrc.yml` 执行。

