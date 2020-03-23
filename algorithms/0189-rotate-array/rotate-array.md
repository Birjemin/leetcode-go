## 问题
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:
```
Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

Example 2:
```
Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
```

Note:

- Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
- Could you do it in-place with O(1) extra space?

## 分析
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

这种解法改变了内存存储，不符合要求
```golang
func rotate1(nums []int, k int) {
    length:= len(nums)
    nums = append(nums[length-k:], nums[:length-k]...)
}
```

## 最高分
```golang
func rotate(nums []int, k int) {
    k %= len(nums)
    reverse(nums, 0, len(nums)-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, len(nums)-1)
    return
}

func reverse(nums []int, i, j int) {
    for ; i < j; i, j = i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
    return
}
```

## 实现
最简单的方式，但是效率极低
```golang
func rotate(nums []int, k int) {
    length := len(nums)
    if length == 1 {
        return
    }
    for i := 0; i < k; i++ {
        for j := length; j > 1; j-- {
            nums[j-1], nums[j-2] = nums[j-2], nums[j-1];
        }
    }
}
```

## 改进
三次翻转，想不到。。。
```golang

```

## 反思

## 参考