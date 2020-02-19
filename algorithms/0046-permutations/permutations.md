## 问题
Given a collection of distinct integers, return all possible permutations.

Example:
```
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

## 分析
输出全排列的组合，其实和39，40类似，考察DFS算法

## 最高分
```golang
func permute(nums []int) [][]int {
    answer := make([][]int, 0)
    aux(&answer, 0, nums)
    return answer
}

func aux(answer *[][]int, idx int, nums []int) {
    if idx == len(nums) {
        c := make([]int, len(nums))
        copy(c, nums)
        *answer = append(*answer, c)
        return
    }
    for i := idx; i < len(nums); i++ {
        nums[idx], nums[i] = nums[i], nums[idx]
        aux(answer, idx + 1, nums)
        nums[i], nums[idx] = nums[idx], nums[i]
    }
    return
}
```

## 实现
注意map是引用，所以得copy，不然会改变原有的结构
```golang
func permute(nums []int) [][]int {
    length := len(nums)
    if length == 1 {
        return [][]int{
            nums,
        }
    }
    var ret [][]int
    dfs(&ret, nums, []int{})
    return ret
}

func dfs(ret *[][]int, nums, temp []int) {
    if len(nums) == 0 {
        b := make([]int, len(temp))
        copy(b, temp)
        *ret = append(*ret, b)
    }
    for i, v := range nums {
        if i == 0 {
            dfs(ret, nums[i+1:], append(temp, v))
        } else {
            b := make([]int, len(nums))
            copy(b, nums)
            dfs(ret, append(b[:i], b[i+1:]...), append(temp, v))
        }
    }
}
```

## 改进
```golang

```

## 反思

## 参考