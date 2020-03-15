## 问题
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
```
Input: 3
Output: [1,3,3,1]
Follow up:
```

Could you optimize your algorithm to use only O(k) extra space?

## 分析
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k+1 行。

## 最高分
两次循环
```golang
func getRow(rowIndex int) []int {
    result := []int{1}
    for i := 0; i < rowIndex; i++ {
        result = append(result, 1)
        for j := i; j > 0; j-- {
            result[j] += result[j-1]
        }
    }
    return result
}
```


## 实现
118的基础上改进，其实就是将dp算法改一下~~~
```golang
func getRow(rowIndex int) []int {
    if rowIndex == 0 {
        return []int{1}
    }
    ret := make([]int, rowIndex+1)
    dp(&ret, rowIndex)
    return ret
}

func dp(r *[]int, n int) {
    (*r)[n] = 1
    if n == 0 {
        return
    }
    dp(r, n-1)
    // cal the collection of n-1
    for i := n - 1; i > 0; i-- {
        (*r)[i] += (*r)[i-1]
    }
}
```

## 改进
```golang

```

## 反思

## 参考