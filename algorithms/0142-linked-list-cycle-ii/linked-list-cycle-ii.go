package linked_list_cycle_ii

type ListNode struct {
    Val  int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }
    slow := head
    for fast := slow.Next; fast != nil; fast, slow = fast.Next.Next, slow.Next {
        // find the circle
        if slow == fast {
            for fast = fast.Next; head != fast; fast, head = fast.Next, head.Next {
            }
            return head
        }
        if fast.Next == nil {
            break
        }
    }
    return nil
}

func detectCycle1(head *ListNode) *ListNode {

    if head == nil || head.Next == nil {
        return nil
    }

    fast := head.Next

    // find the cross point
    for slow := head; fast != slow; slow, fast = slow.Next, fast.Next.Next {
        // if it is not exists
        if fast == nil || fast.Next == nil {
            return nil
        }
    }

    for fast = fast.Next; head != fast; fast, head = fast.Next, head.Next {
    }

    return head
}
