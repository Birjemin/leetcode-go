## 问题
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:
```
Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
```

## 分析
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。

## 最高分
两个指针，然后移动
```golang
func partition(head *ListNode, x int) *ListNode {
    first, second := &ListNode{}, &ListNode{}
    dummyHeadFirst, dummyHeadSecond := first, second

    for head != nil {
        if head.Val < x {
            first.Next = head
            first = first.Next
        } else {
            second.Next = head
            second = second.Next
        }
        head = head.Next
    }
    first.Next, second.Next = nil, nil // cut the connection
    if dummyHeadSecond.Next != nil {
        first.Next = dummyHeadSecond.Next
    }
    return dummyHeadFirst.Next
}
```

## 实现
两个指针，分别指向
```golang
func partition(head *ListNode, x int) *ListNode {
    var slow, fast, after *ListNode

    for tmp := head; tmp != nil; tmp = tmp.Next {
        if tmp.Val < x {
            if slow == nil {
                slow, head = tmp, tmp
            } else {
                slow.Next = tmp
                slow = slow.Next
            }
        } else {
            if fast == nil {
                fast, after = tmp, tmp
            } else {
                fast.Next = tmp
                fast = fast.Next
            }
        }
    }

    if slow == nil {
        return after
    }

    if fast == nil {
        slow.Next = nil
        return head
    }

    slow.Next, fast.Next = after, nil
    return head
}
```

## 改进
头节点设置为空节点，这样不需要判断是否为nil
```golang
func partition1(head *ListNode, x int) *ListNode {
    slow, fast := &ListNode{}, &ListNode{}
    ptr1, ptr2 := slow, fast
    for ; head != nil; head = head.Next {
        if head.Val < x {
            ptr1.Next = head
            ptr1 = ptr1.Next
        } else {
            ptr2.Next = head
            ptr2 = ptr2.Next
        }
    }
    ptr1.Next, ptr2.Next = fast.Next, nil
    return slow.Next
}
```

## 反思

## 参考