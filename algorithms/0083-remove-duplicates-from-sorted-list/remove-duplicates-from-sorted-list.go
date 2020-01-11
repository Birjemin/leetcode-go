package remove_duplicates_from_sorted_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates1(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    res := &ListNode{}
    temp := res
    var first bool
    for ; head != nil; head = head.Next {
        if first == false {
            temp.Val, first = head.Val, true
        } else {
            if temp.Val != head.Val {
                temp.Next = &ListNode{Val: head.Val, Next: nil}
                temp = temp.Next
            }
        }
    }
    return res
}

func deleteDuplicates2(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    temp, current := head, head
    for ;current != nil; current = current.Next {
        if temp.Val != current.Val {
            temp.Next, temp = current, current
        }
    }
    temp.Next = nil
    return head
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    current := head
    for current.Next != nil {
        if current.Val == current.Next.Val {
            current.Next = current.Next.Next
        } else {
            current = current.Next
        }
    }
    return head
}