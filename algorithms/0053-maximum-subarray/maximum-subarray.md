## 问题
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## 分析
得到一个数列中的最大子序列，并返回结果

## 最高分
```golang
func maxSubArray(nums []int) int {
	max, sum := nums[0],nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
```


## 实现
最直接的思路，不停的计算，先找到每次去找循环的最大值
```golang
func maxSubArray(nums []int) int {
    length := len(nums)
    if length == 1 {
        return nums[0]
    }
    flag := nums[0]

    for idx, n := range nums {
        // 当前循环n大于上一次循环的flag，则flag更新
        if flag < n {
            flag = n
        }
        // 在当前循环中找到最大值flag
        for right := idx + 1; right < length; right++ {
            n += nums[right]
            if flag < n {
                flag = n
            }
        }
    }
    return flag
}
```

## 改进
最大子序列，如果序列有正负数，则子序列第一项一定是正数，如果只有负数，则为最大的负数
```golang
func maxSubArray(nums []int) int {
    max, sum := nums[0], nums[0]
    for _, v := range nums[1:] {
        if sum < 0 {
            // 如果是负数，则sum重新计算，v可能还是负数，那么下一次循环会sum会覆盖，但是直到sum不是负数
            sum = v
        } else {
            // 如果是正数，则sum加
            sum += v
        }
        // 如果sum > max则覆盖结果
        if max < sum {
            max = sum
        }
    }
    return max
}

```

## 反思

## 参考
