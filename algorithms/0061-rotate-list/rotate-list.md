## 问题
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:
```
Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
```

Example 2:
```
Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
```

## 分析
旋转链表，将每个节点向后移动k个节点
- 先得到节点数目count
- 连接收尾，移动未知
- k需要和count取余，减少移动次数

## 最高分
// 没有对k进行处理。。。
```golang
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
        return head
    }

    slow, fast := head, head
    for i := 0; i < k; i++ {
        if fast.Next == nil {
            return rotateRight(head, k%(i+1))
        }
        fast = fast.Next
    }

    for fast.Next != nil {
        slow, fast = slow.Next, fast.Next
    }
    newHead := slow.Next
    slow.Next, fast.Next = nil, head
    return newHead
}
```

## 实现
```golang
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
        return head
    }

    slow, fast := head, head
    for i := 0; i < k; i++ {
        if fast.Next == nil {
            return rotateRight(head, k%(i+1))
        }
        fast = fast.Next
    }

    for fast.Next != nil {
        slow, fast = slow.Next, fast.Next
    }
    newHead := slow.Next
    slow.Next, fast.Next = nil, head
    return newHead
}
```

## 改进
移除多余的变量
```golang
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }
    end, count := head, 1
    // find the end of ListNode and the number of ListNode
    for ; end.Next != nil; count++ {
        end = end.Next
    }
    if count == 1 {
        return head
    }
    if k %= count; k == 0 {
        return head
    }
    // connect the between of head node and end node，and move the end pointer to the first node
    end.Next, end = head, head
    // find the position of k
    for k = count - k; k > 1; k-- {
        end = end.Next
    }
    // replace the pointer
    head, end.Next = end.Next, nil
    return head
}
```

## 反思

## 参考