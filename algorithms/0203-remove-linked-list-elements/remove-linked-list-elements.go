package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func removeElements1(head *ListNode, val int) *ListNode {
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
