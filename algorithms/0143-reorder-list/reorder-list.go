package reorder_list

type ListNode struct {
    Val  int
    Next *ListNode
}

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
        t, head.Next, nHead = head.Next, nHead, head
    }
    return nHead
}
