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
```
* 时间复杂度：
* 空间复杂度：

## 改进
尝试缩减代码，全部放在一个循环中
```golang

```
* 时间复杂度：
* 空间复杂度：

## 反思

https://leetcode.com/problemset/algorithms/
https://leetcode.com/problems/add-two-numbers/
https://leetcode-cn.com/problems/add-two-numbers/
https://github.com/wind-liang/leetcode/blob/master/leetCode-2-Add-Two-Numbers.md
https://github.com/aQuaYi/LeetCode-in-Go/tree/master/Algorithms/0002.add-two-numbers

## 参考
1. []()