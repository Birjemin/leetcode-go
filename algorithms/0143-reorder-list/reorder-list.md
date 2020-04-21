## 问题
Given a singly linked list L: `L0→L1→…→Ln-1→Ln`,
reorder it to: `L0→Ln→L1→Ln-1→L2→Ln-2→…`

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:
```
Given 1->2->3->4, reorder it to 1->4->2->3.
```

Example 2:
```
Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
```

## 分析
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

最简单的方式分成两部分，后半部分翻转，然后再拼接

## 最高分
巧妙的使用了快慢指针~~
```golang
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }
    
    var prev *ListNode
    for slow != nil {
        slow.Next, prev, slow = prev, slow, slow.Next
    }

    first := head
    for prev.Next != nil{
        first.Next, first = prev, first.Next
        prev.Next, prev = first, prev.Next
    }
}
```

## 实现
分成两部分，后半部分翻转，然后再逐次拼接
```golang
func reorderList(head *ListNode) {
    tmp := head
    var count int
    // calculate the number of listNode
    for tmp != nil {
        count++
        tmp = tmp.Next
    }

    var newHead *ListNode
    point := (count + 1) / 2

    // the part which should be reversed of listNode
    for tmp, count = head, 0; tmp != nil; tmp = tmp.Next {
        count++
        if point == count {
            newHead, tmp.Next = tmp.Next, nil
        }
    }

    // reverse
    newHead = reverseNode(newHead)

    // link two listNode
    tmp1 := head
    var tmp2 *ListNode
    for newHead != nil {
        tmp, tmp1.Next = tmp1.Next, newHead
        tmp2, newHead.Next = newHead.Next, tmp
        tmp1, newHead = tmp, tmp2
    }
}

func reverseNode(head *ListNode) *ListNode {
    var t, nHead *ListNode
    for ; head != nil; head = t {
        t = head.Next
        head.Next, nHead = nHead, head
    }
    return nHead
}

```

## 改进
```golang

```

## 反思

## 参考