## 问题
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Example 1:
```
Input: 1->2->3->3->4->4->5
Output: 1->2->5
```

Example 2:
```
Input: 1->1->1->2->3
Output: 2->3
```

## 分析
删除重复的节点，和83题有区别，和48题处理方式类似
- 快慢指针

## 最高分
```golang
func deleteDuplicates(head *ListNode) *ListNode {
    // 长度 <=1 的 list ，可以直接返回
    if head == nil || head.Next == nil {
        return head
    }

    // 要么 head 重复了，那就删除 head
    if head.Val == head.Next.Val {
        for head.Next != nil && head.Val == head.Next.Val {
            head = head.Next
        }
        return deleteDuplicates(head.Next)
    }

    // 要么 head 不重复，递归处理 head 后面的节点
    head.Next = deleteDuplicates(head.Next)
    return head
}
```

## 实现
快慢指针法，一个遍历，一个标记不重复的位置
```golang
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var slow *ListNode
    fast := head
    var repeat bool

    for ; fast.Next != nil; fast = fast.Next {
        // when fast.Val is not equal to fast.Next.Val
        if fast.Val != fast.Next.Val {
            // if fast.Val is repeated, continue
            if repeat {
                repeat = false
                continue
            }
            // init variable
            if slow == nil {
                head, slow = fast, fast
            } else {
                // move slow pointer
                slow.Next = fast
                slow = slow.Next
            }
        } else {
            repeat = true
        }
    }

    if slow == nil {
        if !repeat {
            return fast
        }
        return nil
    } else {
        if !repeat {
            slow.Next = fast
        } else {
            slow.Next = nil
        }
    }

    return head
}
```

## 改进
```golang

```

## 反思

## 参考