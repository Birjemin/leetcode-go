package add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int       // 值
	Next *ListNode // 下一位地址
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{}
	temp := head
	var a int
	for l1 != nil || l2 != nil || a != 0 {
		if l1 != nil {
			a += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			a += l2.Val
			l2 = l2.Next
		}

		temp.Val = a % 10
		a /= 10
		if l1 == nil && l2 == nil && a == 0 {
			break
		}

		temp.Next = &ListNode{Val: 0, Next: nil}
		temp = temp.Next
	}
	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{}
	temp := head
	var a int
	for l1 != nil || l2 != nil || a != 0 {
		if l1 != nil {
			a += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			a += l2.Val
			l2 = l2.Next
		}

		temp.Next = &ListNode{Val: a % 10, Next: nil}
		a /= 10
		temp = temp.Next
	}
	// 第一个不符合
	return head.Next
}