## 问题
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

```
Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```
Example 2:

```
Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
```

## 分析
求出candidates中之和为target的所有组合，candidates中的数字不可以重复使用，其实考察的还是dsf遍历

## 最高分
```golang
func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    
    line:=make([]int,0)
    res:=make([][]int,0)
    calc(0, target, candidates, line, &res)
    
    return res
}

func calc(index int, t int, c []int, line []int, res *[][]int){
    for i:=index;i<len(c);{
        line = append(line, c[i])
        sum:=0
        for _,v := range line{
            sum+=v 
        }
        if sum == t{
            temp:=make([]int,len(line))
            copy(temp, line)
            *res = append(*res, temp)
            return
        }
        if sum > t{
            return
        }
        
        calc(i+1, t, c, line, res)
        line = line[:len(line)-1]

        i++
        for ;i<len(c) && c[i] == c[i-1]; {
            i++
        }
    }
}
```

## 实现
- 如果找到了，则跳出当前循环，回溯到上一级
- 如果有相同的值，需要跳到下一个值
```golang
func combinationSum2(candidates []int, target int) [][]int {
    var ret [][]int
    sort.Ints(candidates)
    dfs(&ret, candidates, []int{}, target)
    return ret
}

func dfs(ret *[][]int, candidates []int, solution []int, target int) bool {
    if target == 0 {
        // *ret = append(*ret, solution)
        b := make([]int, len(solution))
        copy(b, solution)
        *ret = append(*ret, b)
        return true
    }
    for i, v := range candidates {
        // adding this expression could reduce the memory storage
        if v > target {
            return false
        }
        // next value equal to the current value, just jump up this value
        if i != 0 && candidates[i-1] == v {
            continue
        }
        // if find the value, then break this circle
        if dfs(ret, candidates[i+1:], append(solution, v), target-v) {
            break
        }
    }
    return false
}
```

## 改进
```golang

```

## 反思

## 参考
