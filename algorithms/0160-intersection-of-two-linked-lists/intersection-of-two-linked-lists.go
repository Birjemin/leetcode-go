package intersection_of_two_linked_lists

type ListNode struct {
    Val  int
    Next *ListNode
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
    for ; headA != nil; headA = headA.Next {
        for tmp := headB; tmp != nil; tmp = tmp.Next {
            if headA == tmp {
                return headA
            }
        }
    }
    return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    tmpA, tmpB := headA, headB

    for tmpA != tmpB {
        if tmpA == nil {
            tmpA = headB
        } else {
            tmpA = tmpA.Next
        }
        if tmpB == nil {
            tmpB = headA
        } else {
            tmpB = tmpB.Next
        }
    }
    return tmpA
}
