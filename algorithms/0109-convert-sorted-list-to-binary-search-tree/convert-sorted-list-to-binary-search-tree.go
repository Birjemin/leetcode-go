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
    if head == nil {
        return nil
    }
    temp := toArr(head)
    return cal(temp)
}

func cal(arr []int) *TreeNode {
    length := len(arr)
    if length == 0 {
        return nil
    }
    mid := length / 2
    return &TreeNode{
        Val:   arr[mid],
        Left:  cal(arr[:mid]),
        Right: cal(arr[mid+1:]),
    }
}

func toArr(head *ListNode) (ret []int) {
    for ; head != nil; head = head.Next {
        ret = append(ret, head.Val)
    }
    return
}