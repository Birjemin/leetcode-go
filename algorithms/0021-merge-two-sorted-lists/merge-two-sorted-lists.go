package merge_two_sorted_lists

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

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    var head, temp *ListNode
    if l1.Val > l2.Val {
        head = l2
        temp = l2
        l2 = l2.Next
    } else {
        head = l1
        temp = l1
        l1 = l1.Next
    }
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
           temp.Next = l2
           l2 = l2.Next
        } else {
           temp.Next = l1
           l1 = l1.Next
        }
        temp = temp.Next
    }

    if l1 == nil {
        temp.Next = l2
    } else {
        temp.Next = l1
    }
    return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    var head *ListNode
    if l1.Val > l2.Val {
        head = l2
        head.Next = mergeTwoLists(l2.Next, l1)
    } else {
        head = l1
        head.Next = mergeTwoLists(l1.Next, l2)
    }
    return head
}
