## 问题
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:
```
Input: 4->2->1->3
Output: 1->2->3->4
```

Example 2:
```
Input: -1->5->3->4->0
Output: -1->0->3->4->5
```

## 分析
对链表进行排序(不可以借用147的思想，插入排序的时间复杂度为n~n^2)
快速排序的复杂度满足要求（好好复习一下什么是快速排序）
堆排序（有趣的方式）、归并排序（需要借助21题的答案合并两个有序链表）

归并排序：先将数组不断细分成最小的单位，然后每个单位分别排序，排序完毕后合并，重复以上过程最后就可以得到排序结果。

快速排序：先选定一个基准元素，然后以该基准元素划分成两个数组，再在被划分的部分重复以上过程，最后可以得到排序结果。

此题为链表，适合归并排序，快速排序就会很复杂（转成数组就好了）

## 最高分
类似的方式
```golang
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l, r := head, head
	// deal with len of 2
	for r != nil && r.Next != nil && r.Next.Next != nil {
		l, r = l.Next, r.Next.Next
	}
	r = l.Next
	l.Next = nil
	l = head

	l, r = sortList(l), sortList(r)

	h := &ListNode{}
	t := h
	for l != nil && r != nil {
		if l.Val < r.Val {
			t.Next = l
			l = l.Next
		} else {
			t.Next = r
			r = r.Next
		}
		t = t.Next
	}
	if l != nil {
		t.Next = l
	}
	if r != nil {
		t.Next = r
	}
	return h.Next
}
```

## 实现
先使用归并排序分解，然后使用21题解答
```golang
func sortList(head *ListNode) *ListNode {

    if head == nil {
        return nil
    } else if head.Next == nil {
        return head
    }

    // split
    slow, fast := head, head.Next

    // find delimiter and recursive
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }
    fast, slow.Next = slow.Next, nil

    return mergeTwoLists(sortList(head), sortList(fast))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    if l1.Val > l2.Val {
        l2.Next = mergeTwoLists(l1, l2.Next)
        return l2
    } else {
        l1.Next = mergeTwoLists(l1.Next, l2)
        return l1
    }
}

```

## 改进
```golang

```

## 反思

## 参考