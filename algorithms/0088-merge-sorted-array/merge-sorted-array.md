## 问题
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:
```
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
```

## 分析
直接的思路就是复制一个temp，然后遍历

## 最高分
```golang
func merge(nums1 []int, m int, nums2 []int, n int) {
    i := m - 1
    j := n - 1
    k := m + n - 1
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
    for j >= 0 {
        nums1[k] = nums2[j]
        k--
        j--
    }
}
```

## 实现
```golang
func merge(nums1 []int, m int, nums2 []int, n int) {
    temp := make([]int, m)
    // 将nums1复制给temp
    copy(temp, nums1)
    // j -> nums1 k -> nums2
    j, k := 0, 0
    for i :=0; i < m + n; i++ {
        // nums1用完
        if j >= m {
            nums1[i] = nums2[k]
            k++
            continue
        }
        // nums2用完
        if k >= n {
            nums1[i] = temp[j]
            j++
            continue
        }
        if temp[j] > nums2[k] {
            nums1[i] = nums2[k]
            k++
        } else {
            nums1[i] = temp[j]
            j++
        }
    }
}
```

## 改进
比较大小，那么最后一位肯定是最大值，依次填充最后面的值即可
```golang
func merge(nums1 []int, m int, nums2 []int, n int) {
    i, j, k := m-1, n-1, m+n-1
    for ; i >= 0 && j >= 0; k-- {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
    }
    // j >= 0
    for ; j >= 0; j, k = j-1, k-1 {
        nums1[k] = nums2[j]
    }
}
```

## 反思
排序可以从前到后，也可以从后到前~~

## 参考