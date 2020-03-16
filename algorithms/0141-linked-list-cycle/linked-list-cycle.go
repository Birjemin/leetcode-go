package linked_list_cycle

type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    for fast := head.Next; fast != nil; {
        if head == fast {
            return true
        }
        head = head.Next
        if fast.Next == nil {
            return false
        }
        fast = fast.Next.Next
    }
    return false
}
