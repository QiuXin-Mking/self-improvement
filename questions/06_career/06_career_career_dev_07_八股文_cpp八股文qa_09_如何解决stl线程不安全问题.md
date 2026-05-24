# q
如何解决STL vector在多线程并发push_back时的不安全问题？
# a
1. 加互斥锁（mutex）实现独占锁，适合写少读多且需要强一致性（满足CAP中的Consistent）的场景。  
2. 加共享锁（如std::shared_mutex），适合需要高可用（满足CAP中的Availability）的场景。

