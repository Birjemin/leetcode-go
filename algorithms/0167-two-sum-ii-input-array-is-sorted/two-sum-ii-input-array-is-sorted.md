## 问题
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

- Your returned answers (both index1 and index2) are not zero-based.
- You may assume that each input would have exactly one solution and you may not use the same element twice.

Example:

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

## 分析
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
- 快慢指针法，一个从后往前fast，一个从前往后slow，如果和大于target，则fast-1，反之slow+1
## 最高分
```golang
func twoSum(numbers []int, target int) []int {
    i, j := 0,len(numbers)-1

    for ;i<j;{
        sum := numbers[i]+numbers[j]
        if (sum)==target {
            break
        } else if (sum) < target {
            i++
        } else {
            j--
        }
    }
    return []int{i + 1, j + 1}
}
```


## 实现
```golang
func twoSum(numbers []int, target int) []int {
    slow, fast := 1, len(numbers)
    for slow < fast {
        switch sum := numbers[slow-1] + numbers[fast-1] - target; {
        case sum > 0:
            fast -= 1
        case sum < 0:
            slow += 1
        default:
            return []int{slow, fast}
        }
    }
    return []int{}
}
```

## 改进
```golang

```

## 反思

## 参考