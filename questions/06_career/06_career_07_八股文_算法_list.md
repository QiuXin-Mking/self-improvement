# q
在寻找链表中间节点的快慢指针实现中，`while` 循环条件写为 `fast->next != nullptr && fast->next->next != nullptr` 和写为 `fast != nullptr && fast->next != nullptr`，对返回的中间节点位置有什么不同？
# a
- 使用 `while (fast->next != nullptr && fast->next->next != nullptr)`：快指针最终只能停在链表倒数第二个节点（无法抵达 `nullptr`）。  
  - 若链表长度为奇数，快指针跳转次数为偶数，`slow` 停在正中间节点。  
  - 若链表长度为偶数，快指针跳转次数为奇数，`slow` 停在中间两个节点中偏左的那个。
- 使用 `while (fast != nullptr && fast->next != nullptr)`：快指针最终可以越过尾节点指向 `nullptr`。  
  - 若链表长度为奇数，`slow` 停在正中间节点。  
  - 若链表长度为偶数，`slow` 停在中间两个节点中偏右的那个。

```cpp
// 示例: 两种写法的差异
// 写法1
ListNode* middleNode1(ListNode* head) {
    ListNode* slow = head;
    ListNode* fast = head;
    while (fast->next != nullptr && fast->next->next != nullptr) {
        slow = slow->next;
        fast = fast->next->next;
    }
    return slow;
}

// 写法2
ListNode* middleNode2(ListNode* head) {
    ListNode* slow = head;
    ListNode* fast = head;
    while (fast != nullptr && fast->next != nullptr) {
        slow = slow->next;
        fast = fast->next->next;
    }
    return slow;
}
```

