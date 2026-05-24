# q
单链表节点 `ListNode` 的 C++ 结构体定义包含哪两个成员和哪三种构造函数？
# a
```cpp
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};
```

# q
如何用快慢指针找到单链表的中间节点？（要求：5 个节点返回第 3 个，4 个节点返回第 2 个）
# a
```cpp
ListNode* middleNode(ListNode* head) {
    ListNode* slow = head;
    ListNode* fast = head;
    while (fast->next != nullptr && fast->next->next != nullptr) {
        slow = slow->next;
        fast = fast->next->next;
    }
    return slow;
}
// 对于 1 2 3 4 5，返回节点 3
// 对于 1 2 3 4，返回节点 2
```

# q
迭代法反转单链表的完整代码是什么？
# a
```cpp
ListNode* reverseList(ListNode* head) {
    ListNode* prev = nullptr;
    ListNode* curr = head;
    while (curr != nullptr) {
        ListNode* next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
}
```

# q
合并两个有序链表的标准实现是怎样的？
# a
```cpp
ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
    ListNode dummy(0);
    ListNode* tail = &dummy;

    while (list1 && list2) {
        if (list1->val <= list2->val) {
            tail->next = list1;
            list1 = list1->next;
        } else {
            tail->next = list2;
            list2 = list2->next;
        }
        tail = tail->next;
    }

    tail->next = (list1 != nullptr) ? list1 : list2;
    return dummy.next;
}
```

# q
在遍历链表时，`for (int i = 0; i < n - 1; i++)` 有什么陷阱？
# a
当 `n == 1` 时，这个 `for` 循环可能错误地执行一次。推荐使用 `while` 循环配合条件判断来避免该问题：
```cpp
int i = 0;
while (i < n - 1) {
    prev = curr;
    curr = curr->next;
    i++;
}
```

