# q
cur 和 prev 的完整拼写分别是什么？
# a
cur 是 current，prev 是 previous。

# q
如何用 Python 实现反转链表？
# a
```python
class solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        cur = head
        while cur:
            nxt = cur.next
            cur.next = prev
            prev = cur
            cur = nxt
        return prev
```

# q
反转链表完成后，返回的节点是什么？循环结束条件是什么？
# a
返回的节点是 `prev`。结束条件是 `cur == None`，此时 `prev` 指向原链表的最后一个节点，即反转后的新头节点。

# q
如何用 Python 实现反转链表 II（指定区间反转）？
# a
```python
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        p0 = dummy = ListNode(next=head)
        for _ in range(left - 1):
            p0 = p0.next
        
        pre = None
        cur = p0.next
        for _ in range(right - left + 1):
            nxt = cur.next
            cur.next = pre
            pre = cur
            cur = nxt
        
        p0.next.next = cur
        p0.next = pre
        return dummy.next
```

