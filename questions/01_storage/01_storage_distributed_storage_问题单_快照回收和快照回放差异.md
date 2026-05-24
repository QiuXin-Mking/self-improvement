# q
快照回收涉及哪些关键步骤？
# a
发送 mdlog，执行 free_block，最终真正执行删除。

# q
快照回放过程中涉及哪些关键函数？
# a
block_map_replay_delete、osd_replay_delete_log、mdlog_replay_log、osd_process_free_blk_handle。

# q
快照回收与快照回放的核心差异是什么？
# a
快照回收遵循“存在才能删除”的原则，通过发送 mdlog 触发 free_block 来真正删除；快照回放则是基于日志回放（replay）执行删除操作。

