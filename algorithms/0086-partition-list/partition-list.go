package partition_list

type ListNode struct {
    Val  int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    var slow, fast, after *ListNode

    for tmp := head; tmp != nil; tmp = tmp.Next {
        if tmp.Val < x {
            if slow == nil {
                slow, head = tmp, tmp
            } else {
                slow.Next = tmp
                slow = slow.Next
            }
        } else {
            if fast == nil {
                fast, after = tmp, tmp
            } else {
                fast.Next = tmp
                fast = fast.Next
            }
        }
    }

    if slow == nil {
        return after
    }

    if fast == nil {
        slow.Next = nil
        return head
    }

    slow.Next, fast.Next = after, nil
    return head
}

func partition1(head *ListNode, x int) *ListNode {
    slow, fast := &ListNode{}, &ListNode{}
    ptr1, ptr2 := slow, fast
    for ; head != nil; head = head.Next {
        if head.Val < x {
            ptr1.Next = head
            ptr1 = ptr1.Next
        } else {
            ptr2.Next = head
            ptr2 = ptr2.Next
        }
    }
    ptr1.Next, ptr2.Next = fast.Next, nil
    return slow.Next
}
