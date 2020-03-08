package reverse_linked_list_ii

type ListNode struct {
    Val  int
    Next *ListNode
}

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
