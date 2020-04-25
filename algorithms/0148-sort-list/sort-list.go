package sort_list

type ListNode struct {
    Val  int
    Next *ListNode
}

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
