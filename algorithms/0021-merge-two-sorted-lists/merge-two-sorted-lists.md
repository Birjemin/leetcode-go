## 问题
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

```
Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
```

## 分析


## 最高分
```golang
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    }
    l2.Next = mergeTwoLists(l1, l2.Next)
    return l2
}
```


## 实现
```golang
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    var head, temp *ListNode
    if l1.Val > l2.Val {
        head = l2
        temp = l2
        l2 = l2.Next
    } else {
        head = l1
        temp = l1
        l1 = l1.Next
    }
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
           temp.Next = l2
           l2 = l2.Next
        } else {
           temp.Next = l1
           l1 = l1.Next
        }
        temp = temp.Next
    }

    if l1 == nil {
        temp.Next = l2
    } else {
        temp.Next = l1
    }
    return head
}
```

## 改进
改成递归~~~
```golang
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    var head *ListNode
    if l1.Val > l2.Val {
        head = l2
        head.Next = mergeTwoLists(l2.Next, l1)
    } else {
        head = l1
        head.Next = mergeTwoLists(l1.Next, l2)
    }
    return head
}
```

## 反思
递归可以降维，便于管理

## 参考