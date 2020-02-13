## 问题
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
```
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

## 分析
1 2 7 4 3 1
下一个排列为：
1 3 1 2 4 7
步骤如下：
- 从数组后往前，找到最长的降序排列位置（7的位置是2）: 1 2 7 4 3 1
- 把这个降序排列，转换成升序排列（7 - 1， 4 - 3调换）: 1 2 1 3 4 7
- 把序列前的元素，与序列中，第一个大于他的元素互换（2和后面3调换）: 1 3 1 2 4 7

## 最高分
```golang

```

## 实现
根据规则查找
```golang
func nextPermutation(nums []int) {
    length := len(nums)
    p, j := length-1, 0
    // find position
    for ; p > 0; p-- {
        if nums[p-1] < nums[p] {
            break
        }
    }
    // reverse the rest of array
    for ; length-p > j*2; j++ {
        nums[p+j], nums[length-j-1] = nums[length-j-1], nums[p+j]
    }
    // all array had been reversed
    if p == 0 {
        return
    }

    // find position for the next array
    for j = 0; nums[p-1] >= nums[p+j]; j++ {
    }
    nums[p+j], nums[p-1] = nums[p-1], nums[p+j]
    return
}
```

## 改进
```golang

```

## 反思

## 参考
