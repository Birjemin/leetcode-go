## 问题
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.

## 分析

## 最高分
```golang

```

## 实现
```golang
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    h := head.Next
    h.Next, head.Next = head, swapPairs(h.Next)
    return h
}
```

## 改进
```golang

```

## 反思

## 参考
