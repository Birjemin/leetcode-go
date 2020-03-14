## 问题
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

```
      0
     / \
   -3   9
   /   /
 -10  5
```

## 分析


## 最高分
减少了一个判断，不使用切片
```golang
func sortedArrayToBST(nums []int) *TreeNode {
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, i, j int) *TreeNode {
    if len(nums) == 0 || i > j || i < 0 || j >= len(nums) {
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
```

## 实现
自己回溯
```golang
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
```

## 改进
```golang

```

## 反思

## 参考