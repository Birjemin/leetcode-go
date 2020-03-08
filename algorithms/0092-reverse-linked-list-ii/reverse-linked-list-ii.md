## 问题
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:
```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
```

## 分析
对m~n之间的数组进行翻转，然后连接，注意一下m == 1, n == count(list)的情况

## 最高分
```golang
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || m >= n {
        return head
    }
    newHead := &ListNode{Val: 0, Next: head}
    pre := newHead
    for count := 0; pre.Next != nil && count < m-1; count++ {
        pre = pre.Next
    }
    if pre.Next == nil {
        return head
    }
    cur := pre.Next
    for i := 0; i < n-m; i++ {
        tmp := pre.Next
        pre.Next = cur.Next
        cur.Next = cur.Next.Next
        pre.Next.Next = tmp
    }
    return newHead.Next
}
```

## 实现
```golang
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    var nHead, nEnd, lEnd, t2 *ListNode
    tmp := head
    for i := 1; i <= n; i++ {
        if i <= m-1 {
            if i == m-1 {
                // get the pointer for left end: lEnd
                lEnd = tmp
            }
            tmp = tmp.Next
            continue
        }
        t2 = tmp.Next
        // if new list's end is nil
        if i == m {
            // get the pointer for new list's end: nEnd
            tmp.Next, nEnd = nil, tmp
        } else {
            // link the list
            tmp.Next = nHead
        }
        // update the new list's Head: nHead
        nHead, tmp = tmp, t2
    }
    if nEnd != nil {
        nEnd.Next = tmp
    }
    if lEnd == nil {
        return nHead
    }
    lEnd.Next = nHead
    return head
}
```

## 改进
```golang

```

## 反思

## 参考