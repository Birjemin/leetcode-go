## 问题
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Example 1:
```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

Example 2:
```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

Example 3:
```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```

Follow up:

Can you solve it using O(1) (i.e. constant) memory?

## 分析
判断链表是否存在环形
- 两个指针，一个一步一步跑，一个2倍跑，总会相遇则存在，反之不存在

## 最高分
快慢指针
```golang
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    slow := head
    fast := head.Next

    for {
        if fast == nil || fast.Next == nil {
            break
        }
        if slow == fast {
            return true
        }

        fast = fast.Next.Next
        slow = slow.Next

    }

    return false

}
```


## 实现
快慢指针，总会相遇
```golang
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    for fast := head.Next; fast != nil; fast = fast.Next.Next {
        if head == fast {
            return true
        }
        head = head.Next
        if fast.Next == nil {
            return false
        }
    }
    return false
}
```

## 改进
```golang

```

## 反思

## 参考