# q
单链表节点 `ListNode` 的结构定义是什么？请写出包含两个数据成员和三种构造函数的完整代码。
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
如何用迭代方法反转一个单链表？请写出函数实现。
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
如何使用快慢指针法找到单链表的中间节点？要求对于5个节点返回第3个节点，4个节点返回第2个节点。
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
// 示例：1->2->3->4->5 返回3；1->2->3->4 返回2
```

# q
如何合并两个有序链表？请写出使用虚拟头节点的实现。
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
在寻找并删除链表倒数第 N 个节点时，使用如下 `for` 循环移动指针到倒数第 N+1 个节点有何陷阱？应如何改写？
```cpp
for (int i = 0; i < n - 1; i++) {
    prev = curr;
    curr = curr->next;
}
```
# a
当 `n == 1` 时，循环条件 `i < 0` 为假，本应不执行，但由于写法问题或边界处理错误，可能错误执行一次。更安全的写法是使用 `while`：
```cpp
int i = 0;
while (i < n - 1) {
    prev = curr;
    curr = curr->next;
    i++;
}
```
或提前处理 `n == 1` 的边界情况。

