package remove_duplicates_from_sorted_list_ii

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var slow *ListNode
    fast := head
    var repeat bool

    for ; fast.Next != nil; fast = fast.Next {
        // when fast.Val is not equal to fast.Next.Val
        if fast.Val != fast.Next.Val {
            // if fast.Val is repeated, continue
            if repeat {
                repeat = false
                continue
            }
            // init variable
            if slow == nil {
                head, slow = fast, fast
            } else {
                // move slow pointer
                slow.Next = fast
                slow = slow.Next
            }
        } else {
            repeat = true
        }
    }

    if slow == nil {
        if !repeat {
            return fast
        }
        return nil
    } else {
        if !repeat {
            slow.Next = fast
        } else {
            slow.Next = nil
        }
    }

    return head
}
