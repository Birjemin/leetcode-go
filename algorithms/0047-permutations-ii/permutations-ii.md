## 问题
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:
```
Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```

## 分析
和46题一样，dfs的使用，先排序，然后去重

## 最高分
```golang
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	doPermute(&res, []int{}, nums)
	return res
}

func doPermute(res *[][]int, currentPerm []int, nums []int) {
	if len(nums) == 0 {
		*res = append(*res, currentPerm)
		return
	}

	// deal with the first index
	prev := nums[0]
	newPerm := make([]int, len(currentPerm)+1)
	copy(newPerm[:len(currentPerm)], currentPerm)
	newPerm[len(currentPerm)] = nums[0]
	doPermute(res, newPerm, nums[1:])

	if len(nums) == 1 {
		return
	}
	// deal with remaining, while skipping if the value is same as the previous index
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			continue
		}
		prev = nums[i]
		newPerm := make([]int, len(currentPerm)+1)
		copy(newPerm[:len(currentPerm)], currentPerm)
		newPerm[len(currentPerm)] = nums[i]

		newNums := make([]int, len(nums)-1)
		copy(newNums[:i], nums[:i])
		copy(newNums[i:], nums[i+1:])
		doPermute(res, newPerm, newNums)
	}
}
```

## 实现
先排序，然后跳过相同的值
```golang
func permuteUnique(nums []int) [][]int {
    length := len(nums)
    if length == 1 {
        return [][]int{
            nums,
        }
    }
    var ret [][]int
    sort.Ints(nums)
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
            // if the value is equal to the front value, just continue the circle
            if i != 0 && nums[i-1] == v {
                continue
            }
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