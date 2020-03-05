## 问题
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

Example:
```
Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
```

## 分析
和78题类似，区别在于collection有重复的元素，dfs算法

## 最高分
```golang
func subsetsWithDup(nums []int) [][]int {
    res := [][]int{[]int{}}
    var cur []int

    sort.Ints(nums)
    doSubsetsWithDup(nums, cur, &res)
    return res
}

func doSubsetsWithDup(nums []int, cur []int, res *[][]int) {
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        cur = append(cur, nums[i])
        newRes := make([]int, len(cur))
        copy(newRes, cur)
        *res = append(*res, newRes)

        doSubsetsWithDup(nums[i+1:], cur, res)
        cur = cur[:len(cur)-1]
    }
}
```

## 实现
78题目的答案类似，重复的元素直接跳过就好了
```golang
func subsetsWithDup(nums []int) [][]int {
    var ret [][]int
    sort.Ints(nums)
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
        if i > start && nums[i] == nums[i-1] {
            continue
        }
        dfs(ret, append(tmp, nums[i]), nums, i+1, end)
    }
}
```

## 改进
```golang

```

## 反思

## 参考