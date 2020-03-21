## 问题
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:
```
Input: [3,2,3]
Output: 3
```

Example 2:
```
Input: [2,2,1,1,1,2,2]
Output: 2
```

## 分析
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

## 最高分
```golang

```


## 实现
最简单的方式就是加一个map处理
```golang
func majorityElement(nums []int) int {
    m := make(map[int]int, len(nums)/2)
    for _, v := range nums {
        if _, ok := m[v]; ok {
            m[v]++
        } else {
            m[v] = 1
        }
    }
    var ret, count int
    for i, v := range m {
        if count < v {
            count, ret = v, i
        }
    }
    return ret
}

```

## 改进
先排序，因为length > n/2，这说明n/2位置的数字绝对是所求的值
```golang
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}
```

## 改进
Boyer-Moore 投票算法，刘辟刘辟
```
func majorityElement(nums []int) int {
    maj := 0
    num := 0
    for _, v := range nums {
        if v != maj && num <= 0 {
            maj = v
        } else if v != maj {
            num -- 
        } else {
            num ++
        }
    }
    return maj
}
```

## 反思

## 参考