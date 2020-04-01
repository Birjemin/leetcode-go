package convert_sorted_list_to_binary_search_tree

type ListNode struct {
    Val  int
    Next *ListNode
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
    idx := findIndex(head)
    if idx == nil {
        return nil
    }
    if idx.Next == nil {
        return &TreeNode{Val: idx.Val}
    }

    val, n := idx.Next.Val, idx.Next.Next
    idx.Next = nil

    return &TreeNode{
        Val:   val,
        Left:  sortedListToBST(head),
        Right: sortedListToBST(n)}
}

func findIndex(slow *ListNode) *ListNode {
    if slow == nil {
        return nil
    }
    if slow.Next == nil {
        return slow
    }
    fast := slow.Next
    for {
        fast = fast.Next
        if fast == nil || fast.Next == nil {
            return slow
        }
        fast, slow = fast.Next, slow.Next
    }
}
