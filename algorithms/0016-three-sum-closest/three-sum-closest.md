## 问题
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:
```
Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
```

## 分析
找出数组nums中三个数之和与target最接近的值（每组数组唯一解）。
和15题类似

## 最高分
```golang
func threeSumClosest(nums []int, target int) int {
	near, result := math.MaxInt32, math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			subTarget := target - sum
			if subTarget == 0 {
				return sum
			}
			if subTarget > 0 {
				left++
			} else {
				right--
				subTarget = -subTarget
			}
			if near < 0 {
				near = -near
			} else if near > subTarget {
				near = subTarget
				result = sum
			}
		}
	}
	return result
}
```

## 实现
```golang
func threeSumClosest(nums []int, target int) int {
    length := len(nums)
    var left, right, sub int
    ret := math.MaxInt32
    // sort nums
    sort.Ints(nums)
    for i := 0; i < length-2; i++ {
        left = i + 1
        right = length - 1
        // while left < right
        for left < right {
            // jump repeat left value
            if left != i+1 && nums[left] == nums[left-1] {
                left++
                continue
            }
            sub = nums[i] + nums[right] + nums[left] - target
            // got it!
            if abs(sub) < abs(ret-target) {
                ret = sub + target
            }
            switch {
            case sub < 0:
                left++
            case sub > 0:
                right--
            default:
                return ret
            }
        }
    }
    return ret
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

## 改进
```golang

```

## 反思

## 参考
