package rotate_list

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

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }
    end, count := head, 1
    // find the end of ListNode and the number of ListNode
    for ; end.Next != nil; count++ {
        end = end.Next
    }
    if count == 1 {
        return head
    }
    if k %= count; k == 0 {
        return head
    }
    // connect the between of head node and end node，and move the end pointer to the first node
    end.Next, end = head, head
    // find the position of k
    for k = count - k; k > 1; k-- {
        end = end.Next
    }
    // replace the pointer
    head, end.Next = end.Next, nil
    return head
}
