## 问题
Given a singly linked list, determine if it is a palindrome.

Example 1:
```
Input: 1->2
Output: false
```

Example 2:
```
Input: 1->2->2->1
Output: true
```

Follow up:
- Could you do it in O(n) time and O(1) space?

## 分析
请判断一个链表是否为回文链表。最好使用O(1)空间复杂度。
206和92题目的升级，先分段a,b，翻转a，判断两者是否相等即可，143题的方式就可以解答

## 最高分
```golang
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    
    var (
        slow *ListNode = head
        fast *ListNode = head
    )
    
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next;
        slow = slow.Next;
    }
    
    slow.Next = reverse(slow.Next)
    slow = slow.Next
    
    for slow != nil {
        if slow.Val != head.Val {
            return false
        }
        slow = slow.Next
        head = head.Next
    }
    
    return true
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        temp := head.Next
        head.Next = prev
        prev = head
        head = temp
    }
    return prev
}
```

## 实现
先split两份，然后翻转一份，比较是否相等
```golang
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// split
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
```

## 改进
```golang

```

## 反思

## 参考