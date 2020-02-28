## 问题
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

Example:
```
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

## 分析
dfs实现，和39题、40题类似

## 最高分
```golang
func combine(n int, k int) [][]int {
    combination := make([]int, k)
    res := [][]int{}

    var dfs func(int, int)
    dfs = func(idx, begin int) {
        if idx == k {
            temp := make([]int, k)
            copy(temp, combination)
            res = append(res, temp)
            return
        }

        for i := begin; i <= n+1-k+idx; i++ {
            combination[idx] = i
            dfs(idx+1, i+1)
        }
    }

    dfs(0, 1)

    return res
}
```


## 实现
dfs搜索即可
```golang
func combine(n int, k int) [][]int {
    var ret [][]int
    dfs(&ret, []int{}, 1, n, k)
    return ret
}

func dfs(ret *[][]int, tmp []int, min, max, k int) {
    if k == 0 {
        b := make([]int, len(tmp))
        copy(b, tmp)
        *ret = append(*ret, b)
        return
    }
    for i := min; i <= max-k+1; i++ {
        dfs(ret, append(tmp, i), i+1, max, k-1)
    }
}
```

## 改进
```golang

```

## 反思

## 参考