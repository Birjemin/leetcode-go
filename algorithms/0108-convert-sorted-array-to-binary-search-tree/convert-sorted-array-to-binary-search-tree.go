package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    length := len(nums)
    if length == 0 {
        return nil
    }
    p := length / 2
    return &TreeNode{
        Val:   nums[p],
        Left:  sortedArrayToBST(nums[:p]),
        Right: sortedArrayToBST(nums[p+1:]),
    }
}

func sortedArrayToBST1(nums []int) *TreeNode {
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, i, j int) *TreeNode {
    length := len(nums)
    if length == 0 || i > j || i < 0 || j >= length {
        return nil
    }
    if i == j {
        return &TreeNode{Val: nums[i]}
    }
    mid := i + (j-i)/2
    return &TreeNode{
        Val:   nums[mid],
        Left:  helper(nums, i, mid-1),
        Right: helper(nums, mid+1, j),
    }
}