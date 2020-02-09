## 问题
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:
```
Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
```

## 分析
和15题类似解答

## 最高分
```golang
func fourSum(nums []int, target int) [][]int {
    var result [][]int
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[length-1]+nums[length-2]+nums[length-2] < target {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			head, end := j+1, length-1
			for head < end {
				tempt := nums[i] + nums[j] + nums[head] + nums[end]
				if (tempt == target) {
					var lengthArray []int
					lengthArray = append(lengthArray, nums[i])
					lengthArray = append(lengthArray, nums[j])
					lengthArray = append(lengthArray, nums[head])
					lengthArray = append(lengthArray, nums[end])
					result = append(result, lengthArray)
					head++;
					for head < end && nums[head] == nums[head-1] {
						head++;
					}
				} else if tempt > target {
					end--
				} else {
					head++
				}
			}
		}
	}
	return result
}
```

## 实现
```golang
func fourSum(nums []int, target int) [][]int {
    length := len(nums)
    var ret [][]int
    if length < 3 {
        return ret
    }
    var left, right, sum int
    // sort nums
    sort.Ints(nums)
    for j := 0; j < length-3; j++ {
        // jump repeat j value
        if j > 0 && nums[j] == nums[j-1] {
            continue
        }
        for i := j + 1; i < length-2; i++ {
            // impossible value
            if target > 0 && nums[j]+nums[i] > target {
                break
            }
            // jump repeat i value
            if i > j+1 && nums[i] == nums[i-1] {
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
                // got it!
                sum = nums[j] + nums[i] + nums[right] + nums[left]
                switch {
                case sum < target:
                    left++
                case sum > target:
                    right--
                default:
                    ret = append(ret, []int{nums[j], nums[i], nums[left], nums[right]})
                    left++
                    right--
                }
            }
        }
    }
    return ret
}
```

## 改进
根据最高分的解答，增加了边界检查，提高了30%提交~~
```golang
func fourSum(nums []int, target int) [][]int {
    length := len(nums)
    var ret [][]int
    if length < 3 {
        return ret
    }
    var left, right, sum int
    // sort nums
    sort.Ints(nums)
    for j := 0; j < length-3; j++ {
        // jump repeat j value
        if j > 0 && nums[j] == nums[j-1] {
            continue
        }
        // boundary check
        if nums[j]+nums[length-3]+nums[length-2]+nums[length-1] < target {
            continue
        }
        for i := j + 1; i < length-2; i++ {
            // impossible value
            if target > 0 && nums[j]+nums[i] > target {
                break
            }
            // jump repeat i value
            if i > j+1 && nums[i] == nums[i-1] {
                continue
            }
            // boundary check
            if nums[j]+nums[i]+nums[length-2]+nums[length-1] < target {
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
                // got it!
                sum = nums[j] + nums[i] + nums[right] + nums[left]
                switch {
                case sum < target:
                    left++
                case sum > target:
                    right--
                default:
                    ret = append(ret, []int{nums[j], nums[i], nums[left], nums[right]})
                    left++
                    right--
                }
            }
        }
    }
    return ret
}

```

## 反思

## 参考
