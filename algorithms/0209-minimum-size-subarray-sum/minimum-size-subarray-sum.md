## 问题
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example: 
```
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
```

Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n). 

## 分析
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续(!!!!连续~)子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

## 最高分
```golang
func minSubArrayLen(s int, nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    sums := make([]int, len(nums))
    sums[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        sums[i] = sums[i-1] + nums[i]
    }
    
    minLen := len(nums)+1

    for i := 0; i < len(sums); i++ {
        if sums[i] >= s {
            left := binarySearch(0, i-1, sums[i]-s, sums)
            minLen = min(minLen, i-left)
        }
    }
    
    if minLen == len(nums)+1 {
        return 0
    }
    return minLen
}

func binarySearch(left, right int, target int, sums []int) int {
    for left < right {
        mid := left + (right-left)/2
        if sums[mid] < target {
            left = mid+1
        } else {
            right = mid
        }
    }
    if sums[left] > target {
        return left - 1
    }
    return left
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```

## 实现
多推演一下就出来了，快慢指针的变种
```golang
func minSubArrayLen(s int, nums []int) int {

	length := len(nums)

	if length == 0 {
		return 0
	}

	var sum int
	var count int
	var slow int

	for fast, v := range nums {
		sum += v
		count++

		// find !
		if sum >= s {
			for slow < fast && fast < length && count != 1 {
				sum -= nums[slow]
				slow++
				if sum >= s {
					count--
				} else if fast < length-1 {
					sum += nums[fast+1]
					fast++
				}
			}
			return count
		}
	}
	return 0
}
```

## 改进
```golang
func minSubArrayLen(s int, nums []int) int {

	length := len(nums)

	if length == 0 {
		return 0
	}

	var (
		count int
		slow  int
		fast  int
		flag  bool
	)

	for i, v := range nums {
		s -= v

		// find !
		if s <= 0 {
			fast, count, flag = i, i+1, true
			break
		}
	}

	// not found
	if !flag {
		return 0
	}

	for slow < fast && fast < length && count != 1 {
		s += nums[slow]
		slow++
		if s <= 0 {
			count--
		} else if fast < length-1 {
			s -= nums[fast+1]
			fast++
		}
	}
	return count
}
```
## 改进
二分法的思想：

对于长度为 n 的数组，我们先去判断长度为 n/2 的连续数字中最大的和是否大于等于 s。

如果大于等于 s ，那么我们需要减少长度，继续判断所有长度为 n/4 的连续数字
如果小于 s，我们需要增加长度，我们继续判断所有长度为 (n/2 + n) / 2，也就是 3n/4 的连续数字。

个人感觉有点诡异~不过这个方法挺有趣的~（分治的思想）

## 反思

## 参考