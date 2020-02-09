package remove_nth_node_from_end_of_list

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var length int
    temp := head
    // count ListNode length
    for temp != nil {
        length++
        temp = temp.Next
    }
    // just one node
    if length == 1 {
        return nil
    } else if n == length {
        // head node
        return head.Next
    } else {
        n, length = length-n, 0
        for temp = head; temp != nil; temp = temp.Next {
            length++
            if n == length {
                if temp.Next != nil {
                    temp.Next = temp.Next.Next
                }
                break
            }
        }
        return head
    }
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
    temp := head
    // get temp position
    for i := 0; i < n; i++ {
        temp = temp.Next
    }
    // n == length of ListNode
    if temp == nil {
        return head.Next
    }
    // head node
    prev := head
    // both of prev and temp should attempt to move
    for temp.Next != nil {
        temp, prev = temp.Next, prev.Next
    }
    // if not equal nil, connection the next node
    if prev.Next != nil {
        prev.Next = prev.Next.Next
    }
    return head
}
