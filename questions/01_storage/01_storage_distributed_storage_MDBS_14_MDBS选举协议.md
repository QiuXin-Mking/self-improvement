# q
Paxos协议是什么？在分布式系统中解决什么问题？
# a
Paxos是一种基于消息传递且具有高度容错特性的一致性算法，用于在可能发生机器宕机、网络异常（延迟、丢失、重复、乱序、网络分区）的分布式系统中，快速且正确地在集群内部对某个数据的值达成一致，并保证不会破坏整个系统的一致性。

# q
Raft协议中节点有哪三种角色？各有什么作用？
# a
- Leader：处理客户端交互和日志复制操作，一般只有一个。
- Follower：普通节点，跟随Leader，同步数据。
- Candidate：候选者节点，仅在选举期间出现；当Follower对Leader心跳超时时，转变为Candidate并发起投票竞选Leader。获得过半数投票后成为新Leader。

# q
Raft协议主要包括哪三项任务？
# a
1. 竞选Leader
2. 日志复制
3. 安全

# q
Paxos协议的两个阶段（Prepare和Accept）的核心步骤是什么？
# a
第一阶段（Prepare）：
1. 提议者生成提案编号n。
2. 向所有节点广播prepare(n)请求。
3. 接收者比较n和minProposal，若n>minProposal则更新minProposal=n；如果已认可一个值，则返回(acceptedProposal, acceptedValue)，否则返回OK。
4. 提议者收到过半数响应后，若存在acceptedValue，保留编号最高的acceptedValue到本地。

第二阶段（Accept）：
5. 提议者广播accept(n, value)到所有节点。
6. 接收者若n>=minProposal，则接受该提案，持久化acceptedProposal=n, acceptedValue=value，并返回；否则拒绝并返回minProposal。
7. 提议者收到过半数响应后，若发现有返回值>n，表示有更新的提议，回到第一阶段重新发起；否则value达成一致。

