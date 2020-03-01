## 问题
Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:
```
Given nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.

It doesn't matter what you leave beyond the returned length.
```

Example 2:
```
Given nums = [0,0,1,1,1,1,2,3,3],

Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.

It doesn't matter what values are set beyond the returned length.
```

Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```
## 分析
删除排序数组中的重复项 II，26题的升级版本
快慢指针
重要！！！

## 最高分
一种很经典的做法，很难想到，快慢指针法，一个记录要替换的位置，一个负责遍历
```golang
func removeDuplicates(nums []int) int {
    length := len(nums)
    if length <= 2 {
        return length
    }
    j := 1
    for i := 2; i < length; i++ {
        // equal
        if nums[j-1] != nums[i] {
            j++
            nums[j] = nums[i]
        }
    }
    return j + 1
}
```

## 实现
```golang
func removeDuplicates(nums []int) int {
    length := len(nums)
    if length <= 2 {
        return length
    }
    j := 1
    for i := 2; i < length; i++ {
        // equal
        if nums[j-1] != nums[i] {
            j++
            nums[j] = nums[i]
        }
    }
    return j + 1
}
```

## 改进
```golang

```

## 反思

## 参考