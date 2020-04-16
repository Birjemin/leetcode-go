## 问题
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

 

Example 1:
```
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```
![img](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)

Example 2:
```
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```
![img](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)

Example 3:
```
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```
![img](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)

## 分析
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

简单，和141类似，先找到相交点，然后一个从head头节点重新跑，一个从相交点跑，相交的时候就是相交点

## 最高分
```golang
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}

	if slow != fast {
		return nil
	}

	for slow != head {
		slow, head = slow.Next, head.Next
	}
	return slow
}
```

## 实现
```golang
func detectCycle(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }
    slow := head
    for fast := slow.Next; fast != nil; fast, slow = fast.Next.Next, slow.Next {
        // find the circle
        if slow == fast {
            for fast = fast.Next; head != fast; fast, head = fast.Next, head.Next {
            }
            return head
        }
        if fast.Next == nil {
            break
        }
    }
    return nil
}
```

## 改进
```golang
func detectCycle(head *ListNode) *ListNode {

    if head == nil || head.Next == nil {
        return nil
    }

    fast := head.Next

    // find the cross point
    for slow := head; fast != slow; slow, fast = slow.Next, fast.Next.Next {
        // if it is not exists
        if fast == nil || fast.Next == nil {
            return nil
        }
    }

    for fast = fast.Next; head != fast; fast, head = fast.Next, head.Next {
    }

    return head
}
```

## 反思

## 参考