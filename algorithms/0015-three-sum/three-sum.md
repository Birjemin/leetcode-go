## 问题
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:
```
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 分析
这道题目就是从一个数组中找出所有的三个数字的子数组，子数组不重复，并且它们的和是0，和第一题很像，可以借鉴第一题思路：先数组排序，得到有序数组，然后使用双指针法进行摆动求解


## 最高分
- 首先对数组进行排序，排序后固定一个数 nums[i]，再使用左右指针指向 nums[i]后面的两端，数字分别为 nums[L] 和 nums[R]，计算三个数的和 sum 判断是否满足为 0，满足则添加进结果集
- 如果 nums[i] == nums[i−1]，则说明该数字重复，会导致结果重复，所以应该跳过
- 当 sum == 0 时，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
- 当 sum == 0 时，nums[R] == nums[R-1] 则会导致结果重复，应该跳过，R--

```golang
func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue//To prevent the repeat
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return results
}
```

## 实现
```golang
func threeSum(nums []int) [][]int {
    length := len(nums)
    var ret [][]int
    if length < 3 {
        return ret
    }
    var left, right, sum int
    // sort nums
    sort.Ints(nums)
    for i := 0; i < length-2; i++ {
        // impossible value
        if nums[i] > 0 {
            return ret
        }
        // jump repeat i value
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        left = i + 1
        right = length - 1
        // while left < right
        for left < right {
            // jump repeat left value
            if left != i+1 && nums[left] == nums[left-1] {
                left++
                continue
            }
            // question: why not jumping repeat right value?
            // got it!
            sum = nums[i] + nums[right] + nums[left]
            switch {
            case sum < 0:
                left++
            case sum > 0:
                right--
            default:
                ret = append(ret, []int{nums[i], nums[left], nums[right]})
                left++
                right--
            }
        }
    }
    return ret
}
```

## 改进
```golang

```

## 反思
双指针again

## 参考
