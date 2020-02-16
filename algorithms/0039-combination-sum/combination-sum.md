## 问题
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:
```
Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
```

Example 2:
```
Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

## 分析
求出candidates中之和为target的所有组合，candidates中的数字可以重复使用，其实考察的就是dsf遍历

## 最高分
```golang
func combinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := make([][]int, 0)
    line := make([]int,0)
    calc(target, 0, line, candidates, &res)
    return res
}

func calc(target, index int, line []int, c []int, pres *[][]int){
    for i:= index;i<len(c);i++{
        line = append(line, c[i])
        //current sum
        sum:=0
        for _,v:=range line{
            sum+=v
        }
        if sum == target {
            temp := make([]int,len(line))
            copy(temp, line)
            *pres = append(*pres, temp)
            return
        }else if sum > target{
            return
        }
        // recursive
        calc(target, i, line, c, pres)
        //clean for next round
        line = line[:len(line)-1]
    }
}
```


## 实现
就是使用dsf迭代，和之前的回文判断类似
```golang
func combinationSum(candidates []int, target int) [][]int {
    var ret [][]int
    sort.Ints(candidates)
    dsf(&ret, candidates, []int{}, target)
    return ret
}

func dsf(ret *[][]int, candidates []int, solution []int, target int) {
    if target == 0 {
        // *ret = append(*ret, solution)
        b := make([]int, len(solution))
        copy(b, solution)
        *ret = append(*ret, b)
        return
    }
    for i, v := range candidates {
        // adding this expression could reduce the memory storage
        if v > target {
            return
        }
        dsf(ret, candidates[i:], append(solution, v), target-v)
    }
}
```

## 改进
根据最高分的解答，使用一个循环进行处理
```golang
func combinationSum(candidates []int, target int) [][]int {
    var ret [][]int
    sort.Ints(candidates)
    dsf(&ret, candidates, []int{}, target)
    return ret
}

func dsf(ret *[][]int, candidates []int, solution []int, target int) {
    for i, v := range candidates {
        if v > target {
            return
        }
        // add v
        solution = append(solution, v)
        if target == v {
            // *ret = append(*ret, solution)
            b := make([]int, len(solution))
            copy(b, solution)
            *ret = append(*ret, b)
            return
        }
        dsf(ret, candidates[i:], solution, target-v)
        // should subtract the last element of solution
        solution =solution[:len(solution)-1]
    }
}
```

## 反思

## 参考
