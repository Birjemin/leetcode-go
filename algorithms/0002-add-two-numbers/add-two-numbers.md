## 问题

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## 分析
使用链表来处理加减法

## 最高分
```golang
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry, dummy := 0, new(ListNode)
    for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
        if l1 != nil {
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            carry += l2.Val
            l2 = l2.Next
        }
        node.Next = &ListNode{carry%10, nil}
        carry /= 10
    }
    return dummy.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	each_sums := 0

	for l1 != nil || l2 != nil || each_sums != 0{
		if l1 != nil{
			each_sums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			each_sums += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: each_sums % 10}
		each_sums /= 10
		current = current.Next
	}

	return result.Next
}
```

## 实现
```golang
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
		} else {
			sum += 0
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		} else {
			sum += 0
		}

		// 余数
		z = (sum + a) % 10
		// 进位
		a = (sum + a) / 10

		temp.Val = z

		if l1 == nil && l2 == nil && a == 0 {
			break
		}
		temp.Next = &ListNode{Val: 0, Next: nil}
		temp = temp.Next
		sum = 0
	}
	return head
}
```

## 改进
参考别人的代码，简化逻辑
```golang

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
```

## 再次改进
看了最高分的，使用空间换取时间
```golang
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
	// 第一个节点不符合，所以返回head.Next
	return head.Next
}
```
## 反思
感觉最高分的很讨巧，没有想到会从第二个节点开始。。。

## 参考
1. [https://leetcode.com/problems/add-two-numbers/discuss/253942/Golang-12ms-100.00-5MB](https://leetcode.com/problems/add-two-numbers/discuss/253942/Golang-12ms-100.00-5MB)
2. [https://leetcode.com/problems/add-two-numbers/discuss/322184/golang-concise-solution](https://leetcode.com/problems/add-two-numbers/discuss/322184/golang-concise-solution)