## 问题
Remove all elements from a linked list of integers that have value val.

Example:
```
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
```

## 分析
删除链表中等于给定值 val 的所有节点。

送分题

## 最高分
```golang
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	curr := head

	for ;curr.Next!=nil; {
		if curr.Next.Val==val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	if head.Val==val {
		return head.Next
	}
	return head
}
```

## 实现
快慢指针
```golang
func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	for slow, fast := head, head.Next; fast != nil; fast = fast.Next {
		if fast.Val == val {
			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}
	}

	return head
}
```

## 改进
新增一个头节点，方便处理
```golang
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Next: head}

	for tmp := newHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return newHead.Next
}
```

## 反思

## 参考