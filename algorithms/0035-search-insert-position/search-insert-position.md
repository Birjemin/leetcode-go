## 问题
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
```
Input: [1,3,5,6], 5
Output: 2
```
Example 2:
```
Input: [1,3,5,6], 2
Output: 1
```
Example 3:
```
Input: [1,3,5,6], 7
Output: 4
```
Example 4:
```
Input: [1,3,5,6], 0
Output: 0
```

## 分析


## 最高分
```golang
func searchInsert(nums []int, target int) int {
    previous := 0
    for i, v := range nums {
        if v >= target{
            return i
        }
        previous++
    }
    return previous
}
```

## 实现
```golang
func searchInsert(nums []int, target int) int {
    for k, v := range nums {
        if v >= target {
            return k
        }
    }
    return len(nums)
}
```

## 改进
减少内存使用
```golang
func searchInsert(nums []int, target int) int {
    var count int
    for k, v := range nums {
        if v >= target {
            return k
        }
        count++
    }
    return count
}
```

## 反思

## 参考
