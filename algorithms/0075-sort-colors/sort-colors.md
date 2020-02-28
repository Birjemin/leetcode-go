## 问题
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:
```
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

Follow up:

- A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
- Could you come up with a one-pass algorithm using only constant space?

## 分析
荷兰国旗问题，提示给了一种方式，就是计算出每个值的个数，然后累加

## 最高分
```golang
func sortColors(a []int) {
    i, j, k := 0, 0, len(a)-1
    // for 循环中， nums[i:j] 中始终全是 1
    // 循环结束后，
    // nums[:i] 中全是 0
    // nums[j:] 中全是 2
    for j <= k {
        switch a[j] {
        case 0:
            a[i], a[j] = a[j], a[i]
            i++
            j++
        case 1:
            j++
        case 2:
            a[j], a[k] = a[k], a[j]
            k--
        }
    }

}
```

## 实现
三个变量，将left监听0的边界，right监听2的边界，curr负责移动
```golang
func sortColors(nums []int) {
    right := len(nums) - 1
    if right <= 0 {
        return
    }
    var left, curr int
    for curr <= right {
        switch nums[curr] {
        case 0:
            nums[left], nums[curr] = nums[curr], nums[left]
            left++
            curr++
        case 2:
            nums[curr], nums[right] = nums[right], nums[curr]
            right--
        default:
            curr++
        }
    }
    return
}
```

## 改进
```golang

```

## 反思

## 参考