## 问题
Given a linked list, remove the n-th node from the end of list and return its head.

Example:
```
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
```
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?

## 分析
删除链表的倒数第N个节点。
- 可能为head结点
- 一次移动求解

## 最高分
设置 2 个指针，一个指针距离前一个指针n个距离。同时移动这2个指针，2个指针都移动相同的距离。当一个指针移动到了终点，那么前一个指针就是倒数第n个节点了
```golang
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    h := head
	for n > 0 {
		n--
		if h != nil {
			h = h.Next
		} else {
			return head
		}
	}
	if h == nil {
		return head.Next
	}
	s := head
	for h.Next != nil {
		h = h.Next
		s = s.Next
	}
	if s.Next != nil {
		s.Next = s.Next.Next
	}
	return head
}
```

## 实现
先得到链表长度，然后计算即可
```golang
type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    temp := head
    var length int
    // count ListNode length
    for temp != nil {
        length++
        temp = temp.Next
    }
    // just one node
    if length == 1 {
        return nil
    } else if n == length {
        // head node
        return head.Next
    } else {
        n, length = length-n, 0
        for temp = head; temp != nil; temp = temp.Next {
            length++
            if n == length {
                if temp.Next != nil {
                    temp.Next = temp.Next.Next
                }
                break
            }
        }
        return head
    }
}
```

## 改进
参考最高分解答一下
```golang
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    temp := head
    // get temp position
    for i := 0; i < n; i++ {
        temp = temp.Next
    }
    // n == length of ListNode
    if temp == nil {
        return head.Next
    }
    // head node
    prev := head
    // both of prev and temp should attempt to move
    for temp.Next != nil {
        temp, prev = temp.Next, prev.Next
    }
    // if not equal nil, connection the next node
    if prev.Next != nil {
        prev.Next = prev.Next.Next
    }
    return head
}

```

## 反思

## 参考
