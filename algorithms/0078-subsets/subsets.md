## 问题
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:
```
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```

## 分析
dfs实现，和77题、39题、40题类似

## 最高分
```golang
func subsets(nums []int) [][]int {
    res := make([][]int, 1, 1024)
    for _, n := range nums {
        for _, r := range res {
            res = append(res, append([]int{n}, r...))
        }
    }
    return res
}
```


## 实现
```golang
func subsets(nums []int) [][]int {
    var ret [][]int
    dfs(&ret, []int{}, nums, 0, len(nums))
    return ret
}

func dfs(ret *[][]int, tmp, nums []int, start, end int) {
    b := make([]int, len(tmp))
    copy(b, tmp)
    *ret = append(*ret, b)
    if start == end {
        return
    }
    for i := start; i < end; i++ {
        dfs(ret, append(tmp, nums[i]), nums, i+1, end)
    }
}
```

## 改进
```golang

```

## 反思

## 参考