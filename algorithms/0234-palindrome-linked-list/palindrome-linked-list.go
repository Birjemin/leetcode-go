package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// split list
	slow, fast := head, head.Next
	for {
		if fast.Next == nil {
			fast, slow.Next = slow.Next, nil
			break
		} else if fast.Next.Next == nil {
			fast, slow.Next = slow.Next.Next, nil
			break
		}
		fast, slow = fast.Next.Next, slow.Next
	}

	// reverse slow
	slow = nil
	for head != nil {
		head, slow, head.Next = head.Next, head, slow
	}

	// check
	for fast != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast, slow = fast.Next, slow.Next
	}
	return true
}
