## 问题
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?

## 分析
反转一个单链表。使用迭代或递归地反转链表。

提示已经给了解决方案

92题的简化

## 最高分
```golang
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		cur, pre, cur.Next = cur.Next, cur, pre
	}
	return pre
}
```

## 实现
迭代(等价于最高分解答)
```golang
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head.Next
	head.Next = nil

	var t *ListNode
	for fast != nil {
		t, fast = fast, fast.Next
		t.Next, head = head, t
	}

	return head
}
```

## 改进
递归
```golang
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head.Next
	head.Next = nil

	head, _ = cal(head, fast)

	return head
}

func cal(h1, h2 *ListNode) (*ListNode, *ListNode) {
	if h2 == nil {
		return h1, nil
	}
	t := h2.Next
	h2.Next = h1
	return cal(h2, t)
}

```

## 反思

## 参考