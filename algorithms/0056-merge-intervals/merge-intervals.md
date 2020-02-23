## 问题
Given a collection of intervals, merge all overlapping intervals.

Example 1:
```
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

Example 2:
```
Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
```

## 分析
给出一个区间的集合，请合并所有重叠的区间。
- 先排序，然后判断会比较简单

## 最高分
```golang

```


## 实现
先排序，然后登记判断，存在则更新，不存在则添加
```golang
func merge(intervals [][]int) [][]int {
    length := len(intervals)
    if length == 0 {
        return [][]int{}
    } else if length == 1 {
        return intervals
    }
    
    // sort the arrays
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // init the ret
    ret, j := [][]int{{intervals[0][0], intervals[0][1]}}, 0
    for i := 1; i < length; i++ {
        if ret[j][1] >= intervals[i][0] {
            ret[j][1] = max(ret[j][1], intervals[i][1])
        } else {
            ret = append(ret, intervals[i])
            j++
        }
    }
    return ret
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
```

## 改进
```golang

```

## 反思

## 参考