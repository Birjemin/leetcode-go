## 问题
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
```
Input: [2,2,1]
Output: 1
```

Example 2:
```
Input: [4,1,2,1,2]
Output: 4
```

## 分析
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
- 借助map可以解决
- 文中不推荐申请内存，说明有其他方法

## 最高分
位运算，刘辟
```golang
func singleNumber(nums []int) int {
    length := len(nums)
    for i := 1; i < length; i++ {
        nums[0] ^= nums[i]
    }
    return nums[0]
}
```

## 实现
借助了map数据结构来解决
```golang
func singleNumber(nums []int) int {
    ret := make(map[int]bool, (len(nums)+1)/2)
    for _, v := range nums {
        if ret[v] {
            ret[v] = false
        } else {
            ret[v] = true
        }
    }
    for k, v := range ret {
        if v {
            return k
        }
    }
    return 0
}
```

## 改进
位运算
```golang
func singleNumber(nums []int) int {
    length := len(nums)
    for i := 1; i < length; i++ {
        nums[0] ^= nums[i]
    }
    return nums[0]
}
```

## 反思

## 参考