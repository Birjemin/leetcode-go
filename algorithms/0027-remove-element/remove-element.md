## 问题
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:

```
Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
```

Example 2:

```
Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
```

Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```
// nums is passed in by reference. (i.e., without making a copy)
int len = removeElement(nums, val);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

## 分析
最直接的方式就是循环读取，然后标记找到的数字位

## 最高分
```golang


```

## 实现
```golang
func removeElement1(nums []int, val int) int {
    left, flag, num, length := 0, -1, 0, len(nums)
    for ;left < length; left++ {
        if nums[left] == val {
            if flag == -1 {
                flag = left
            }
        } else {
            if flag != -1 {
                nums[flag], nums[left] = nums[left], nums[flag]
                left = flag
                flag = -1
            }
            num++
        }
    }
    return num
}
```

## 改进
只需要计算不等于的值就好了，其余的值不用管
```golang
func removeElement(nums []int, val int) int {
    var l int
    for _, v := range nums {
        if v != val {
            nums[l] = v
            l++
        }
    }
    return l
}
```

## 反思
只计算需要计算的就好了，不需要将问题想复杂

## 参考
