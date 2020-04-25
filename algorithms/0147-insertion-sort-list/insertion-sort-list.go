package insertion_sort_list

type ListNode struct {
    Val  int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    tmp := head.Next
    for head.Next = nil; tmp != nil; {
        t := tmp
        tmp, t.Next = tmp.Next, nil
        head = insertNode(head, t)
    }
    return head
}

// insert one node
func insertNode(start, target *ListNode) *ListNode {
    // first
    var front *ListNode
    head := start
    for start != nil {
        if start.Val > target.Val {
            target.Next = start
            break
        }
        front, start = start, start.Next
    }
    if front == nil {
        return target
    }
    front.Next = target
    return head
}

func insertionSortList1(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var tmp *ListNode
    tummy := &ListNode{}
    for pre := tummy; head != nil; head = tmp {

        // store the the next node of head(head=tmp)
        // reset the first node
        pre, tmp = tummy, head.Next

        // find the position should be exchanged
        for pre.Next != nil && pre.Next.Val < head.Val {
            pre = pre.Next
        }

        // insert the node
        head.Next, pre.Next = pre.Next, head
    }
    return tummy.Next
}
