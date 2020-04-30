## 问题
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.

Example 1:
```
Input: [3,4,5,1,2] 
Output: 1
```

Example 2:
```
Input: [4,5,6,7,0,1,2]
Output: 0
```

## 分析
假设按照升序排序的数组在预先未知的某个点上进行了旋转。请找出其中最小的元素。
找到旋转点即可

## 最高分
```golang
func findMin(nums []int) int {
	max := len(nums)
    min := nums[0] 
	for i := range nums {
		if i+1 < max  && nums[i] > nums[i+1] {
			min = nums[i+1]
            break
		}
	}
    return min
}
```


## 实现
找到旋转点就好了
```golang
func findMin(nums []int) int {
    for i, v := range nums {
        if i == 0 {
            continue
        }
        if v < nums[i-1] {
            return v
        }
    }
    return nums[0]
}
```

## 改进
二分法，max mid min，如果max > mid，说明旋转点在前方，反之在后方
```golang
func findMin1(nums []int) int {
    left, right := 0, len(nums)-1
    for {

        if right-left == 0 {
            return nums[0]
        } else if right-left == 1 {
            if nums[left] < nums[right] {
                return nums[0]
            }
            return nums[right]
        }

        // split the array
        mid := (left + right) / 2
        if nums[left] > nums[mid] {
            right = mid
        } else {
            left = mid
        }
    }
}
```

## 反思

## 参考