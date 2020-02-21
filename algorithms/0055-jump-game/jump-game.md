## 问题
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:
```
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

Example 2:
```
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

## 分析
给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个位置。dfs思想的应用
下面方式超时！！！！
```
func canJump(nums []int) bool {
    length := len(nums)
    return jump(&nums, 0, length)
}

func jump (nums *[]int, idx,length int) bool{
    if idx >= length-1 {
        return true
    }
    if (*nums)[idx] == 0 {
        return false
    }
    return jump(nums, idx + (*nums)[idx], length)
}
```

## 最高分
```golang
func canJump(nums []int) bool {
	cursor := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= cursor-i {
			cursor = i
		}
	}
	return cursor == 0
}
```

## 实现
很精炼的一种解法，需要耐心看看，和第一题类似，每次等级最大值。
```golang
func canJump(nums []int) bool {
    var k int
    for i, v := range nums {
        // if i > largest jumping distance, that mean never touch the end of array
        if i > k {
            return false
        }
        // get the largest jumping distance
        k = max(k, i+v)
    }
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## 改进
```golang

```

## 反思

## 参考