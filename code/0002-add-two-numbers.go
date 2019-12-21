package code

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

	head := &ListNode{Val: 0, Next: nil}
	temp := head
	var x, y, z, a int
	for l1 != nil || l2 != nil || a != 0 {

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}

		// 余数
		z = (x + y + a) % 10
		// 进位
		a = (x + y + a) / 10

		temp.Val = z

		if l1 == nil && l2 == nil && a == 0 {
			break
		}
		temp.Next = &ListNode{Val: 0, Next: nil}
		temp = temp.Next
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{Val: 0, Next: nil}
	temp := head
	var sum, z, a int
	for l1 != nil || l2 != nil || a != 0 {

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// 余数
		z = (sum + a) % 10
		// 进位
		a = (sum + a) / 10

		temp.Val = z

		println("sum, z, a: ", sum, z, a)
		if l1 == nil && l2 == nil && a == 0 {
			break
		}
		temp.Next = &ListNode{Val: 0, Next: nil}
		temp = temp.Next
		sum = 0
	}
	return head
}

func makeListNode(params []int) *ListNode {
	if len(params) == 0 {
		return nil
	}
	res := &ListNode{
		Val: params[0],
	}
	temp := res
	for i, v := range params {
		if i != 0 {
			temp.Next = &ListNode{Val: v}
			temp = temp.Next
		} else {
			continue
		}
	}
	return res
}
