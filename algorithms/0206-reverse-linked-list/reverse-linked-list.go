package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		cur, pre, cur.Next = cur.Next, cur, pre
	}
	return pre
}

func reverseList2(head *ListNode) *ListNode {
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

func reverseList1(head *ListNode) *ListNode {
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
