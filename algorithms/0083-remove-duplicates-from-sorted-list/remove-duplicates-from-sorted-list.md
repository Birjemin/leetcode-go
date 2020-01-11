## 问题
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
```
Input: 1->1->2
Output: 1->2
```

Example 2:
```
Input: 1->1->2->3->3
Output: 1->2->3
```

## 分析
循环删除

## 最高分
这个很简洁，也很通俗易懂
```golang
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    current := head
    for current.Next != nil {
        if current.Val == current.Next.Val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    return head
}
```

## 实现
新建一个ListNode处理这个事情
```golang
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    res := &ListNode{}
    temp := res
    var first bool
    for ; head != nil; head = head.Next {
        if first == false {
            temp.Val, first = head.Val, true
        } else {
            if temp.Val != head.Val {
                temp.Next = &ListNode{Val: head.Val, Next: nil}
                temp = temp.Next
            }
        }
    }
    return res
}
```

## 改进
不新建ListNode处理，在原来的ListNode上面进行处理
```golang
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    temp, current := head, head
    for ;current != nil; current = current.Next {
        if temp.Val != current.Val {
            temp.Next, temp = current, current
        }
    }
    temp.Next = nil
    return head
}
```

## 反思

## 参考
